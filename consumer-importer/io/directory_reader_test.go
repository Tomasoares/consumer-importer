package io

import (
	"testing"
)

func TestGetFilesFromDirectory(t *testing.T) {
	fileNames := GetFilesFromDirectory("../file-repository/")

	if len(fileNames) == 0 {
		t.Errorf("No files to read!")
	}

	if fileNames[0] == "file-repository/base_teste.txt" {
		t.Errorf("Wrong filename: " + fileNames[0])
	}
}
