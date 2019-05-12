package util

import (
	"testing"
)

func TestRemoveDiatrics(t *testing.T) {
	provided := "áàâã.éèê/ìíî-õôóò:ùúû-"
	expected := "aaaa.eee/iii-oooo:uuu-"

	result := removeDiatrics(provided)

	if result != expected {
		t.Errorf("Diatrics have not been removed: " + result)
	}
}

func TestTrim(t *testing.T) {
	provided := "ao      "
	expected := "ao"

	result := CleanUpString(provided)

	if result != expected {
		t.Errorf("The string still has blank spaces: " + result)
	}
}

func TestUpperCase(t *testing.T) {
	provided := "ABCDEFGHIJKLMNOPQRSTUVXYZ"
	expected := "abcdefghijklmnopqrstuvxyz"

	result := CleanUpString(provided)

	if result != expected {
		t.Errorf("The string still has uppercase letters: " + result)
	}
}
