package handy

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// HasNumber returns true if input string contains at least one digit/number
func HasNumber(s string) bool {
	for _, s := range s {
		if unicode.IsNumber(s) {
			return true
		}
	}

	return false
}

// HasLetter returns true if input string contains at least one letter
func HasLetter(s string) bool {
	for _, s := range s {
		if unicode.IsLetter(s) {
			return true
		}
	}

	return false
}

// HasSymbol returns true if input string contains at least one symbol
// If rune is not a space, letter nor a number, it's considered a symbol
func HasSymbol(s string) bool {
	for _, s := range s {
		if unicode.IsSymbol(s) || (!unicode.IsLetter(s) && !unicode.IsNumber(s) && !unicode.IsSpace(s)) {
			return true
		}
	}

	return false
}

const (
	// CheckStrAllowEmpty allows empty string ""
	CheckStrAllowEmpty = 1
	// CheckStrDenySpaces forbids spaces, tabs, new lines and carriage return
	CheckStrDenySpaces = 2
	// CheckStrDenyNumbers forbids digits/numbers
	CheckStrDenyNumbers = 4
	// CheckStrDenyLetters forbids letters
	CheckStrDenyLetters = 8
	// CheckStrDenySymbols forbids symbols. if it's not a number, letter or space, is considered a symbol
	CheckStrDenySymbols = 16
	// CheckStrDenyMoreThanOneWord forbids more than one word
	CheckStrDenyMoreThanOneWord = 32
	// CheckStrDenyUpperCase forbids uppercase letters
	CheckStrDenyUpperCase = 64
	// CheckStrDenyLowercase forbids lowercase letters
	CheckStrDenyLowercase = 128
	// CheckStrDenyUnicode forbids non-ASCII characters
	CheckStrDenyUnicode = 256
	// CheckStrRequireNumbers demands at least 1 number within string
	CheckStrRequireNumbers = 512
	// CheckStrRequireLetters demands at least 1 letter within string
	CheckStrRequireLetters = 1024
	// CheckStrRequireSymbols demands at least 1 symbol within string. if it's not a number, letter or space, is considered a symbol
	CheckStrRequireSymbols = 2048
	// CheckStrRequireMoreThanOneWord demands at least 2 words in given string input
	CheckStrRequireMoreThanOneWord = 4096
	// CheckStrRequireUpperCase demands at least 1 uppercase letter within string
	CheckStrRequireUpperCase = 8192
	// CheckStrRequireLowercase demands at least 1 lowercase letter within string
	CheckStrRequireLowercase = 16384

	// CheckStrOk means "alright"
	CheckStrOk = 0
	// CheckStrEmptyDenied is self explained
	CheckStrEmptyDenied = -1
	// CheckStrTooShort is self explained
	CheckStrTooShort = -2
	// CheckStrTooLong is self explained
	CheckStrTooLong = -4
	// CheckStrSpaceDenied is self explained
	CheckStrSpaceDenied = -5
	// CheckStrNubersDenied is self explained
	CheckStrNubersDenied = -6
	// CheckStrLettersDenied is self explained
	CheckStrLettersDenied = -7
	// CheckStrSymbolsDenied is self explained
	CheckStrSymbolsDenied = -8
	// CheckStrMoreThanOneWordDenied is self explained
	CheckStrMoreThanOneWordDenied = -9
	// CheckStrUpperCaseDenied is self explained
	CheckStrUpperCaseDenied = -10
	// CheckStrLowercaseDenied is self explained
	CheckStrLowercaseDenied = -11
	// CheckStrUnicodeDenied is self explained
	CheckStrUnicodeDenied = -12

	// CheckStrNumbersNotFound is self explained
	CheckStrNumbersNotFound = -13
	// CheckStrLettersNotFound is self explained
	CheckStrLettersNotFound = -14
	// CheckStrSymbolsNotFound is self explained
	CheckStrSymbolsNotFound = -15
	// CheckStrMoreThanOneWordNotFound is self explained
	CheckStrMoreThanOneWordNotFound = -16
	// CheckStrUpperCaseNotFound is self explained
	CheckStrUpperCaseNotFound = -17
	// CheckStrLowercaseNotFound is self explained
	CheckStrLowercaseNotFound = -18
)

var (
	reEmailFinder = regexp.MustCompile(`([a-zA-Z0-9._-]+@[a-zA-Z0-9._-]+\.[a-zA-Z0-9_-]+)`)
)

// CheckStr validates a string according given complexity rules
// CheckStr first evaluates "Deny" rules, and then "Require" rules.
// minLen=0 means there's no minimum length
// maxLen=0 means there's no maximum length
func CheckStr(seq string, minLen, maxLen uint, rules uint64) int8 {
	strLen := uint(utf8.RuneCountInString(seq))

	if seq == "" {
		if rules&CheckStrAllowEmpty == CheckStrAllowEmpty {
			return CheckStrOk
		}

		return CheckStrEmptyDenied
	}

	if strLen < minLen {
		return CheckStrTooShort
	}

	if maxLen > 0 && strLen > maxLen {
		return CheckStrTooLong
	}

	if rules&CheckStrDenySpaces == CheckStrDenySpaces {
		if strings.ContainsAny(seq, "\n\r\t ") {
			return CheckStrSpaceDenied
		}
	}

	containsNumbers := HasNumber(seq)

	if rules&CheckStrDenyNumbers == CheckStrDenyNumbers {
		if containsNumbers {
			return CheckStrNubersDenied
		}
	}

	containsLetters := HasLetter(seq)

	if rules&CheckStrDenyLetters == CheckStrDenyLetters {
		if containsLetters {
			return CheckStrLettersDenied
		}
	}

	containsSymbols := HasSymbol(seq)

	if rules&CheckStrDenySymbols == CheckStrDenySymbols {
		if containsSymbols {
			return CheckStrSymbolsDenied
		}
	}

	containsMoreThanOneWord := len(strings.Fields(seq)) > 1

	if rules&CheckStrDenyMoreThanOneWord == CheckStrDenyMoreThanOneWord {
		if containsMoreThanOneWord {
			return CheckStrMoreThanOneWordDenied
		}
	}

	containsUppercase := func(s string) bool {
		if containsLetters {
			for _, s := range s {
				if unicode.IsUpper(s) {
					return true
				}
			}
		}

		return false
	}(seq)

	if rules&CheckStrDenyUpperCase == CheckStrDenyUpperCase {
		if containsUppercase {
			return CheckStrUpperCaseDenied
		}
	}

	containsLowercase := func(s string) bool {
		if containsLetters {
			for _, s := range s {
				if unicode.IsLower(s) {
					return true
				}
			}
		}

		return false
	}(seq)

	if rules&CheckStrDenyLowercase == CheckStrDenyLowercase {
		if containsLowercase {
			return CheckStrLowercaseDenied
		}
	}

	if rules&CheckStrDenyUnicode == CheckStrDenyUnicode {
		for _, s := range seq {
			if s > unicode.MaxASCII {
				return CheckStrUnicodeDenied
			}
		}
	}

	if rules&CheckStrRequireNumbers == CheckStrRequireNumbers {
		if !containsNumbers {
			return CheckStrNumbersNotFound
		}
	}

	if rules&CheckStrRequireLetters == CheckStrRequireLetters {
		if !containsLetters {
			return CheckStrLettersNotFound
		}
	}

	if rules&CheckStrRequireSymbols == CheckStrRequireSymbols {
		if !containsSymbols {
			return CheckStrSymbolsNotFound
		}
	}

	if rules&CheckStrRequireMoreThanOneWord == CheckStrRequireMoreThanOneWord {
		if !containsMoreThanOneWord {
			return CheckStrMoreThanOneWordNotFound
		}
	}

	if rules&CheckStrRequireUpperCase == CheckStrRequireUpperCase {
		if !containsUppercase {
			return CheckStrUpperCaseNotFound
		}
	}

	if rules&CheckStrRequireLowercase == CheckStrRequireLowercase {
		if !containsLowercase {
			return CheckStrLowercaseNotFound
		}
	}

	return CheckStrOk
}

// StrContainsEmail returns true if given string contains an email address
func StrContainsEmail(seq string) bool {

	return reEmailFinder.MatchString(seq)
}
