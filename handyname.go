package handy

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	// CheckPersonNameResultOK means the name was validated
	CheckPersonNameResultOK = 0
	// CheckPersonNameResultPolluted The routine only accepts letters, single quotes and spaces
	CheckPersonNameResultPolluted = 1
	// CheckPersonNameResultTooFewWords The funcion requires at least 2 words
	CheckPersonNameResultTooFewWords = 2
	// CheckPersonNameResultTooShort the sum of all characters must be >= 6
	CheckPersonNameResultTooShort = 3
	// CheckPersonNameResultTooSimple The name rule requires that at least one word
	CheckPersonNameResultTooSimple = 4
)

// CheckPersonName returns true if the name contains at least two words, one >= 3 chars and one >=2 chars.
// I understand that this is a particular criteria, but this is the OpenSourceMagic, where you can change and adapt to your own specs.
func CheckPersonName(name string, acceptEmpty bool) uint8 {
	name = strings.TrimSpace(name)

	// If name is empty, AND it's accepted, return ok. Else, cry!
	if name == "" {
		if !acceptEmpty {
			return CheckPersonNameResultTooShort
		}

		return CheckPersonNameResultOK
	}

	// Person names doesn't accept other than letters, spaces and single quotes
	for _, r := range name {
		if !unicode.IsLetter(r) && r != ' ' && r != '\'' && r != '-' {
			return CheckPersonNameResultPolluted
		}
	}

	// A complete name has to be at least 2 words.
	a := strings.Fields(name)

	if len(a) < 2 {
		return CheckPersonNameResultTooFewWords
	}

	// At least two words, one with 3 chars and other with 2
	found2 := false
	found3 := false

	for _, s := range a {
		if !found3 && utf8.RuneCountInString(s) >= 3 {
			found3 = true
			continue
		}

		if !found2 && utf8.RuneCountInString(s) >= 2 {
			found2 = true
			continue
		}
	}

	if !found2 || !found3 {
		return CheckPersonNameResultTooSimple
	}

	return CheckPersonNameResultOK
}

// NameFirstAndLast returns the first and last words/names from the given input, optionally transformed by transformFlags
// Example: handy.NameFirstAndLast("friedrich wilhelm nietzsche", handy.TransformFlagTitleCase) // returns "Friedrich Nietzsche"
func NameFirstAndLast(name string, transformFlags uint) string {
	name = strings.Replace(name, "\t", ` `, -1)

	if transformFlags != TransformNone {
		name = Transform(name, utf8.RuneCountInString(name), transformFlags)
	}

	name = strings.TrimSpace(name)

	if name == `` {
		return ``
	}

	words := strings.Split(name, ` `)

	wl := len(words)

	if wl <= 0 {
		return ``
	}

	if wl == 1 {
		return words[0]
	}

	return fmt.Sprintf(`%s %s`, words[0], words[wl-1])
}

// NameInitials returns the first and last words/names from the given input, optionally transformed by transformFlags
func NameInitials(name string, transformFlags uint) string {
	name = strings.TrimSpace(name)

	name = strings.Replace(name, "\t", ` `, -1)

	if name == "" {
		return ``
	}

	if transformFlags != TransformNone {
		name = Transform(name, utf8.RuneCountInString(name), transformFlags)
	}

	words := strings.Split(name, ` `)

	wl := len(words)

	if wl <= 0 {
		return ``
	}

	var a []string

	for _, s := range words {
		for _, letter := range s {
			a = append(a, string(letter))
			break
		}
	}

	return strings.Join(a, ` `)
}
