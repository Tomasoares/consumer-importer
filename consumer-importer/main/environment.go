package main

import (
	"consumer-importer/service/util"
	"os"
)

var dbProperties = util.DbProperties{
	Host:     os.Getenv("PG_HOST"),
	Port:     os.Getenv("PG_PORT"),
	User:     os.Getenv("PG_USER"),
	Password: os.Getenv("PG_PASSWORD"),
	Dbname:   os.Getenv("PG_DATABASE"),
}

var directory = os.Getenv("FILE_REPOSITORY")
