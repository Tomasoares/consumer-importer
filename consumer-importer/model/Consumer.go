package model

import "time"

type Consumer struct {
	CPF                *string
	Private            bool
	Incompleto         bool
	DataUltimaCompra   time.Time
	TicketMedio        *float64
	TicketUltimaCompra *float64
	LojaMaisFrequente  *string
	LojaUltimaCompra   *string
}
