package validator

import "testing"

func TestCPFValidation(t *testing.T) {
	sample := "04109164125"

	if !ValidateCpf(sample) {
		t.Errorf("should've validated the CPF: " + sample)
	}
}

func TestCPFValidationWithMoreDigits(t *testing.T) {
	sample := "041091641256"

	if ValidateCpf(sample) {
		t.Errorf("shouldn't have validated the CPF: " + sample)
	}
}

func TestCPFValidationWithInvalidCPF(t *testing.T) {
	sample := "43738288821"

	if ValidateCpf(sample) {
		t.Errorf("shouldn't have validated the CPF: " + sample)
	}
}

func TestCNPJValidation(t *testing.T) {
	sample := "79379491000183"

	if !ValidateCNPJ(sample) {
		t.Errorf("should've validated the CPF: " + sample)
	}
}

func TestCNPJValidationWithMoreDigits(t *testing.T) {
	sample := "041091641256"

	if ValidateCNPJ(sample) {
		t.Errorf("shouldn't have validated the CPF: " + sample)
	}
}

func TestCNPJValidationWithInvalidCPF(t *testing.T) {
	sample := "153678436555421"

	if ValidateCNPJ(sample) {
		t.Errorf("shouldn't have validated the CPF: " + sample)
	}
}
