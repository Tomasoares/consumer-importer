package service

import "consumer-importer/service/util"

var dbProperties = util.DbProperties{
	Host:     "localhost",
	Port:     "5000",
	User:     "user",
	Password: "postgrespassword",
	Dbname:   "consumer_importer",
}
