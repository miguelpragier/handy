package handy

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	CheckStrAllowEmpty          = 1
	CheckStrDenySpaces          = 2
	CheckStrDenyNumbers         = 4
	CheckStrDenyLetters         = 8
	CheckStrDenySymbols         = 16
	CheckStrDenyMoreThanOneWord = 32
	CheckStrDenyUpperCase       = 64
	CheckStrDenyLowercase       = 128
	CheckStrDenyUnicode         = 256

	CheckStrRequireNumbers         = 512
	CheckStrRequireLetters         = 1024
	CheckStrRequireSymbols         = 2048
	CheckStrRequireMoreThanOneWord = 4096
	CheckStrRequireUpperCase       = 8192
	CheckStrRequireLowercase       = 16384

	CheckStrOk                    = 0
	CheckStrEmptyDenied           = -1
	CheckStrTooShort              = -2
	CheckStrTooLong               = -4
	CheckStrSpaceDenied           = -5
	CheckStrNubersDenied          = -6
	CheckStrLettersDenied         = -7
	CheckStrSymbolsDenied         = -8
	CheckStrMoreThanOneWordDenied = -9
	CheckStrUpperCaseDenied       = -10
	CheckStrLowercaseDenied       = -11
	CheckStrUnicodeDenied         = -12

	CheckStrNumbersNotFound         = -13
	CheckStrLettersNotFound         = -14
	CheckStrSymbolsNotFound         = -15
	CheckStrMoreThanOneWordNotFound = -16
	CheckStrUpperCaseNotFound       = -17
	CheckStrLowercaseNotFound       = -18
)

var (
	reEmailFinder = regexp.MustCompile(`([a-zA-Z0-9._-]+@[a-zA-Z0-9._-]+\.[a-zA-Z0-9_-]+)`)
)

// CheckStr validates a string according given complexity rules
// CheckStr first evaluates "Deny" rules, and then "Require" rules.
// minLen=0 means there's no minimum lenght
// maxLen=0 means there's no maximum lenght
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

	containsNumbers := func(s string) bool {
		for _, s := range s {
			if unicode.IsNumber(s) {
				return true
			}
		}

		return false
	}(seq)

	if rules&CheckStrDenyNumbers == CheckStrDenyNumbers {
		if containsNumbers {
			return CheckStrNubersDenied
		}
	}

	containsLetters := func(s string) bool {
		for _, s := range s {
			if unicode.IsLetter(s) {
				return true
			}
		}

		return false
	}(seq)

	if rules&CheckStrDenyLetters == CheckStrDenyLetters {
		if containsLetters {
			return CheckStrLettersDenied
		}
	}

	containsSymbols := func(s string) bool {
		for _, s := range s {
			if unicode.IsSymbol(s) || ( !unicode.IsLetter(s) && !unicode.IsNumber(s) && !unicode.IsSpace(s) ) {
				return true
			}
		}

		return false
	}(seq)

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
