package migrate

import (
	"bufio"
	"database/sql"
	"errors"
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
}

type Config struct {
	Table string
}

type version struct {
	version int64
	dirty   *bool
}

func New(baseDir, driver, dsn string, cfg *Config) (*migration, error) {
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
	m := &migration{DB: db, Config: cfg, baseDir: baseDir}
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
	files, lastVersion, err := m.getFiles("down")
	if lastVersion.version != 0 && confirm("Are you sure to rollback migrations from database?") {
		if err != nil {
			return err
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
		_, err = m.Exec(
			fmt.Sprintf(`truncate table %s;`, m.Config.Table),
		)
		return err
	}
	fmt.Println("no change")
	return nil
}

// *Create new migration
func (m *migration) Create() error {
	var name string
	fmt.Print("Enter migration name: ")
	fmt.Scanf("%s", &name)
	baseName := fmt.Sprintf("%s/%s_%s", m.baseDir, time.Now().Format("20060201150405999"), name)
	file, err := os.Create(fmt.Sprintf("%s.up.sql", baseName))
	if err != nil {
		return err
	}
	defer file.Close()
	fmt.Printf(" migration %s.up.sql\n", baseName)
	_, err = os.Create(fmt.Sprintf("%s.down.sql", baseName))
	if err != nil {
		return err
	}
	fmt.Printf(" migration %s.down.sql\n", baseName)
	return nil
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
