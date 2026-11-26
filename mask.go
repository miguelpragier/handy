package handy

import (
	"errors"
	"unicode/utf8"
)

// Mask creates a whole-string rule and panics if pattern is invalid.
// See CompileMask for the placeholder syntax.
func Mask(pattern string) Rule {
	rule, err := CompileMask(pattern)
	if err != nil {
		panic(err)
	}
	return rule
}

// CompileMask creates a whole-string rule that reports the code "mask".
// '#' matches an ASCII digit and '@' matches an ASCII letter. Other runes
// are literal. A backslash escapes '#', '@', or another backslash.
// An empty pattern matches an empty string.
func CompileMask(pattern string) (Rule, error) {
	if !utf8.ValidString(pattern) {
		return nil, errors.New("handy: mask must be valid UTF-8")
	}

	type token struct {
		char    rune
		literal bool
	}
	var tokens []token
	escaped := false
	for _, r := range pattern {
		if escaped {
			if r != '#' && r != '@' && r != '\\' {
				return nil, errors.New("handy: invalid mask escape")
			}
			tokens = append(tokens, token{char: r, literal: true})
			escaped = false
			continue
		}
		if r == '\\' {
			escaped = true
			continue
		}
		tokens = append(tokens, token{char: r, literal: r != '#' && r != '@'})
	}
	if escaped {
		return nil, errors.New("handy: unfinished mask escape")
	}

	return func(value string) []Violation {
		for _, token := range tokens {
			r, size := utf8.DecodeRuneInString(value)
			if size == 0 || (r == utf8.RuneError && size == 1) {
				return failure("mask", false)
			}
			matches := r == token.char
			if !token.literal {
				switch token.char {
				case '#':
					matches = r >= '0' && r <= '9'
				case '@':
					matches = r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z'
				}
			}
			if !matches {
				return failure("mask", false)
			}
			value = value[size:]
		}
		return failure("mask", value == "")
	}, nil
}
