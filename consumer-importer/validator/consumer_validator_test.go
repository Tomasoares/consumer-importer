package validator

import (
	"consumer-importer/model"
	"testing"
)

func TestValidConsumer(t *testing.T) {
	cpf := "04109164125"
	lojaMaisFrequente := "79379491000183"
	lojaUltimaCompra := "79379491000850"

	c := model.Consumer{
		CPF:               &cpf,
		LojaMaisFrequente: &lojaMaisFrequente,
		LojaUltimaCompra:  &lojaUltimaCompra,
	}

	err := ValidateConsumer(&c)

	if err != nil {
		t.Errorf("should've been a valid consumer: " + err.Error())
	}
}

func TestInalidCPFConsumer(t *testing.T) {
	cpf := "24102162125"
	lojaMaisFrequente := "79379491000183"
	lojaUltimaCompra := "79379491000850"

	c := model.Consumer{
		CPF:               &cpf,
		LojaMaisFrequente: &lojaMaisFrequente,
		LojaUltimaCompra:  &lojaUltimaCompra,
	}

	if ValidateConsumer(&c) == nil {
		t.Errorf("should've been a consumer with invalid CPF")
	}
}

func TestInvalidLojaMaisFrequenteConsumer(t *testing.T) {
	cpf := "04109164125"
	lojaMaisFrequente := "793791000183"
	lojaUltimaCompra := "79379491000850"

	c := model.Consumer{
		CPF:               &cpf,
		LojaMaisFrequente: &lojaMaisFrequente,
		LojaUltimaCompra:  &lojaUltimaCompra,
	}

	if ValidateConsumer(&c) == nil {
		t.Errorf("should've been a consumer with invalid CNPJ from LojaMaisFrequente")
	}
}

func TestInvalidLojaUltimaCompraConsumer(t *testing.T) {
	cpf := "04109164125"
	lojaMaisFrequente := "79379491000183"
	lojaUltimaCompra := "793794900850"

	c := model.Consumer{
		CPF:               &cpf,
		LojaMaisFrequente: &lojaMaisFrequente,
		LojaUltimaCompra:  &lojaUltimaCompra,
	}

	if ValidateConsumer(&c) == nil {
		t.Errorf("should've been a consumer with invalid CNPJ from LojaUltimaCompra")
	}
}
