package filereader

import (
	"fmt"
	"testing"
	"time"
)

func TestWithNullSampleType(t *testing.T) {
	sample := "041.091.641-25     0           0           NULL                  NULL                  NULL                    NULL                NULL"
	row := FormatFile(sample)

	ValidateAttribute("041.091.641-25     ", row.CPF, "CPF", t)
	ValidateAttribute("0           ", row.Private, "Private", t)
	ValidateAttribute("0           ", row.Incompleto, "Incompleto", t)
	ValidateAttribute("NULL                  ", row.DataUltimaCompra, "DataUtlimaCompra", t)
	ValidateAttribute("NULL                  ", row.TicketMedio, "TicketMedio", t)
	ValidateAttribute("NULL                    ", row.TicketUltimaCompra, "TicketUltimaCompra", t)
	ValidateAttribute("NULL                ", row.LojaMaisFrequente, "LojaMaisFrequente", t)
	ValidateAttribute("NULL", row.LojaUltimaCompra, "LojaUltimaCompra", t)
}

func TestWithFilledSampleType(t *testing.T) {
	sample := "866.315.609-00     0           0           2010-01-13            335,38                335,38                  79.379.491/0001-83  79.379.491/0001-83"
	row := FormatFile(sample)

	ValidateAttribute("866.315.609-00     ", row.CPF, "CPF", t)
	ValidateAttribute("0           ", row.Private, "Private", t)
	ValidateAttribute("0           ", row.Incompleto, "Incompleto", t)
	ValidateAttribute("2010-01-13            ", row.DataUltimaCompra, "DataUtlimaCompra", t)
	ValidateAttribute("335,38                ", row.TicketMedio, "TicketMedio", t)
	ValidateAttribute("335,38                  ", row.TicketUltimaCompra, "TicketUltimaCompra", t)
	ValidateAttribute("79.379.491/0001-83  ", row.LojaMaisFrequente, "LojaMaisFrequente", t)
	ValidateAttribute("79.379.491/0001-83", row.LojaUltimaCompra, "LojaUltimaCompra", t)
}

func ValidateAttribute(expected, received, fieldName string, t *testing.T) {
	if received != expected {
		t.Errorf("unexpected " + fieldName + ", expected: '" + expected + "', received: '" + received + "'")
	}
}

func TestStringParserCPF(t *testing.T) {
	sample := "866.315.609-00     "
	expected := "866.315.609-00"
	result := parseString(sample)

	if *result != expected {
		t.Errorf("incorret parsing, expected: '" + expected + "', received: '" + *result + "'")
	}
}

func TestStringParserCNPJ(t *testing.T) {
	sample := "79.379.491/0001-83  "
	expected := "79.379.491/0001-83"
	result := parseString(sample)

	if *result != expected {
		t.Errorf("incorret parsing, expected: '" + expected + "', received: '" + *result + "'")
	}
}

func TestStringParserNull(t *testing.T) {
	sample := "NULL                "
	result := parseString(sample)

	if result != nil {
		t.Errorf("incorret parsing, expected: 'nil', received" + *result + "'")
	}
}

func TestBooleanParserFalse(t *testing.T) {
	sample := "0           "
	result, _ := parseBoolean(sample)

	if result {
		t.Errorf("incorret parsing, expected: 'false', received 'true'")
	}
}

func TestBooleanParserTrue(t *testing.T) {
	sample := "1           "
	result, _ := parseBoolean(sample)

	if !result {
		t.Errorf("incorret parsing, expected: 'true', received 'false'")
	}
}

func TestBooleanParserError(t *testing.T) {
	sample := "t           "
	_, err := parseBoolean(sample)

	if err == nil {
		t.Errorf("incorret parsing, expected err not nil")
	}
}

func TestStringParserDate(t *testing.T) {
	sample := "2010-01-13            "
	expected := time.Date(2010, 1, 13, 0, 0, 0, 0, time.UTC)
	result := parseDate(sample)

	if !result.Equal(expected) {
		t.Errorf("incorret parsing, expected: '" + expected.String() + "', received: '" + result.String() + "'")
	}
}
func TestStringParserDateNull(t *testing.T) {
	sample := "NULL                "
	expected := time.Time{}
	result := parseDate(sample)

	if !result.Equal(expected) {
		t.Errorf("incorret parsing, expected: '" + expected.String() + "', received: '" + result.String() + "'")
	}
}
func TestFloatParser(t *testing.T) {
	sample := "335,38                "
	expected := 335.38
	result := parseFloat(sample)

	if *result != expected {
		t.Errorf("incorret parsing, expected: '" + fmt.Sprintf("%f", expected) + "', received: '" + fmt.Sprintf("%f", *result) + "'")
	}
}
func TestFloatParserNull(t *testing.T) {
	sample := "NULL                "
	result := parseFloat(sample)

	if result != nil {
		t.Errorf("incorret parsing, expected: 'nil', received: '" + fmt.Sprintf("%f", *result) + "'")
	}
}
