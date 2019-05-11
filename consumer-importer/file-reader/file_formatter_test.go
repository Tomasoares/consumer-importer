package filereader

import (
	"testing"
)

func TestFileFormatterWithNullSampleType(t *testing.T) {
	sample := "041.091.641-25     0           0           NULL                  NULL                  NULL                    NULL                NULL"
	row := FormatFile(sample)

	ValidateAttribute("041.091.641-25     "		, row.CPF				, "CPF", t)
	ValidateAttribute("0           "			, row.Private			, "Private", t)
	ValidateAttribute("0           "			, row.Incompleto		, "Incompleto", t)
	ValidateAttribute("NULL                  "	, row.DataUltimaCompra	, "DataUtlimaCompra", t)
	ValidateAttribute("NULL                  "	, row.TicketMedio		, "TicketMedio", t)
	ValidateAttribute("NULL                    ", row.TicketUltimaCompra, "TicketUltimaCompra", t)
	ValidateAttribute("NULL                "	, row.LojaMaisFrequente	, "LojaMaisFrequente", t)
	ValidateAttribute("NULL"					, row.LojaUltimaCompra	, "LojaUltimaCompra", t)
}

func TestFileFormatterWithFiiledSampleType(t *testing.T) {
	sample := "866.315.609-00     0           0           2010-01-13            335,38                335,38                  79.379.491/0001-83  79.379.491/0001-83"
	row := FormatFile(sample)

	ValidateAttribute("866.315.609-00     "		, row.CPF				, "CPF", t)
	ValidateAttribute("0           "			, row.Private			, "Private", t)
	ValidateAttribute("0           "			, row.Incompleto		, "Incompleto", t)
	ValidateAttribute("2010-01-13            "	, row.DataUltimaCompra	, "DataUtlimaCompra", t)
	ValidateAttribute("335,38                "	, row.TicketMedio		, "TicketMedio", t)
	ValidateAttribute("335,38                  ", row.TicketUltimaCompra, "TicketUltimaCompra", t)
	ValidateAttribute("79.379.491/0001-83  "	, row.LojaMaisFrequente	, "LojaMaisFrequente", t)
	ValidateAttribute("79.379.491/0001-83"		, row.LojaUltimaCompra	, "LojaUltimaCompra", t)
}

func ValidateAttribute(expected, received, fieldName string, t *testing.T) {
	if received != expected {
		t.Errorf("unexpected "+fieldName+", expected: '"+expected+"', received: '" + received + "'")
	}
}