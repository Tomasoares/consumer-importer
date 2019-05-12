package service

import (
	"consumer-importer/service/util"
	"testing"
)

var dbProperties = util.DbProperties{
	Host:     "localhost",
	Port:     5000,
	User:     "user",
	Password: "postgrespassword",
	Dbname:   "consumer_importer",
}

func TestDBConnection(t *testing.T) {
	service := ConsumerService{dbProperties}
	db := service.openConnection()
	service.closeConnection(db)
}

func TestInitializationTables(t *testing.T) {
	service := ConsumerService{dbProperties}
	service.InitializeDBTables()
}
