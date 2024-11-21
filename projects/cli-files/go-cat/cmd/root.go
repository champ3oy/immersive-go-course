package cmd

import (
	"fmt"
	"os"
)

func Execute() error {
	args := os.Args
	if len(args) > 1 {
		content, err := os.ReadFile(args[1])
		if err != nil {
			return fmt.Errorf("error %v", err)
		}

		fmt.Fprintln(os.Stdout, string(content))
		return nil
	}

	return fmt.Errorf("no argument was passed to the command")
}
