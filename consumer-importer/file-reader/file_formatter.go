package filereader

import (
	"consumer-importer/model"
	"consumer-importer/util"
	"errors"
	"strconv"
	"strings"
	"time"
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

func ToConsumer(row util.FileRow) model.Consumer {
	private, err := parseBoolean(row.Private)

	if err != nil {
		panic(err)
	}

	incompleto, err := parseBoolean(row.Incompleto)

	if err != nil {
		panic(err)
	}

	return model.Consumer{
		CPF:                parseString(row.CPF),
		Private:            private,
		Incompleto:         incompleto,
		DataUltimaCompra:   parseDate(row.DataUltimaCompra),
		TicketMedio:        parseFloat(row.TicketMedio),
		TicketUltimaCompra: parseFloat(row.TicketUltimaCompra),
		LojaMaisFrequente:  parseString(row.LojaMaisFrequente),
		LojaUltimaCompra:   parseString(row.LojaUltimaCompra),
	}
}

func parseString(str string) *string {
	cleanedStr := util.CleanUpString(str)

	if strings.ToLower(cleanedStr) == "null" {
		return nil
	}

	return &cleanedStr
}

func parseBoolean(str string) (bool, error) {
	cleanedStr := util.CleanUpString(str)

	if cleanedStr == "1" {
		return true, nil
	} else if cleanedStr == "0" {
		return false, nil
	}

	return false, errors.New("invalid boolean type")
}

func parseDate(str string) time.Time {
	cleanedStr := util.CleanUpString(str)

	if strings.ToLower(cleanedStr) == "null" {
		return time.Time{}
	}

	time, _ := time.Parse("2006-01-02", cleanedStr)

	return time
}

func parseFloat(str string) *float64 {
	floatStr := strings.ReplaceAll(util.CleanUpString(str), ",", ".")

	if strings.ToLower(floatStr) == "null" {
		return nil
	}

	converted, _ := strconv.ParseFloat(floatStr, 64)

	return &converted
}
