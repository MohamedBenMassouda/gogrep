package src

import (
	"fmt"
	"os"
)

func CheckArgs() (string, string) {
	args := os.Args[1:]

	if args[0] == "-h" || args[0] == "--help" {
		fmt.Println("Usage: gogrep [string] [path]")
		os.Exit(0)
	} else if Contains(args, "-f") || Contains(args, "--files") {
		files := GetDirFiles(".")

		PrintAllFiles(files)
		os.Exit(0)
	} else if Contains(args, "-d") || Contains(args, "--dirs") {
		dirs := GetDirs(".")
		PrintAllDirs(dirs)
		os.Exit(0)
	}

	if len(args) < 1 || args[0] == "" {
		fmt.Println("Please enter a string to search")
		os.Exit(0)
	} else if len(args) > 2 {
		fmt.Println("Too many arguments")
		os.Exit(0)
	}

	var path = "."

	if len(args) >= 2 {
		path = args[1]
	}

	if path == "~" {
		path, _ = os.UserHomeDir()
	}

	return path, args[0]
}
