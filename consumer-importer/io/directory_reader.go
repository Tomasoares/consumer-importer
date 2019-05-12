package io

import (
	"io/ioutil"
)

//GetFilesFromDirectory get all existing files within a directory
func GetFilesFromDirectory(directory string) []string {
	files, err := ioutil.ReadDir(directory)

	if err != nil {
		panic(err)
	}

	filesNames := make([]string, len(files))

	for i, f := range files {
		filesNames[i] = directory + f.Name()
	}

	return filesNames
}
