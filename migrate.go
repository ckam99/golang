package migrate

import (
	"bufio"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	TIME_LAYOUT = "20060201150405999"
)

type migration struct {
	*sql.DB
	Config  *Config
	baseDir string
	driver  string
}

type Config struct {
	Table string
}

type version struct {
	version int64
	dirty   *bool
}

func New(driver, dsn, baseDir string, cfg *Config) (*migration, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	if cfg.Table == "" {
		cfg.Table = "migrations"
	}
	m := &migration{DB: db, Config: cfg, baseDir: baseDir, driver: driver}
	return m, m.createTable()
}

// * Run migrations
func (m *migration) Migrate() error {
	files, lastVersion, err := m.getFiles("up")
	if err != nil {
		return err
	}
	var newVersion int
	for _, f := range files {
		b, err := os.ReadFile(m.baseDir + "/" + f.Name())
		if err != nil {
			return err
		}
		_, err = m.Exec(string(b))
		if err != nil {
			log.Fatalln(string(b), err)
			return err
		}
		newVersion, _ = strconv.Atoi(strings.Split(f.Name(), "_")[0])
		fmt.Println(f.Name(), "successfuly migrated")
	}

	if newVersion > 0 && newVersion != int(lastVersion.version) {
		_, err = m.Exec(
			fmt.Sprintf(`insert into %s(version) values($1);`, m.Config.Table),
			newVersion,
		)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("no change")
	}
	return nil
}

// ! Rollback all migrations from database
func (m *migration) Rollback() error {
	if confirm("Are you sure to rollback migrations from database?") {
		files, lastVersion, err := m.getFiles("down")
		if err != nil {
			return err
		}
		if lastVersion.version == 0 || len(files) == 0 {
			fmt.Println("no change")
			return nil
		}
		for _, f := range files {
			b, err := os.ReadFile(m.baseDir + "/" + f.Name())
			if err != nil {
				return err
			}
			_, err = m.Exec(string(b))
			if err != nil {
				log.Fatalln(string(b), err)
				return err
			}
			fmt.Println(f.Name(), "successfuly rollback")
		}
		q := fmt.Sprintf(`truncate table %s;`, m.Config.Table)
		if m.driver == "sqlite3" || m.driver == "sqlite" {
			q = fmt.Sprintf(`delete from %s;`, m.Config.Table)
		}
		_, err = m.Exec(q)
		if err != nil {
			log.Println(q)
			return err
		}
		return nil
	}
	return nil
}

// Create new migration
func CommandLine() {

	const StringHelp = `
Usage:
	migrate <command> [arguments]
The commands are:
	migrate help - Help
	migrate up - Migration database
	migrate down - Database Rollback
	migrate create - Create migration
Use "migrate help" for more information about a command.
	`

	dir := flag.String("dir", "./migration", "Migration directory")
	name := flag.String("name", "", "Migration name")
	database := flag.String("database", "", "Database url")
	driver := flag.String("driver", "", "Database driver")

	flag.Parse()

	if *dir == "" {
		fmt.Println("Migration directory is required")
		return
	}

	if len(flag.Args()) == 0 {
		fmt.Println(StringHelp)
		return
	}

	switch flag.Args()[0] {
	case "up":
		if *database == "" && *driver == "" {
			fmt.Println("Database url and driver are required")
			return
		}
		mi, err := New(*driver, *database, *dir, &Config{})
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if flag.Args()[0] == "up" {
			if err = mi.Migrate(); err != nil {
				fmt.Println(err.Error())
				return
			}
		}
		if flag.Args()[0] == "down" {
			if err = mi.Rollback(); err != nil {
				fmt.Println(err.Error())
				return
			}
		}
		return
	case "create":
		if *name == "" {
			fmt.Println("Migration name required")
			return
		}
		Create(*dir, *name)
		return
	default:
		fmt.Println("Command is not available")
		fmt.Println(StringHelp)
		return
	}
}

func Create(dir, name string) {
	baseName := fmt.Sprintf("%s/%s_%s", dir, time.Now().Format("20060201150405999"), name)
	file, err := os.Create(fmt.Sprintf("%s.up.sql", baseName))
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer file.Close()
	_, err = os.Create(fmt.Sprintf("%s.down.sql", baseName))
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Printf(" migration %s.down.sql\n", baseName)
}

func (m *migration) checkVersion() (version, error) {
	var v version
	if err := m.QueryRow(fmt.Sprintf("select * from %s order by version desc limit 1", m.Config.Table)).
		Scan(&v.version, &v.dirty); err != nil {
		if err != sql.ErrNoRows {
			return version{}, errors.New("migration " + err.Error())
		}
	}
	if v.dirty != nil {
		if *v.dirty {
			return version{}, errors.New("migration is dirty")
		}
	}
	return v, nil
}

func (m *migration) createTable() error {
	_, err := m.Exec(fmt.Sprintf(`create table if not exists %s(
		 version varchar(60) not null unique,
		 dirty bool default(false)
	     );
		`, m.Config.Table))
	return err
}

func (m *migration) getFiles(n string) (fss []fs.DirEntry, lastVersion version, err error) {
	files, err := os.ReadDir(m.baseDir)
	if err != nil {
		return fss, lastVersion, err
	}
	lastVersion, err = m.checkVersion()
	if err != nil {
		return fss, lastVersion, err
	}
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), "."+n+".sql") {
			currentVersion, _ := strconv.Atoi(strings.Split(f.Name(), "_")[0])
			if currentVersion == 0 || int64(currentVersion) <= lastVersion.version {
				continue
			}
			fss = append(fss, f)
		}
	}
	sort.Slice(fss, func(i, j int) bool {
		return fss[i].Name() < fss[j].Name()
	})
	return fss, lastVersion, nil
}

func confirm(s string) bool {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s [y/n]: ", s)
		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}
