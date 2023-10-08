package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

func checkInPath(path string, input string) {
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
		var listFiles []string = getDirFiles(path)

		for _, file := range listFiles {
			checkInPath(path+"/"+file, input)
		}
	} else {
		fmt.Println(getRandomColor() + path)
		lookInContent(getFileContent(path), input)
	}
}

func getDirFiles(path string) []string {
	var files []string

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
		if !fileInfo.IsDir() {
			files = append(files, fileInfo.Name())
		}
	}

	return files
}

func getRandomColor() string {
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	colorBlue := "\033[34m"
	colorPurple := "\033[35m"
	colorCyan := "\033[36m"

	colors := []string{colorRed, colorGreen, colorYellow, colorBlue, colorPurple, colorCyan}

	return colors[rand.Intn(len(colors))]
}

// Get file content
func getFileContent(fileName string) string {
	fileContent, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(fileContent)
}

func lookInContent(content string, input string) {
	var lineNumber int = 1
	for _, line := range strings.Split(content, "\n") {
		if strings.Contains(line, input) {
			// Get the line number
			print(line, input, lineNumber)
		}
		lineNumber++
	}
}

func print(text string, input string, lineNumber int) {
	colorReset := "\033[0m"

	// Print the line normally but the input in color
	tx := getRandomColor() + fmt.Sprint(lineNumber) + colorReset + ":" + text
	fmt.Println(strings.Replace(tx, input, getRandomColor()+input+colorReset, -1))
}

func main() {
	if os.Args[1] == "" {
		fmt.Println("Please enter a string to search")
		return
	}

	var path, _ = os.Getwd()
	if len(os.Args) > 2 {
		path = os.Args[2]
	}

	if path == ".." {
		path = ".."
	} else if path == "~" {
		path, _ = os.UserHomeDir()
	}

	// checkInPath(os.Args[1], os.Args[2])
	var listFiles []string = getDirFiles(path)

	for _, file := range listFiles {
		checkInPath(path+"/"+file, os.Args[1])
	}
}
