package main

import (
	"consumer-importer/io"
	"fmt"
)

func main() {
	fileNames := io.GetFilesFromDirectory("../file-repository/")

	for _, file := range fileNames {
		fmt.Println(file)
	}

}
