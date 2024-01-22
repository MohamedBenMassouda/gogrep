package src

import (
	"fmt"
	"os"
)

func CheckArgs() (string, string) {
	args := os.Args[1:]

	if len(args) < 1 || args[0] == "" {
		fmt.Println("Please enter a string to search")
		os.Exit(0)
	} else if len(args) > 2 {
		fmt.Println("Too many arguments")
		os.Exit(0)
	}

	var path, _ = os.Getwd()

	if len(args) >= 2 {
		path = args[1]
	}

	if path == "~" {
		path, _ = os.UserHomeDir()
	}

	return path, args[0]
}
