package handy

import (
	"strings"
	"unicode/utf8"
)

const (
	// TransformNone No transformations are ordered. Only constraints maximum length
	// TransformNone turns all other flags OFF.
	TransformNone = 1
	// TransformFlagTrim Trim spaces before and after process the input
	// TransformFlagTrim Trims the string, removing leading and trailing spaces
	TransformFlagTrim = 2
	// TransformFlagLowerCase Makes the string lowercase
	// If case transformation flags are combined, the last one remains, considering the following order: TransformFlagTitleCase, TransformFlagLowerCase and TransformFlagUpperCase.
	TransformFlagLowerCase = 4
	// TransformFlagUpperCase Makes the string uppercase
	// If case transformation flags are combined, the last one remains, considering the following order: TransformFlagTitleCase, TransformFlagLowerCase and TransformFlagUpperCase.
	TransformFlagUpperCase = 8
	// TransformFlagOnlyDigits Removes all non-numeric characters
	TransformFlagOnlyDigits = 16
	// TransformFlagOnlyLetters Removes all non-letter characters
	TransformFlagOnlyLetters = 32
	// TransformFlagOnlyLettersAndDigits Leaves only letters and numbers
	TransformFlagOnlyLettersAndDigits = 64
	// TransformFlagHash After process all other flags, applies SHA256 hashing on string for output
	// 	The routine applies handy.StringHash() on given string
	TransformFlagHash = 128
	// TransformFlagTitleCase Makes the string uppercase
	// If case transformation flags are combined, the last one remains, considering the following order: TransformFlagTitleCase, TransformFlagLowerCase and TransformFlagUpperCase.
	TransformFlagTitleCase = 256
	// TransformFlagRemoveDigits Removes all digit characters, without to touch on any other
	// If combined with TransformFlagOnlyLettersAndDigits, TransformFlagOnlyDigits or TransformFlagOnlyLetters, it's ineffective
	TransformFlagRemoveDigits = 512
)

// Transform handles a string according given flags/parametrization, as follows:
// The transformations are made in arbitrary order, what can result in unexpected output. It the input matters, use TransformSerially instead.
// If maxLen==0, truncation is skipped
// The last operations are, by order, truncation and trimming.
func Transform(s string, maxLen int, transformFlags uint) string {
	if s == "" {
		return s
	}

	if transformFlags&TransformNone == TransformNone {
		if maxLen > 0 && utf8.RuneCountInString(s) > maxLen {
			s = string([]rune(s)[:maxLen])
		}

		return s
	}

	if (transformFlags & TransformFlagOnlyLettersAndDigits) == TransformFlagOnlyLettersAndDigits {
		s = OnlyLettersAndNumbers(s)
	}

	if (transformFlags & TransformFlagOnlyDigits) == TransformFlagOnlyDigits {
		s = OnlyDigits(s)
	}

	if (transformFlags & TransformFlagOnlyLetters) == TransformFlagOnlyLetters {
		s = OnlyLetters(s)
	}

	if (transformFlags & TransformFlagRemoveDigits) == TransformFlagRemoveDigits {
		s = RemoveDigits(s)
	}

	// Have to trim before and after, to avoid issues with string truncation and new leading/trailing spaces
	if (transformFlags & TransformFlagTrim) == TransformFlagTrim {
		s = strings.TrimSpace(s)
	}

	if (transformFlags & TransformFlagTitleCase) == TransformFlagTitleCase {
		s = strings.Title(strings.ToLower(s))
	}

	if (transformFlags & TransformFlagLowerCase) == TransformFlagLowerCase {
		s = strings.ToLower(s)
	}

	if (transformFlags & TransformFlagUpperCase) == TransformFlagUpperCase {
		s = strings.ToUpper(s)
	}

	if (transformFlags & TransformFlagHash) == TransformFlagHash {
		s = StringHash(s)
	}

	if s == "" {
		return s
	}

	if maxLen > 0 && utf8.RuneCountInString(s) > maxLen {
		s = string([]rune(s)[:maxLen])
	}

	// Have to trim before and after, to avoid issues with string truncation and new leading/trailing spaces
	if (transformFlags & TransformFlagTrim) == TransformFlagTrim {
		s = strings.TrimSpace(s)
	}

	return s
}

// TransformSerially reformat given string according parameters, in the order these params were sent
// Example: TransformSerially("uh lalah 123", 4, TransformFlagOnlyDigits,TransformFlagHash,TransformFlagUpperCase)
//          First remove non-digits, then hashes string and after make it all uppercase.
// If maxLen==0, truncation is skipped
// Truncation is the last operation
func TransformSerially(s string, maxLen int, transformFlags ...uint) string {
	if s == "" {
		return s
	}

	for _, flag := range transformFlags {
		switch flag {
		case TransformFlagOnlyLettersAndDigits:
			s = OnlyLettersAndNumbers(s)
		case TransformFlagOnlyDigits:
			s = OnlyDigits(s)
		case TransformFlagOnlyLetters:
			s = OnlyLetters(s)
		case TransformFlagTrim:
			s = strings.TrimSpace(s)
		case TransformFlagTitleCase:
			s = strings.ToTitle(s)
		case TransformFlagLowerCase:
			s = strings.ToLower(s)
		case TransformFlagUpperCase:
			s = strings.ToUpper(s)
		case TransformFlagHash:
			s = StringHash(s)
		}
	}

	if maxLen > 0 && utf8.RuneCountInString(s) > maxLen {
		s = string([]rune(s)[:maxLen])
	}

	return s
}
