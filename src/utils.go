package src

import (
	"path/filepath"
	"strings"
)

// TODO: Try to make it more efficient, by reading .gitignore file and ignoring the files in it
var ignoreDirs = []string{".git", "node_modules", ".vscode", "vendor", ".idea", "build"}

func IsImageOrVideoFile(fileName string) bool {
	imageExtensions := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp"}
	videoExtensions := []string{".mp4", ".avi", ".mkv", ".mov", ".wmv"}

	ext := strings.ToLower(filepath.Ext(fileName))

	for _, imgExt := range imageExtensions {
		if ext == imgExt {
			return true
		}
	}

	for _, vidExt := range videoExtensions {
		if ext == vidExt {
			return true
		}
	}

	return false
}

func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
