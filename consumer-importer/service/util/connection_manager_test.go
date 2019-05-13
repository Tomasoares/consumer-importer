package util

import (
	"testing"
)

func TestDBConnection(t *testing.T) {
	dbProperties := DbProperties{
		Host:     "localhost",
		Port:     "5432",
		User:     "pguser",
		Password: "p0stgr3s",
		Dbname:   "consumer_importer",
	}

	db := OpenConnection(&dbProperties)
	CloseConnection(db)
}
