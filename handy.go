// Package handy is a toolbelt with utilities and helpers like validators, sanitizers and string formatters.
// There are routines to filter strings, convert between types, validate passwords with custom rules, easily format dates and much more.
package handy

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

var reDupSpaces = regexp.MustCompile(`\s+`)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

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

//RuneHasSymbol returns true if the given rune contains a symbol
func RuneHasSymbol(ru rune) bool {
	allowedSymbols := "!\"#$%&'()*+´-./:;<=>?@[\\]^_`{|}~"

	for _, r := range allowedSymbols {
		if ru == r {
			return true
		}
	}

	return false
}

// StringHash simply generates a SHA256 hash from the given string
// In case of error, return ""
func StringHash(s string) string {
	h := sha256.New()

	if _, err := h.Write([]byte(s)); err != nil {
		return ""
	}

	sum := h.Sum(nil)

	return fmt.Sprintf("%x", sum)
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

// RandomInt returns a random integer within the given (inclusive) range
func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

// RandomIntArray returns an array filled with random integer numbers
func RandomIntArray(min, max, howMany int) []int {
	var a []int

	for i := 0; i < howMany; i++ {
		a = append(a, RandomInt(min, max))
	}

	return a
}

// RandomReseed restarts the randonSeeder and returns a random integer within the given (inclusive) range
func RandomReseed(min, max int) int {
	x := time.Now().UTC().UnixNano() + int64(rand.Int())

	rand.Seed(x)

	return rand.Intn(max-min) + min
}

// CheckPhone returns true if a given input has between 9 and 14 digits
func CheckPhone(phone string, acceptEmpty bool) bool {
	phone = OnlyDigits(phone)

	return (acceptEmpty && (phone == "")) || ((len([]rune(phone)) >= 9) && (len([]rune(phone)) <= 14))
}

// StringAsFloat tries to convert a string to float, and if it can't, just returns zero
// It's limited to one billion
func StringAsFloat(s string, decimalSeparator, thousandsSeparator rune) float64 {
	if s == "" {
		return 0.0
	}

	const maxLength = 20

	if len([]rune(s)) > maxLength {
		s = s[0:maxLength]
	}

	s = strings.Replace(s, string(thousandsSeparator), "", -1)

	s = strings.Replace(s, string(decimalSeparator), ".", -1)

	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f
	}

	return 0.0
}

// StringAsInteger returns the integer value extracted from string, or zero
func StringAsInteger(s string) int {
	if s == "" {
		return 0
	}

	if i, err := strconv.ParseInt(s, 10, 32); err == nil {
		return int(i)
	}

	return 0
}

// Between checks if param n in between low and high integer params
func Between(n, low, high int) bool {
	return n >= low && n <= high
}

// Tif is a simple implementation of the dear ternary IF operator
func Tif(condition bool, tifThen, tifElse interface{}) interface{} {
	if condition {
		return tifThen
	}

	return tifElse
}

// Truncate limits the length of a given string, trimming or not, according parameters
func Truncate(s string, maxLen int, trim bool) string {
	if s == "" {
		return s
	}

	if len(s) > maxLen {
		s = s[0:maxLen]
	}

	if trim {
		s = strings.TrimSpace(s)
	}

	return s
}

// MatchesAny returns true if any of the given items matches ( equals ) the subject ( search parameter )
func MatchesAny(search interface{}, items ...interface{}) bool {
	for _, v := range items {
		if fmt.Sprintf("%T", search) == fmt.Sprintf("%T", v) {
			if search == v {
				return true
			}
		}
	}

	return false
}

// HasOnlyNumbers returns true if the input is entirely numeric
func HasOnlyNumbers(sequence string) bool {
	if utf8.RuneCountInString(sequence) == 0 {
		return false
	}

	for _, r := range sequence {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

// HasOnlyDigits returns true if the input is entirely numeric
// HasOnlyDigits is an alias for HasOnlyNumbers()
func HasOnlyDigits(sequence string) bool {
	return HasOnlyNumbers(sequence)
}

// HasOnlyLetters returns true if the input is entirely composed by letters
func HasOnlyLetters(sequence string) bool {
	if utf8.RuneCountInString(sequence) == 0 {
		return false
	}

	for _, r := range sequence {
		if !unicode.IsLetter(r) {
			return false
		}
	}

	return true
}

// TrimLen returns the runes count after trim the spaces
func TrimLen(text string) int {
	if text == "" {
		return 0
	}

	text = strings.TrimSpace(text)

	if text == "" {
		return 0
	}

	return utf8.RuneCountInString(text)
}

// CheckMinLen verifies if the rune-count is greater then or equal the given minimum
// It returns true if the given string has length greater than or equal than minLength parameter
func CheckMinLen(value string, minLength int) bool {
	value = strings.TrimSpace(value)

	return TrimLen(value) >= minLength
}

// IsNumericType checks if an interface's concrete type corresponds to some of golang native numeric types
func IsNumericType(x interface{}) bool {
	switch x.(type) {
	case uint:
		return true
	case uint8: // Or byte
		return true
	case uint16:
		return true
	case uint32:
		return true
	case uint64:
		return true
	case int:
		return true
	case int8:
		return true
	case int16:
		return true
	case int32:
		return true
	case float32:
		return true
	case float64:
		return true
	case complex64:
		return true
	case complex128:
		return true
	default:
		return false
	}
}

// Bit returns only uint8(0) or uint8(1).
// It receives an interface, and when it's a number, and when this number is 0 (zero) it returns 0. Otherwise it returns 1 (one)
// If the interface is not a number, it returns 0 (zero)
func Bit(x interface{}) uint8 {
	if IsNumericType(x) && x != 0 {
		return 1
	}

	return 0
}

// Boolean returns the bool version/interpretation of some value;
// It receives an interface, and when this is a number, Boolean() returns flase of zero and true for different from zero.
// If it's a string, try to find "1", "T", "TRUE" to return true.
// Any other case returns false
func Boolean(x interface{}) bool {
	if IsNumericType(x) {
		return x != 0
	}

	if s, ok := x.(string); ok {
		s = Transform(s, 4, TransformFlagLowerCase|TransformFlagTrim)
		return MatchesAny(s, "1", "true", "t")
	}

	return false
}

// Reverse returns the given string written backwards, with letters reversed.
func Reverse(s string) string {
	if utf8.RuneCountInString(s) < 2 {
		return s
	}

	r := []rune(s)

	buffer := make([]rune, len(r))

	for i, j := len(r)-1, 0; i >= 0; i-- {
		buffer[j] = r[i]
		j++
	}

	return string(buffer)
}

// StringReplaceAll keeps replacing until there's no more ocurrences to replace.
func StringReplaceAll(original string, replacementPairs ...string) string {
	if original == "" {
		return original
	}

	r := strings.NewReplacer(replacementPairs...)

	for {
		result := r.Replace(original)

		if original != result {
			original = result
		} else {
			break
		}
	}

	return original
}

// ArrayDifferenceAtoB returns the items from A that doesn't exist in B
func ArrayDifferenceAtoB(a, b []int) []int {
	m := make(map[int]bool)

	for _, item := range b {
		m[item] = true
	}

	var diff []int

	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}

	return diff
}

// ArrayDifference returns all items that doesn't exist in both given arrays
func ArrayDifference(a, b []int) []int {
	diffA := ArrayDifferenceAtoB(a, b)

	diffB := ArrayDifferenceAtoB(b, a)

	return append(diffB, diffA...)
}

// PositiveOrZero checks if a signed number is negative, and in this case returns zero.
func PositiveOrZero(n int) int {
	if n < 0 {
		return 0
	}

	return n
}
