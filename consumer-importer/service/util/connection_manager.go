package util

import (
	"database/sql"
	"fmt"

	//opens connection to postgres database
	_ "github.com/lib/pq"
)

//OpenConnection creates a new connection to a database
func OpenConnection(props *DbProperties) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		props.Host, props.Port, props.User, props.Password, props.Dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

//CloseConnection closes a database connection
func CloseConnection(db *sql.DB) {
	err := db.Close()

	if err != nil {
		panic(err)
	}
}
