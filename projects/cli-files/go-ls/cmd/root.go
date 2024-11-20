package cmd

import (
	"fmt"
	"os"
)

func Execute() error {
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	err := readDir(dir)
	return err
}

func readDir(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("error reading files from dir")
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	return nil
}
