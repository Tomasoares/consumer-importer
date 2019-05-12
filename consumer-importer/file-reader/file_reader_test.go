package filereader

import (
	"testing"
)

var filename = "../file-repository/base_teste.txt"

func TestReadFileSuccessfully(t *testing.T) {
	fileReader := FileReader{
		fileName: filename,
	}

	err := fileReader.Initialize()

	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestCloseUnopenedFileSuccessfully(t *testing.T) {
	fileReader := FileReader{
		fileName: filename,
	}

	err := fileReader.CloseFile()

	if err == nil {
		t.Errorf("shoud've sent error")
	}
}

func TestOpenCloseFileSuccessfully(t *testing.T) {
	fileReader := FileReader{
		fileName: filename,
	}

	fileReader.Initialize()
	err := fileReader.CloseFile()

	if err != nil {
		t.Errorf("error closing the file: " + err.Error())
	}
}

func TestReadFirstLineOfFile(t *testing.T) {
	fileReader := FileReader{
		fileName: filename,
	}

	fileReader.Initialize()
	hasText := fileReader.Next()

	if !hasText {
		t.Errorf("file shouldn't be empty")
	}

	line := fileReader.ReadLine()

	if line == "" {
		t.Errorf("line shouldn't be empty")
	}

	if line != "CPF                PRIVATE     INCOMPLETO  DATA DA ÚLTIMA COMPRA TICKET MÉDIO          TICKET DA ÚLTIMA COMPRA LOJA MAIS FREQUÊNTE LOJA DA ÚLTIMA COMPRA" {
		t.Errorf("line is not the first one")
	}

	fileReader.CloseFile()
}
