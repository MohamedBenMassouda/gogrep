package src

import (
	"fmt"
	"os"
	"strings"
)

func Print(text string, input string, lineNumber int) {
	// if -1 it will highlight all the occurrences
	// 0 will not highlight any occurrences

	var howManyTimesShouldBeHighlighted int = -1
	colorReset := "\033[0m"

	// Print the line normally but the input in color
	tx := GetRandomColor() + fmt.Sprint(lineNumber) + colorReset + ":" + text
	fmt.Println(strings.Replace(tx, input, GetRandomColor()+input+colorReset, howManyTimesShouldBeHighlighted))
}

func PrintAllFiles(path string, ignore ...string) {
	if path == ".git" {
		return
	}
	var listFiles []os.FileInfo = GetDirFiles(path)

	for _, file := range listFiles {
		if file.IsDir() {
			PrintAllFiles(path + "/" + file.Name())
		} else {
			fmt.Println(path + "/" + file.Name())
		}
	}
}
