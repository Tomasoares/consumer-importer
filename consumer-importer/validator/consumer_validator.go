package validator

import (
	"consumer-importer/model"
)

//ValidateConsumer checks if consumer has valid CPF amd CNPJ from LojaMaisFrequente and LojaUltimaCompra fields
func ValidateConsumer(consumer *model.Consumer) bool {
	if consumer.CPF != nil && ValidateCpf(*consumer.CPF) {
		return false
	}

	if consumer.LojaMaisFrequente != nil && ValidateCNPJ(*consumer.LojaMaisFrequente) {
		return false
	}

	if consumer.LojaUltimaCompra != nil && ValidateCNPJ(*consumer.LojaUltimaCompra) {
		return false
	}

	return true
}
