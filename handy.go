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

var (
	reEmail     = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	reDupSpaces = regexp.MustCompile(`\s+`)
)

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

// CheckEmail returns true if the given input is a valid email address
// Observe that CheckEmail doesn't trim nor sanitize string before check
// See https://tools.ietf.org/html/rfc2822#section-3.4.1 for details about email address anatomy
func CheckEmail(email string) bool {
	if email == "" {
		return false
	}

	return reEmail.MatchString(email)
}

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
