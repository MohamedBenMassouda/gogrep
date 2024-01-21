package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// TODO: Try to make it more efficient, by reading .gitignore file and ignoring the files in it
var ignoreDirs = []string{".git", "node_modules", ".vscode", "vendor", ".idea", "build"}

func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

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
		if Contains(ignoreDirs, file.Name()) {
			return
		}

		var listFiles []os.FileInfo = getDirFiles(path)

		for _, file := range listFiles {
			checkInPath(path+"/"+file.Name(), input)
		}
	} else {
		content := getFileContent(path)
		if content != "" {
			lookInContent(path, content, input)
		}
	}
}

func getDirFiles(path string) []os.FileInfo {
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
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(fileContent)
}

func lookInContent(path string, content string, input string) {
	var lineNumber int = 1
	var shouldPrintPath bool = false

	for _, line := range strings.Split(content, "\n") {
		if strings.Contains(line, input) {
			if !shouldPrintPath {
				fmt.Println(getRandomColor() + path)
				shouldPrintPath = true
			}

			// Get the line number
			Print(line, input, lineNumber)
		}
		lineNumber++
	}
}

func Print(text string, input string, lineNumber int) {
	// if -1 it will highlight all the occurrences
	// 0 will not highlight any occurrences

	var howManyTimesShouldBeHighlighted int = -1
	colorReset := "\033[0m"

	// Print the line normally but the input in color
	tx := getRandomColor() + fmt.Sprint(lineNumber) + colorReset + ":" + text
	fmt.Println(strings.Replace(tx, input, getRandomColor()+input+colorReset, howManyTimesShouldBeHighlighted))
}

func PrintAllFiles(path string, ignore ...string) {
	if path == ".git" {
		return
	}
	var listFiles []os.FileInfo = getDirFiles(path)

	for _, file := range listFiles {
		if file.IsDir() {
			PrintAllFiles(path + "/" + file.Name())
		} else {
			fmt.Println(path + "/" + file.Name())
		}
	}
}

func CheckArgs() string {
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

	if path == ".." {
		path = ".."
	} else if path == "~" {
		path, _ = os.UserHomeDir()
	}

	return path
}

func main() {
	path := CheckArgs()
	var input string = os.Args[1]

	if file, _ := os.Stat(path); !file.IsDir() {
		checkInPath(path, input)
		return
	}

	var listFiles []os.FileInfo = getDirFiles(path)

	for _, file := range listFiles {
		checkInPath(path+"/"+file.Name(), input)
	}
}
