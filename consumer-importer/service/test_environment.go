package service

import "consumer-importer/service/util"

var dbProperties = util.DbProperties{
	Host:     "localhost",
	Port:     "5432",
	User:     "pguser",
	Password: "p0stgr3s",
	Dbname:   "consumer_importer",
}
