package src

import "os"

func GetDirs(path string) []os.DirEntry {
	dirs, err := os.ReadDir(path)

	if err != nil {
		panic(err)
	}

	return dirs
}
