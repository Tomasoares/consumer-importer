package filereader

import (
	"bufio"
	"errors"
	"os"
)

type FileReader struct {
	fileName string
	file     *os.File
	scanner  *bufio.Scanner
}

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
	f.file, err = os.Open(f.fileName)
	return err
}

func (f *FileReader) CloseFile() error {
	if f.file != nil {
		return f.file.Close()
	} else {
		return errors.New("no file to close")
	}
}

func (f *FileReader) ReadLine() string {
	return f.scanner.Text()
}

func (f *FileReader) Next() bool {
	return f.scanner.Scan()
}
