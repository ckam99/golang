package migrate

import (
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

type Migration struct {
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

func New(driver, dsn, baseDir string) (*Migration, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	m := &Migration{DB: db, Config: &Config{
		Table: "migrations",
	}, baseDir: baseDir, driver: driver}
	return m, m.createTable()
}

// * Run migrations
func (m *Migration) Migrate() error {
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
func (m *Migration) Rollback() error {
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

// Create new migration
func CommandLine() {
	const Help = `
Usage:
	migrate <command> [arguments]
The commands are:
	help - Help
	up - Migration database
	down - Database Rollback
	create - Create migration
	help - for more information about a command.
`
	dir := flag.String("dir", "./internal/migration", "Migration directory")
	database := flag.String("database", "", "Database url")
	driver := flag.String("driver", "", "Database driver")

	flag.Parse()

	if *dir == "" {
		fmt.Println("Migration directory is required")
		return
	}

	if len(flag.Args()) == 0 {
		fmt.Print(Help)
		os.Exit(1)
	}

	switch flag.Args()[0] {
	case "up", "down":
		if *database == "" && *driver == "" {
			fmt.Println("Database url and driver are required")
			return
		}
		mi, err := New(*driver, *database, *dir)
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
		cmd := flag.NewFlagSet("create", flag.ExitOnError)
		name := cmd.String("name", "", "Migration name")
		if err := cmd.Parse(os.Args[2:]); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if *name == "" {
			fmt.Println("Migration name required")
			return
		}
		Create(*dir, *name)
		return
	default:
		fmt.Println("Command is not available")
		fmt.Print(Help)
		return
	}
}

func Create(dir, name string) {
	baseName := fmt.Sprintf("%s/%s_%s", dir, time.Now().Format(TIME_LAYOUT), name)
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

func (m *Migration) checkVersion() (version, error) {
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

func (m *Migration) createTable() error {
	_, err := m.Exec(fmt.Sprintf(`create table if not exists %s(
		 version varchar(60) not null unique,
		 dirty bool default(false)
	     );
		`, m.Config.Table))
	return err
}

func (m *Migration) getFiles(n string) (fss []fs.DirEntry, lastVersion version, err error) {
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
