package handy

import (
	"regexp"
	"strings"
	"unicode"
)

const (
	CheckStrAllowEmpty          = 1
	CheckStrDenySpaces          = 2
	CheckStrDenyDigits          = 4
	CheckStrDenyLetters         = 8
	CheckStrDenySymbols         = 16
	CheckStrDenyMoreThanOneWord = 32
	CheckStrDenyUpperCase       = 64
	CheckStrDenyLowercase       = 128
	CheckStrDenyUnicode         = 256

	CheckStrRequireDigits          = 512
	CheckStrRequireLetters         = 1024
	CheckStrRequireSymbols         = 2048
	CheckStrRequireMoreThanOneWord = 5096
	CheckStrRequireUpperCase       = 10192
	CheckStrRequireLowercase       = 20384

	CheckStrOk                    = 0
	CheckStrEmptyDenied           = -1
	CheckStrTooShort              = -2
	CheckStrTooLong               = -4
	CheckStrSpaceDenied           = -5
	CheckStrDigitsDenied          = -6
	CheckStrLettersDenied         = -7
	CheckStrSymbolsDenied         = -8
	CheckStrMoreThanOneWordDenied = -9
	CheckStrUpperCaseDenied       = -10
	CheckStrLowercaseDenied       = -11
	CheckStrUnicodeDenied         = -12

	CheckStrDigitsNotFound          = -13
	CheckStrLettersNotFound         = -14
	CheckStrSymbolsNotFound         = -15
	CheckStrMoreThanOneWordNotFound = -16
	CheckStrUpperCaseNotFound       = -17
	CheckStrLowercaseNotFound       = -18
)

var (
	reEmailFinder = regexp.MustCompile(`/([a-zA-Z0-9._-]+@[a-zA-Z0-9._-]+\.[a-zA-Z0-9_-]+)/gi`)
)

// CheckStr validates a string according given complexity rules
// CheckStr first evaluates "Deny" rules, and then "Require" rules.
// minLen=0 means there's no minimum lenght
// maxLen=0 means there's no maximum lenght
func CheckStr(seq string, minLen, maxLen int, rules uint64) int8 {
	if rules == 0 {
		return CheckStrOk
	}

	if seq == "" {
		if rules&CheckStrAllowEmpty == CheckStrAllowEmpty {
			return CheckStrOk
		}

		return CheckStrEmptyDenied
	}

	if len(seq) < minLen {
		return CheckStrTooShort
	}

	if maxLen > 0 && len(seq) > maxLen {
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

	if rules&CheckStrDenyDigits == CheckStrDenyDigits {
		if containsNumbers {
			return CheckStrDigitsDenied
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
			if unicode.IsSymbol(s) {
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

	containsMoreThanOneWord := len(strings.Fields(seq)) > 0

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
				if unicode.IsUpper(s) {
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

	if rules&CheckStrRequireDigits == CheckStrRequireDigits {
		if !containsNumbers {
			return CheckStrDigitsNotFound
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

const (
	StrFindNoneFound  = 0
	StrFindEmailFound = -1
)

// StrFind checks given string to find contained/hidden substrings
func StrFind(seq string, substr uint64) int8 {

	if reEmailFinder.MatchString(seq) {
		return StrFindEmailFound
	}

	return StrFindNoneFound
}
