package main

import (
	"consumer-importer/io"
	"consumer-importer/model"
	"consumer-importer/service"
	"consumer-importer/service/util"
	"fmt"
	"time"
)

var dbProperties = util.DbProperties{
	Host:     "localhost",
	Port:     5000,
	User:     "user",
	Password: "postgrespassword",
	Dbname:   "consumer_importer",
}

var consumerService = service.ConsumerService{Props: &dbProperties}
var fileService = service.FileService{Props: &dbProperties}

var directory = "../file-repository/"
var readFiles = make([]string, 0)

//BeginService starts up the process to read the consumer files and import to a postgres database
func BeginService() {
	fmt.Println("Starting Consumer File Importer service...")

	for {
		fmt.Println("Checking for existing files in directory: " + directory)
		filenames := io.GetFilesFromDirectory(directory)
		unreadFiles := getUnreadFiles(filenames)

		if len(unreadFiles) == 0 {
			fmt.Println("No files to load!")
		} else {
			doImportProcess(filenames)
		}

		time.Sleep(5 * time.Second)
	}
}

func getUnreadFiles(filenames []string) []string {
	unreadFiles := make([]string, 0)

	for _, file := range filenames {
		exists := false

		for _, unreadFile := range readFiles {
			if file == unreadFile {
				exists = true
				break
			}
		}

		if !exists {
			unreadFiles = append(unreadFiles, file)
		}
	}

	return unreadFiles
}

func doImportProcess(filenames []string) {
	for _, filename := range filenames {
		alreadyRead, err := fileService.IsFileAlreadyRead(filename)

		if err != nil {
			panic(err)
		}

		if alreadyRead {
			readFiles = append(readFiles, filename)
			fmt.Println("File " + filename + " has already been read!")
			continue
		}

		path := directory + filename

		fileReader := io.FileReader{
			FilePath: path,
		}

		fileReader.Initialize()
		defer fileReader.CloseFile()

		savedFile := storeFile(filename)

		for fileReader.Next() {
			line := fileReader.ReadLine()
			fmt.Println("Reading line: " + line)

			if io.IsHeader(line) {
				continue
			}

			consumer := io.Format(line)
			consumerService.Save(&consumer)
		}

		fileService.UpdateFileAsSuccessful(*savedFile.ID)
		fmt.Println("File " + filename + " has been successfully read!")
	}
}

func storeFile(filename string) *model.File {
	file := model.File{
		ImportDate: time.Now(),
		Name:       filename,
		Successful: false,
	}

	fileService.Save(&file)

	return &file
}
