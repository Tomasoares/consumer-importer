package validator

import (
	"consumer-importer/model"
	"errors"
)

//ValidateConsumer checks if consumer has valid CPF amd CNPJ from LojaMaisFrequente and LojaUltimaCompra fields
func ValidateConsumer(consumer *model.Consumer) error {

	if consumer.CPF != nil && !ValidateCpf(*consumer.CPF) {
		return errors.New("Invalid CPF")
	}

	if consumer.LojaMaisFrequente != nil && !ValidateCNPJ(*consumer.LojaMaisFrequente) {
		return errors.New("Invalid CNPJ on LojaMaisFrequente field")
	}

	if consumer.LojaUltimaCompra != nil && !ValidateCNPJ(*consumer.LojaUltimaCompra) {
		return errors.New("Invalid CNPJ on LojaUltimaCompra field")
	}

	return nil
}
