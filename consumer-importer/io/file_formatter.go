package io

import (
	"consumer-importer/io/util"
	"consumer-importer/model"
	"errors"
	"strconv"
	"strings"
	"time"
)

//IsHeader checks if the line is a header type
func IsHeader(line string) bool {
	return len(line) == util.HeaderSize
}

//Format format a string line of the imported file to a Consumer object
func Format(line string) model.Consumer {
	row := formatFile(line)
	return toConsumer(row)
}

func formatFile(row string) util.FileRow {
	return util.FileRow{
		CPF:                row[:util.CPFEnd],
		Private:            row[util.CPFEnd:util.PrivateEnd],
		Incompleto:         row[util.PrivateEnd:util.IncompletoEnd],
		DataUltimaCompra:   row[util.IncompletoEnd:util.DataUltimaCompraEnd],
		TicketMedio:        row[util.DataUltimaCompraEnd:util.TicketMedioEnd],
		TicketUltimaCompra: row[util.TicketMedioEnd:util.TicketUltimaCompraEnd],
		LojaMaisFrequente:  row[util.TicketUltimaCompraEnd:util.LojaMaisFrequenteEnd],
		LojaUltimaCompra:   row[util.LojaMaisFrequenteEnd:],
	}
}

func toConsumer(row util.FileRow) model.Consumer {
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

	removedDots := strings.ReplaceAll(cleanedStr, ".", "")
	removedDashes := strings.ReplaceAll(removedDots, "-", "")
	finalStr := strings.ReplaceAll(removedDashes, "/", "")

	return &finalStr
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
