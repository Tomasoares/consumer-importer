package io

import (
	"bufio"
	"errors"
	"os"
)

//FileReader struct responsible for reading a file
type FileReader struct {
	FilePath string
	file     *os.File
	scanner  *bufio.Scanner
}

//Initialize opens the file and initialize a scanner reader
func (f *FileReader) Initialize() error {
	err := f.openFile()

	if err != nil {
		return err
	}

	f.scanner = bufio.NewScanner(f.file)

	return f.scanner.Err()
}

func (f *FileReader) openFile() error {
	var err error
	f.file, err = os.Open(f.FilePath)
	return err
}

//CloseFile close the file
func (f *FileReader) CloseFile() error {
	if f.file != nil {
		return f.file.Close()
	}

	return errors.New("no file to close")
}

//ReadLine read the current line pointed by the scanner
func (f *FileReader) ReadLine() string {
	return f.scanner.Text()
}

//Next check if file has next line and move scanner pointer to that line
func (f *FileReader) Next() bool {
	return f.scanner.Scan()
}
