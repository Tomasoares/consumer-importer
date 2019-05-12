package service

import (
	"testing"
)

func TestInitializationTables(t *testing.T) {
	service := DatabaseInitializerService{&dbProperties}
	service.InitializeDBTables()
}
