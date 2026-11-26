package handy

import (
	"crypto/subtle"
	"unicode/utf8"
)

// CheckSecret runs rules in order on the original value and returns only
// violations. It rejects malformed UTF-8 before calling any rules.
func CheckSecret(value string, rules ...Rule) []Violation {
	if !utf8.ValidString(value) {
		return failure(invalidUTF8, false)
	}
	return applyRules(nil, value, rules)
}

// Blocklist rejects exact matches from the supplied values with the code
// "blocklist". It copies the values into a lookup when the rule is created.
func Blocklist(values []string) Rule {
	blocked := make(map[string]struct{}, len(values))
	for _, value := range values {
		blocked[value] = struct{}{}
	}
	return func(value string) []Violation {
		_, found := blocked[value]
		return failure("blocklist", !found)
	}
}

// MinRunes requires at least min runes that satisfy predicate. It reports
// "min_runes" on failure and panics for a negative min or nil predicate.
func MinRunes(min int, predicate func(rune) bool) Rule {
	if min < 0 || predicate == nil {
		panic("handy: MinRunes requires a nonnegative minimum and a predicate")
	}
	return func(value string) []Violation {
		remaining := min
		if remaining == 0 {
			return nil
		}
		for _, r := range value {
			if predicate(r) {
				remaining--
				if remaining == 0 {
					return nil
				}
			}
		}
		return failure("min_runes", false)
	}
}

// Confirm compares the values byte for byte and reports "confirmation" on
// mismatch. Use CheckSecret to validate the value separately.
func Confirm(value, confirmation string) []Violation {
	return failure("confirmation", subtle.ConstantTimeCompare([]byte(value), []byte(confirmation)) == 1)
}
