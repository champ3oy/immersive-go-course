package main

import (
	"fmt"
	"go-cat/cmd"
	"os"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(0)
}
