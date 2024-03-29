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

func PrintAllFiles(files []os.FileInfo) {
	for _, file := range files {
		fmt.Println(GetRandomColor() + file.Name() + "\033[0m")
	}
}

func PrintAllDirs(dirs []os.DirEntry) {
	for _, dir := range dirs {
		if dir.IsDir() {
			fmt.Println(GetRandomColor() + dir.Name() + "\033[0m")
		}
	}
}
