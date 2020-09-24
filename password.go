package handy

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	// CheckNewPassword() Possible results

	// CheckNewPasswordResultOK Means the checking ran alright
	CheckNewPasswordResultOK = 0
	// CheckNewPasswordResultDivergent Password is different from confirmation
	CheckNewPasswordResultDivergent = 1
	// CheckNewPasswordResultTooShort Password is too short
	CheckNewPasswordResultTooShort = 2
	// CheckNewPasswordResultTooSimple Given string doesn't satisfy complexity rules
	CheckNewPasswordResultTooSimple = 3

	// CheckNewPassword() Complexity Rules

	// CheckNewPasswordComplexityLowest There's no rules besides the minimum length
	// >>> This flag turns all others off <<<
	CheckNewPasswordComplexityLowest = 1
	// CheckNewPasswordComplexityRequireLetter At least one letter is required in order to aprove password
	CheckNewPasswordComplexityRequireLetter = 2
	// CheckNewPasswordComplexityRequireUpperCase At least one uppercase letter is required in order to aprove password.
	// Only works if CheckNewPasswordComplexityRequireLetter is included/activated
	CheckNewPasswordComplexityRequireUpperCase = 4
	// CheckNewPasswordComplexityRequireNumber At least one number is required in order to aprove password
	CheckNewPasswordComplexityRequireNumber = 8
	// CheckNewPasswordComplexityRequireSpace The password must contain at least one space
	CheckNewPasswordComplexityRequireSpace = 16
	// CheckNewPasswordComplexityRequireSymbol User have to include at least one special character, like # or -
	CheckNewPasswordComplexityRequireSymbol = 32
)

// CheckNewPassword Run some basic checks on new password strings, based on given options
// This routine requires at least 4 (four) characters
// Example requiring only basic minimum length: CheckNewPassword("lalala", "lalala", 10, CheckNewPasswordComplexityLowest)
// Example requiring number and symbol: CheckNewPassword("lalala", "lalala", 10, CheckNewPasswordComplexityRequireNumber|CheckNewPasswordComplexityRequireSymbol)
func CheckNewPassword(password, passwordConfirmation string, minimumlength uint, flagComplexity uint8) uint8 {
	const minPasswordLengthDefault = 4

	if minimumlength < minPasswordLengthDefault {
		minimumlength = 4
	}

	if utf8.RuneCountInString(strings.TrimSpace(password)) < int(minimumlength) {
		return CheckNewPasswordResultTooShort
	}

	if password != passwordConfirmation {
		return CheckNewPasswordResultDivergent
	}

	letterFound := false
	numberFound := false
	symbolFound := false
	spaceFound := false
	upperCaseFound := false

	if flagComplexity&CheckNewPasswordComplexityLowest != CheckNewPasswordComplexityLowest {
		for _, r := range password {
			if unicode.IsLetter(r) {
				letterFound = true

				if unicode.IsUpper(r) {
					upperCaseFound = true
				}
			}

			if unicode.IsNumber(r) {
				numberFound = true
			}

			if RuneHasSymbol(r) {
				symbolFound = true
			}

			if r == ' ' {
				spaceFound = true
			}
		}
	}

	if flagComplexity&CheckNewPasswordComplexityRequireLetter == CheckNewPasswordComplexityRequireLetter {
		if !letterFound {
			return CheckNewPasswordResultTooSimple
		}

		// Only checks uppercase if letter is required
		if flagComplexity&CheckNewPasswordComplexityRequireUpperCase == CheckNewPasswordComplexityRequireUpperCase {
			if !upperCaseFound {
				return CheckNewPasswordResultTooSimple
			}
		}
	}

	if flagComplexity&CheckNewPasswordComplexityRequireNumber == CheckNewPasswordComplexityRequireNumber {
		if !numberFound {
			return CheckNewPasswordResultTooSimple
		}
	}

	if flagComplexity&CheckNewPasswordComplexityRequireSymbol == CheckNewPasswordComplexityRequireSymbol {
		if !symbolFound {
			return CheckNewPasswordResultTooSimple
		}
	}

	if flagComplexity&CheckNewPasswordComplexityRequireSpace == CheckNewPasswordComplexityRequireSpace {
		if !spaceFound {
			return CheckNewPasswordResultTooSimple
		}
	}

	return CheckNewPasswordResultOK
}
