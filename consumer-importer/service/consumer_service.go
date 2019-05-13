package service

import (
	"consumer-importer/model"
	"consumer-importer/service/util"
	"consumer-importer/validator"
)

// ConsumerService service responsible for managing Consumer data to the postgres database
type ConsumerService struct {
	Props *util.DbProperties
}

//Save store a Consumer object into Consumer Database Table
func (c *ConsumerService) Save(dto *model.Consumer) error {
	err := validator.ValidateConsumer(dto)
	valid := err == nil

	var validationMessage *string

	if err != nil {
		strErr := err.Error()
		validationMessage = &strErr
	} else {
		validationMessage = nil
	}

	db := util.OpenConnection(c.Props)
	defer util.CloseConnection(db)

	sqlStatement := `INSERT INTO consumer 
		(id_file, 
		cpf, 
		private, 
		incompleto,	
		data_ultima_compra,	
		ticket_medio,
		ticket_ultima_compra, 
		loja_mais_frequente, 
		loja_ultima_compra,
		valid,
		validation_message)
	VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	RETURNING ID;`

	return db.QueryRow(sqlStatement,
		dto.IDFile,
		dto.CPF,
		dto.Private,
		dto.Incompleto,
		dto.DataUltimaCompra,
		dto.TicketUltimaCompra,
		dto.TicketMedio,
		dto.LojaMaisFrequente,
		dto.LojaUltimaCompra,
		valid,
		validationMessage).Scan(&dto.ID)
}

//Delete delete a consumer from database
func (c *ConsumerService) Delete(ID int64) (bool, error) {
	db := util.OpenConnection(c.Props)
	defer util.CloseConnection(db)

	sqlStatement := `DELETE FROM consumer
		WHERE id = $1`

	result, err := db.Exec(sqlStatement, ID)
	rowsAffected, _ := result.RowsAffected()

	return rowsAffected == 1, err
}
