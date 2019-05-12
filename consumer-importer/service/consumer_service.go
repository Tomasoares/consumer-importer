package service

import (
	"consumer-importer/model"
	"consumer-importer/service/util"
)

// ConsumerService service responsible for managing Consumer data to the postgres database
type ConsumerService struct {
	Props *util.DbProperties
}

//Save store a Consumer object into Consumer Database Table
func (c *ConsumerService) Save(dto *model.Consumer) {
	db := util.OpenConnection(c.Props)
	defer util.CloseConnection(db)

	sqlStatement := ``

	_, err := db.Exec(sqlStatement)

	if err != nil {
		panic(err)
	}
}
