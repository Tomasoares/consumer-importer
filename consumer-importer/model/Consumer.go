package model

import "time"

type Consumer struct { 
	CPF string
	Private bool
	Incompleto bool
	DataUltimaCompra time.Time
	TicketMedio float32
	TicketUltimaCompra float32
	LojaMaisFrequente string
	LojaDaUltimaCompra string
}
