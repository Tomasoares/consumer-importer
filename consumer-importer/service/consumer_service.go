package service

import (
	"consumer-importer/service/util"
	"database/sql"
	"fmt"

	//opens connection to postgres database
	_ "github.com/lib/pq"
)

type ConsumerService struct {
	props util.DbProperties
}

func (c *ConsumerService) openConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.props.Host, c.props.Port, c.props.User, c.props.Password, c.props.Dbname)

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

func (c *ConsumerService) closeConnection(db *sql.DB) {
	err := db.Close()

	if err != nil {
		panic(err)
	}
}

func (c *ConsumerService) InitializeDBTables() {
	db := c.openConnection()
	defer c.closeConnection(db)

	sqlStatement := `CREATE TABLE IF NOT EXISTS test ( id SERIAL PRIMARY KEY );`

	_, err := db.Exec(sqlStatement)

	if err != nil {
		panic(err)
	}
}
