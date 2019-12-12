package handy

import (
	"log"
	"math/rand"
	"unicode"
	"unicode/utf8"
)

// RandomString generates a string sequence based on given params/rules
func RandomString(minLen, maxLen int, allowUnicode, allowNumbers, allowSymbols, allowSpaces bool) string {
	if minLen > maxLen {
		log.Println("handy.RandomString(minLen is greater than maxLen)")
		return ""
	}

	if maxLen == 0 {
		log.Println("handy.RandomString(maxLen should be greater than zero)")
		return ""
	}

	if minLen == 0 {
		minLen = 1
	}

	strLen := rand.Intn(maxLen-1) + 1

	if minLen < maxLen {
		strLen = rand.Intn(maxLen-minLen) + minLen
	}

	str := make([]rune, strLen)

	maxRune := utf8.MaxRune

	if !allowUnicode {
		maxRune = 127
	}

	// Checks if the space is at beggining or at string end
	// to avoid leading or trailing spaces
	firstOrLast := func(i int) bool {
		return i == 0 || i == strLen-1
	}

	for i := 0; i < strLen; {
		r := rand.Int31n(maxRune-1) + 1

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
