package util

import (
	"testing"
)

func TestDBConnection(t *testing.T) {
	dbProperties := DbProperties{
		Host:     "localhost",
		Port:     "5000",
		User:     "user",
		Password: "postgrespassword",
		Dbname:   "consumer_importer",
	}

	db := OpenConnection(&dbProperties)
	CloseConnection(db)
}
