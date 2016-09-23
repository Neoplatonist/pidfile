// Package pidfile manages pid files.
package pidfile

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

// ReadPid reads pid from file
func ReadPid(path string) (int, error) {
	if path == "" {
		return 0, errors.New("Path is empty")
	}

	s, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, err
	}

	pid, err := strconv.Atoi(string(bytes.TrimSpace(s)))
	if err != nil {
		return 0, errors.New("Error parsing pid from path")
	}

	return pid, nil
}

// WritePid writes pid to file
func WritePid(pid int, path string) error {
	if pid == 0 || path == "" {
		return errors.New("Pid equals 0 or Path is empty")
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	s := strconv.Itoa(pid)

	t, err := f.WriteString(s)
	if err != nil {
		return err
	}

	f.Sync()

	log.Printf("Saved %d to %s", t, path)

	return nil
}
