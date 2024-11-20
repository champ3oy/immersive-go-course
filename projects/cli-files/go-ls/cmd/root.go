package cmd

import (
	"fmt"
	"os"
)

func Execute() error {
	if len(os.Args[1]) > 0 {
		files, err := os.ReadDir(os.Args[1])
		if err != nil {
			return fmt.Errorf("error reading files from dir")
		}

		for _, file := range files {
			fmt.Println(file.Name())
		}

		return nil

	} else {
		files, err := os.ReadDir(".")
		if err != nil {
			return fmt.Errorf("error reading files from dir")
		}

		for _, file := range files {
			fmt.Println(file.Name())
		}

		return nil
	}
}
