package handy

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Violation identifies a failed rule without exposing the input.
type Violation struct {
	Code string
}

// Rule reports violations in their presentation order.
type Rule func(string) []Violation

const invalidUTF8 = "invalid_utf8"

func applyRules(violations []Violation, value string, rules []Rule) []Violation {
	for _, rule := range rules {
		violations = append(violations, rule(value)...)
	}
	return violations
}

// Transform changes a string before the normalized rules run.
type Transform func(string) string

// StringTransform is the original name of Transform.
type StringTransform = Transform

// Policy validates the original input, applies transforms, then validates the result.
// Empty optional values bypass NormalizedRules.
type Policy struct {
	OriginalRules   []Rule
	Transforms      []StringTransform
	NormalizedRules []Rule
	Optional        bool
}

// New creates a policy that transforms input before applying its rules.
func New(transforms []Transform, rules ...Rule) Policy {
	return Policy{Transforms: transforms, NormalizedRules: rules}
}

// Result contains both representations and the rule failures.
type Result struct {
	Original   string
	Value      string
	Violations []Violation
}

// Valid reports whether all applicable rules succeeded.
func (r Result) Valid() bool { return len(r.Violations) == 0 }

// Check applies the policy. Invalid UTF-8 is rejected before Unicode operations.
func (p Policy) Check(input string) Result {
	result := Result{Original: input, Value: input}
	if !utf8.ValidString(input) {
		result.Violations = append(result.Violations, Violation{Code: invalidUTF8})
		return result
	}
	result.Violations = applyRules(result.Violations, input, p.OriginalRules)
	for _, transform := range p.Transforms {
		result.Value = transform(result.Value)
		if !utf8.ValidString(result.Value) {
			result.Violations = append(result.Violations, Violation{Code: invalidUTF8})
			return result
		}
	}
	if p.Optional && result.Value == "" {
		return result
	}
	result.Violations = applyRules(result.Violations, result.Value, p.NormalizedRules)
	return result
}

// TrimSpace removes leading and trailing Unicode whitespace.
func TrimSpace(s string) string { return strings.TrimSpace(s) }

// CompactSpace replaces runs of Unicode whitespace with one ASCII space.
func CompactSpace(s string) string { return strings.Join(strings.Fields(s), " ") }

// Lower converts Unicode letters to lowercase.
func Lower(s string) string { return strings.ToLower(s) }

// Upper converts Unicode letters to uppercase.
func Upper(s string) string { return strings.ToUpper(s) }

// Required rejects empty strings.
func Required(s string) []Violation {
	if s == "" {
		return []Violation{{Code: "required"}}
	}
	return nil
}

func failure(code string, ok bool) []Violation {
	if !ok {
		return []Violation{{Code: code}}
	}
	return nil
}

// ByteLength checks an inclusive byte length interval.
func ByteLength(min, max int) Rule {
	if min < 0 || max < min {
		panic("handy: invalid length interval")
	}
	return func(s string) []Violation { return failure("byte_length", len(s) >= min && len(s) <= max) }
}

// RuneLength checks an inclusive Unicode code point length interval.
func RuneLength(min, max int) Rule {
	if min < 0 || max < min {
		panic("handy: invalid length interval")
	}
	return func(s string) []Violation {
		n := utf8.RuneCountInString(s)
		return failure("rune_length", n >= min && n <= max)
	}
}

// AllowedRunes accepts strings containing only the specified characters.
func AllowedRunes(allowed string) Rule {
	return func(s string) []Violation {
		for _, r := range s {
			if !strings.ContainsRune(allowed, r) {
				return []Violation{{Code: "characters"}}
			}
		}
		return nil
	}
}

// RunesMatching accepts strings whose runes all satisfy predicate.
func RunesMatching(predicate func(rune) bool) Rule {
	if predicate == nil {
		panic("handy: nil rune predicate")
	}
	return func(s string) []Violation {
		for _, r := range s {
			if !predicate(r) {
				return []Violation{{Code: "characters"}}
			}
		}
		return nil
	}
}

// StartsWithASCIIAlpha requires an ASCII letter as the first character.
func StartsWithASCIIAlpha(s string) []Violation {
	return failure("prefix", len(s) > 0 && (s[0] >= 'a' && s[0] <= 'z' || s[0] >= 'A' && s[0] <= 'Z'))
}

// Prefix requires the given prefix.
func Prefix(prefix string) Rule {
	return func(s string) []Violation { return failure("prefix", strings.HasPrefix(s, prefix)) }
}

// Suffix requires the given suffix.
func Suffix(suffix string) Rule {
	return func(s string) []Violation { return failure("suffix", strings.HasSuffix(s, suffix)) }
}

// Pattern requires a regular expression match. Anchor the expression to match the entire string.
func Pattern(pattern *regexp.Regexp) Rule {
	if pattern == nil {
		panic("handy: nil pattern")
	}
	return func(s string) []Violation { return failure("pattern", pattern.MatchString(s)) }
}

// LettersAndSpaces accepts Unicode letters and spaces.
func LettersAndSpaces(r rune) bool { return unicode.IsLetter(r) || unicode.IsSpace(r) }

// ForbiddenRunes rejects any occurrence of the specified characters.
func ForbiddenRunes(forbidden string) Rule {
	return func(s string) []Violation {
		for _, r := range s {
			if strings.ContainsRune(forbidden, r) {
				return []Violation{{Code: "forbidden_characters"}}
			}
		}
		return nil
	}
}

// Separators requires separators to occur only between nonempty segments.
// An empty input has no separators and is accepted.
func Separators(separators string) Rule {
	if separators == "" {
		panic("handy: empty separators")
	}
	return func(s string) []Violation {
		previous := false
		first := true
		for _, r := range s {
			current := strings.ContainsRune(separators, r)
			if current && (first || previous) {
				return []Violation{{Code: "separators"}}
			}
			first = false
			previous = current
		}
		return failure("separators", !previous)
	}
}

// Trim removes leading and trailing Unicode whitespace.
func Trim(s string) string { return TrimSpace(s) }

// Compact replaces runs of Unicode whitespace with one ASCII space.
func Compact(s string) string { return CompactSpace(s) }

// ByteLen checks an inclusive byte length interval.
func ByteLen(min, max int) Rule { return ByteLength(min, max) }

// RuneLen checks an inclusive Unicode code point length interval.
func RuneLen(min, max int) Rule { return RuneLength(min, max) }

// Allow accepts strings containing only the specified runes.
func Allow(runes string) Rule { return AllowedRunes(runes) }

// Deny rejects strings containing any of the specified runes.
func Deny(runes string) Rule { return ForbiddenRunes(runes) }

// MatchRunes accepts strings whose runes all satisfy predicate.
func MatchRunes(predicate func(rune) bool) Rule { return RunesMatching(predicate) }

// ASCIIAlphaStart requires an ASCII letter as the first character.
func ASCIIAlphaStart(s string) []Violation { return StartsWithASCIIAlpha(s) }
