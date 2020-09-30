package handy

import (
	"strconv"
	"strings"
)

// StringAsFloat tries to convert a string to float, and if it can't, just returns zero
// It's limited to one billion
func StringAsFloat(s string, decimalSeparator, thousandsSeparator rune) float64 {
	if s == "" {
		return 0.0
	}

	const maxLength = 20

	if len([]rune(s)) > maxLength {
		s = s[0:maxLength]
	}

	s = strings.Replace(s, string(thousandsSeparator), "", -1)

	s = strings.Replace(s, string(decimalSeparator), ".", -1)

	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f
	}

	return 0.0
}

// StringAsInteger returns the integer value extracted from string, or zero
func StringAsInteger(s string) int {
	if s == "" {
		return 0
	}

	if i, err := strconv.ParseInt(s, 10, 32); err == nil {
		return int(i)
	}

	return 0
}
