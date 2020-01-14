package handy

import (
	"log"
	"math/rand"
	"unicode"
	"unicode/utf8"
)

// RandomString generates a string sequence based on given params/rules
func RandomString(minLen, maxLen int, allowUnicode, allowNumbers, allowSymbols, allowSpaces bool) string {
	switch {
	case minLen > maxLen:
		log.Println("handy.RandomString(minLen is greater than maxLen)")
		return ""

	case maxLen == 0:
		log.Println("handy.RandomString(maxLen should be greater than zero)")
		return ""

	case minLen == 0:
		minLen = 1
	}

	// If minLen==maxLen, force fixed size string
	strLen := maxLen

	// but if minLen<>maxLen, string length must be between minLen and maxLen
	if minLen < maxLen {
		strLen = rand.Intn(maxLen-minLen) + minLen
	}

	str := make([]rune, strLen)

	// Checks if the space is at beggining or at string end
	// to avoid leading or trailing spaces
	firstOrLast := func(i int) bool {
		return i == 0 || i == strLen-1
	}

	const minimumPrintableRune = 32

	var (
		asciiTable    = []rune(" !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~")
		asciiTableLen = int32(len(asciiTable))
		r             rune
	)

	for i := 0; i < strLen; {
		if !allowUnicode {
			p := rand.Int31n(asciiTableLen)
			r = asciiTable[p]
		} else {
			r = rand.Int31n(utf8.MaxRune-minimumPrintableRune) + minimumPrintableRune
		}

		switch {
		case !unicode.IsPrint(r):
			continue

		case unicode.IsLetter(r):
			str[i] = r
			i++

		case unicode.IsNumber(r) || unicode.IsDigit(r):
			if allowNumbers {
				str[i] = r
				i++
			}

		case unicode.IsSymbol(r) || unicode.IsPunct(r):
			if allowSymbols {
				str[i] = r
				i++
			}

		case unicode.IsSpace(r):
			if allowSpaces && !firstOrLast(i) {
				str[i] = r
				i++
			}
		}
	}

	return string(str)
}
