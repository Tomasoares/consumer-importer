package service

import (
	"consumer-importer/model"
	"testing"
	"time"
)

var fileService = FileService{
	&dbProperties,
}

func TestStoreFile(t *testing.T) {
	file := model.File{
		Name:       "test",
		ImportDate: time.Now(),
		Successful: false,
	}

	err := fileService.Save(&file)

	if file.ID == nil {
		t.Errorf("Error inserting a File: ID shouldn't be nil")
	}

	if err != nil {
		t.Errorf("File saving file: " + err.Error())
	}

	fileService.Delete(*file.ID)
}

func TestDeleteFile(t *testing.T) {
	file := model.File{
		Name:       "test",
		ImportDate: time.Now(),
		Successful: false,
	}

	fileService.Save(&file)
	successful, err := fileService.Delete(*file.ID)

	if err != nil {
		t.Errorf("Error deleting a File: " + err.Error())
	}

	if !successful {
		t.Errorf("File hasn't been deleted")
	}
}

func TestFindFile(t *testing.T) {
	file := model.File{
		Name:       "test",
		ImportDate: time.Date(2019, 5, 12, 0, 0, 0, 0, time.FixedZone("", 0)),
		Successful: false,
	}

	fileService.Save(&file)
	defer fileService.Delete(*file.ID)
	searchedFile, err := fileService.Find(*file.ID)

	if err != nil {
		t.Errorf("File searching file: " + err.Error())
	}

	if searchedFile == nil {
		t.Errorf("File hasn't been found")
	}

	if file.String() != searchedFile.String() {
		t.Errorf("Searched file is different from the original: " + file.String() +
			", searched: " + searchedFile.String())
	}
}

func TestUpdateFileAsSuccessful(t *testing.T) {
	file := model.File{
		Name:       "test",
		ImportDate: time.Date(2019, 5, 12, 0, 0, 0, 0, time.FixedZone("", 0)),
		Successful: false,
	}

	fileService.Save(&file)
	defer fileService.Delete(*file.ID)
	success, err := fileService.UpdateFileAsSuccessful(*file.ID)

	if err != nil {
		t.Errorf("Error updating file: " + err.Error())
	}

	if !success {
		t.Errorf("File hasn't been updated")
	}
}

func TestIsFileAlreadyBeenRead(t *testing.T) {
	file := model.File{
		Name:       "test",
		ImportDate: time.Date(2019, 5, 12, 0, 0, 0, 0, time.FixedZone("", 0)),
		Successful: true,
	}

	fileService.Save(&file)
	defer fileService.Delete(*file.ID)
	success, err := fileService.IsFileAlreadyRead(file.Name)

	if err != nil {
		t.Errorf("Error trying to invoke isFileAlreadyRead: " + err.Error())
	}

	if !success {
		t.Errorf("File should've been read")
	}
}
