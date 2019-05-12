package service

import (
	"consumer-importer/model"
	"testing"
	"time"
)

var consumerService = ConsumerService{
	&dbProperties,
}

func TestStoreConsumer(t *testing.T) {
	cpf := "86631560900"
	lojaMaisFrequente := "79379491000183"
	lojaUltimaCompra := "79379491000183"
	ticketMedio := 335.38
	ticketUltimaCompra := 335.38

	consumer := model.Consumer{
		CPF:                &cpf,
		DataUltimaCompra:   time.Date(2010, 1, 13, 0, 0, 0, 0, time.UTC),
		IDFile:             nil,
		Incompleto:         false,
		LojaMaisFrequente:  &lojaMaisFrequente,
		LojaUltimaCompra:   &lojaUltimaCompra,
		Private:            false,
		TicketMedio:        &ticketMedio,
		TicketUltimaCompra: &ticketUltimaCompra,
	}

	err := consumerService.Save(&consumer)

	if consumer.ID == nil {
		t.Errorf("Error inserting a Consumer: ID shouldn't be nil")
	}

	if err != nil {
		t.Errorf("File saving a consumer: " + err.Error())
	}

	consumerService.Delete(*consumer.ID)
}
