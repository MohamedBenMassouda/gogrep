package src

import (
	"bufio"
	"os"
)

func Run() {

	// Check if data is being piped to stdin
	stat, _ := os.Stdin.Stat()

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// data is being piped to stdin
		scanner := bufio.NewScanner(os.Stdin)

		var input string = ""
		for scanner.Scan() {
			input += scanner.Text() + "\n"
		}

		if input != "" {
			LookInContent("", input, os.Args[1])
		}

		return
	}

	path, input := CheckArgs()

	if file, _ := os.Stat(path); !file.IsDir() {
		CheckInPath(path, input)
		return
	}

	var listFiles []os.FileInfo = GetDirFiles(path)

	for _, file := range listFiles {
		CheckInPath(path+"/"+file.Name(), input)
	}

}
