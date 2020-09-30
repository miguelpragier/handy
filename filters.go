package handy

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

var reDupSpaces = regexp.MustCompile(`\s+`)

// DedupSpaces removes duplicated spaces, tabs and newLine characters
// I.E: Replaces two tabs for one single tab
func DedupSpaces(s string) string {
	if s == "" {
		return ""
	}

	s = reDupSpaces.ReplaceAllString(s, " ")

	return s
}

// CleanSpaces removes duplicated spaces, tabs and newLine characters and then trim string's both sides
func CleanSpaces(s string) string {
	if s == "" {
		return ""
	}

	return strings.TrimSpace(DedupSpaces(s))
}

// OnlyLetters returns only the letters from the given string, after strip all the rest ( numbers, spaces, etc. )
func OnlyLetters(sequence string) string {
	if utf8.RuneCountInString(sequence) == 0 {
		return ""
	}

	var letters []rune

	for _, r := range sequence {
		if unicode.IsLetter(r) {
			letters = append(letters, r)
		}
	}

	return string(letters)
}

// OnlyDigits returns only the numbers from the given string, after strip all the rest ( letters, spaces, etc. )
func OnlyDigits(sequence string) string {
	if utf8.RuneCountInString(sequence) > 0 {
		re, _ := regexp.Compile(`[\D]`)

		sequence = re.ReplaceAllString(sequence, "")
	}

	return sequence
}

// OnlyLettersAndNumbers returns only the letters and numbers from the given string, after strip all the rest, like spaces and special symbols.
func OnlyLettersAndNumbers(sequence string) string {
	if utf8.RuneCountInString(sequence) == 0 {
		return ""
	}

	var alphanumeric []rune

	for _, r := range sequence {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			alphanumeric = append(alphanumeric, r)
		}
	}

	return string(alphanumeric)
}

// RemoveDigits returns the given string without digit/numeric runes
func RemoveDigits(sequence string) string {
	if utf8.RuneCountInString(sequence) == 0 {
		return ""
	}

	var rs []rune

	for _, r := range sequence {
		if !unicode.IsDigit(r) {
			rs = append(rs, r)
		}
	}

	return string(rs)
}
