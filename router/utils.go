package router

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && len(value) > 0 {
		return value
	}
	return fallback
}
func GetEnvInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok && len(value) > 0 {
		if i, err := strconv.ParseInt(value, 10, 32); err == nil {
			return int(i)
		}
	}
	return fallback
}
