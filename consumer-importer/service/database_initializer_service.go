package service

import (
	"consumer-importer/service/util"
	"database/sql"
)

// DatabaseInitializerService service responsible for creating all the tables for the postgres database
type DatabaseInitializerService struct {
	Props *util.DbProperties
}

//InitializeDBTables Initialize the tables related to the Consumer Service
func (d *DatabaseInitializerService) InitializeDBTables() {
	db := util.OpenConnection(d.Props)
	defer util.CloseConnection(db)

	d.initializeFileTable(db)
	d.initializeConsumerTable(db)
}

func (d *DatabaseInitializerService) initializeFileTable(db *sql.DB) {
	sqlStatement := `CREATE TABLE IF NOT EXISTS file ( 
		id 						SERIAL PRIMARY KEY,
		import_date				date,
		name					varchar(255),
		successful				bool
	);`

	_, err := db.Exec(sqlStatement)

	if err != nil {
		panic(err)
	}
}

func (d *DatabaseInitializerService) initializeConsumerTable(db *sql.DB) {
	sqlStatement := `CREATE TABLE IF NOT EXISTS consumer ( 
		id 						SERIAL PRIMARY KEY,
		id_file					int,
		cpf 					varchar(11),
		private 				bool,
		incompleto 				bool,
		data_ultima_compra 		date,
		ticket_medio 			DOUBLE PRECISION,
		ticket_ultima_compra 	DOUBLE PRECISION,
		loja_mais_frequente		varchar(14),
		loja_ultima_compra		varchar(14),
		constraint fk_consumer_file foreign key (id_file) references file (id)
	);`

	_, err := db.Exec(sqlStatement)

	if err != nil {
		panic(err)
	}
}
