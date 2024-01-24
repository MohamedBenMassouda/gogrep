package src

import (
	"fmt"
	"os"
	"strings"
)

func GetDirFiles(path string) []os.FileInfo {
	var files []os.FileInfo

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return files
	}

	fileInfos, err := file.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return files
	}

	for _, fileInfo := range fileInfos {
		files = append(files, fileInfo)
	}

	return files
}

// Get file content
func GetFileContent(fileName string) string {
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(fileContent)
}

func LookInContent(path string, content string, input string) {
	var lineNumber int = 1
	var shouldPrintPath bool = false

	for _, line := range strings.Split(content, "\n") {
		if strings.Contains(line, input) {
			if !shouldPrintPath {
				fmt.Println(GetRandomColor() + path)
				shouldPrintPath = true
			}

			// Get the line number
			Print(line, input, lineNumber)
		}
		lineNumber++
	}
}
