package src

import (
	"fmt"
	"os"
)

func CheckInPath(path string, input string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("Path does not exist")
		return
	}

	file, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	if file.IsDir() {
		if Contains(ignoreDirs, file.Name()) {
			return
		}

		var listFiles []os.FileInfo = GetDirFiles(path)

		for _, file := range listFiles {
			CheckInPath(path+"/"+file.Name(), input)
		}
	} else {
		if IsImageOrVideoFile(path) {
			return
		}

		content := GetFileContent(path)
		if content != "" {
			LookInContent(path, content, input)
		}
	}
}
