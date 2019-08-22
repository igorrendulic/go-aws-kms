package ckms

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
)

var (
	// ErrFileNotFound if path to config file incorrect
	ErrFileNotFound = errors.New("file not found")

	// ErrFileExistsNoOverwrite not allowed to overwrite existing files
	ErrFileExistsNoOverwrite = errors.New("cannot overrite existing file")
)

func loadFile(path string) ([]byte, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, ErrFileNotFound
	}
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return content, nil
}

// SaveFile saved the file in custom path (created dir if it doesn't exist)
func SaveFile(raw []byte, path string) error {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return ErrFileExistsNoOverwrite
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = f.Write(raw)
	if err != nil {
		return err
	}
	return nil
}
