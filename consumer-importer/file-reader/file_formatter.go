package filereader

import (
	"consumer-importer/util"
)

func FormatFile(row string) util.FileRow {
	return util.FileRow{
		CPF:                row[:util.CPF_END],
		Private:            row[util.CPF_END:util.PRIVATE_END],
		Incompleto:         row[util.PRIVATE_END:util.INCOMPLETO_END],
		DataUltimaCompra:   row[util.INCOMPLETO_END:util.DATA_ULTIMA_COMPRA_END],
		TicketMedio:        row[util.DATA_ULTIMA_COMPRA_END:util.TICKET_MEDIO_END],
		TicketUltimaCompra: row[util.TICKET_MEDIO_END:util.TICKET_ULTIMA_COMPRA_END],
		LojaMaisFrequente:  row[util.TICKET_ULTIMA_COMPRA_END:util.LOJA_MAIS_FREQUENTE_END],
		LojaUltimaCompra:   row[util.LOJA_MAIS_FREQUENTE_END:],
	}
}

func ToConsumer(row util.FileRow) {

}
