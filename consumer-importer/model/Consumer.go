package model

import "time"

//Consumer formatted struct with data related to a Consumer
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
