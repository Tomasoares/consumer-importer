package util

import (
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func isNonspacingMarks(r rune) bool {
	return unicode.Is(unicode.Mn, r)
}

func removeDiatrics(str string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isNonspacingMarks), norm.NFC)
	result, _, _ := transform.String(t, str)
	return result
}

//CleanUpString remove diatrics, turn into lower case and remove blank spaces from a string
func CleanUpString(str string) string {
	noDiatrics := removeDiatrics(str)
	noDiatricsLowerCase := strings.ToLower(noDiatrics)
	return strings.Trim(noDiatricsLowerCase, " ")
}
