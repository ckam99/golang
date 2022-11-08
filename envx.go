package envx

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Get(key string, defaultValue interface{}) interface{} {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	return defaultValue
}

func GetString(key string, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	return defaultValue
}

func GetInt(key string, defaultValue int) int {
	val, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return defaultValue
	}
	return val
}

func Scan(files ...string) error {
	if len(files) == 0 {
		files = []string{".env"}
	}
	for _, f := range files {
		file, err := os.Open(f)
		defer func() {
			if err := file.Close(); err != nil {
				log.Println(err)
			}
		}()
		if err != nil {
			return err
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			keys := strings.Split(scanner.Text(), "=")
			if len(keys) < 2 {
				continue
			}
			if err := os.Setenv(keys[0], keys[1]); err != nil {
				return err
			}
		}
		if err := scanner.Err(); err != nil {
			return err
		}
	}
	return nil
}

func SetEnv(f string, value interface{}) error {
	return os.Setenv(f, fmt.Sprint(value))
}

// GetLookup: get env if does not exists will be set
func GetLookup(key string, defaultValue interface{}) interface{} {
	if err := Scan(); err != nil {
		fmt.Println(err)
	}
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	return defaultValue
}
