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

// CheckEmail returns true if the given sequence is a valid email address
// See https://tools.ietf.org/html/rfc2822#section-3.4.1 for details about email address anatomy
func CheckEmail(email string) bool {
	if email == "" {
		return false
	}

	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	return re.MatchString(email)
}

const (
	CheckNewPasswordResultOK        = 0
	CheckNewPasswordResultDivergent = 1
	CheckNewPasswordResultTooShort  = 2
	CheckNewPasswordResultTooSimple = 3
)

// Run some basic checks on new password strings
// My rule requires at least six chars, with at least one letter and at least one number.
func CheckNewPassword(password, passwordConfirmation string) uint8 {
	const minPasswordLength = 6

	if utf8.RuneCountInString(strings.TrimSpace(password)) < minPasswordLength {
		return CheckNewPasswordResultTooShort
	}

	if password != passwordConfirmation {
		return CheckNewPasswordResultDivergent
	}

	letters := OnlyLetters(password)

	digits := OnlyDigits(password)

	if letters == "" || digits == "" {
		return CheckNewPasswordResultTooSimple
	}

	return CheckNewPasswordResultOK
}

// StringHash simply generates a SHA256 hash from the given string
func StringHash(password string) string {
	h := sha256.New()

	h.Write([]byte(password))

	sum := h.Sum(nil)

	return fmt.Sprintf("%x", sum)
}

// OnlyLetters returns only the letters from the given string, after strip all the rest ( numbers, spaces, etc. )
func OnlyLetters(sequence string) string {
	if utf8.RuneCountInString(sequence) == 0 {
		return ""
	}

	var letters []rune

	for _, r := range []rune(sequence) {
		if unicode.IsLetter(r) {
			letters = append(letters, r)
		}
	}

	return string(letters)
}

// OnlyDigits returns only the numbers from the given string, after strip all the rest ( letters, spaces, etc. )
func OnlyDigits(sequence string) string {
	if utf8.RuneCountInString(sequence) > 0 {
		re, _ := regexp.Compile("[\\D]")

		sequence = re.ReplaceAllString(sequence, "")
	}

	return sequence
}

// OnlyLettersAndNumbers returns only the letters and numbers from the given string, after strip all the rest, like spaces and special symbols.
func OnlyLettersAndNumbers(sequence string) string {
	if utf8.RuneCountInString(sequence) == 0 {
		return ""
	}

	var aplhanumeric []rune

	for _, r := range []rune(sequence) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			aplhanumeric = append(aplhanumeric, r)
		}
	}

	return string(aplhanumeric)
}

// RandomInt returns a rondom integer within the given (inclusive) range
func RandomInt(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())

	return rand.Intn(max) + min
}

// CheckPhone returns true if a given sequence has between 9 and 14 digits
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

	const BillionLength = 12

	if len([]rune(s)) > BillionLength {
		s = s[0:12]
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
	} else {
		return tifElse
	}
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
	TransformNone                     = uint8(0)
	TransformFlagTrim                 = uint8(2)
	TransformFlagLowerCase            = uint8(4)
	TransformFlagUpperCase            = uint8(8)
	TransformFlagOnlyDigits           = uint8(16)
	TransformFlagOnlyLetters          = uint8(32)
	TransformFlagOnlyLettersAndDigits = uint8(64)
	TransformFlagHash                 = uint8(128)
)

// Transform handles a string according given flags/parametrization, as follows:
// Available Flags to be used alone or combined:
//	TransformNone - Does nothing. It's only for truncation.
//	TransformFlagTrim - Trim spaces before and after proccess the input
//	TransformFlagLowerCase - Change string case to lower
//	TransformFlagUpperCase - Change string case to upper
//	TransformFlagOnlyDigits - Filter/strip all but digits
//	TransformFlagOnlyLetters - Filter/strip all but letters
//	TransformFlagOnlyLettersAndDigits - Filter/strip all but numbers and letters. Removes spaces, punctuation and special symbols
// 	TransformFlagHash - Apply handy.StringHash() routine to string
func Transform(s string, maxLen int, transformFlags uint8) string {
	if s == "" {
		return s
	}

	if (transformFlags & TransformFlagOnlyLettersAndDigits) == TransformFlagOnlyLettersAndDigits {
		s = OnlyLettersAndNumbers(s)

		if s == "" {
			return s
		}
	} else if (transformFlags & TransformFlagOnlyDigits) == TransformFlagOnlyDigits {
		s = OnlyDigits(s)

		if s == "" {
			return s
		}
	} else if (transformFlags & TransformFlagOnlyLetters) == TransformFlagOnlyLetters {
		s = OnlyLetters(s)

		if s == "" {
			return s
		}
	}

	// Have to trim before and after, to avoid issues with string truncation and new leading/trailing spaces
	if (transformFlags & TransformFlagTrim) == TransformFlagTrim {
		s = strings.TrimSpace(s)
	}

	if utf8.RuneCountInString(s) > maxLen {
		s = string([]rune(s)[:maxLen])
	}

	// Have to trim before and after, to avoid issues with string truncation and new leading/trailing spaces
	if (transformFlags & TransformFlagTrim) == TransformFlagTrim {
		s = strings.TrimSpace(s)
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

// HasOnlyNumbers returns true if the sequence is entirely numeric
func HasOnlyNumbers(sequence string) bool {
	if utf8.RuneCountInString(sequence) == 0 {
		return false
	}

	for _, r := range []rune(sequence) {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

// HasOnlyNumbers returns true if the sequence is entirely composed by letters
func HasOnlyLetters(sequence string) bool {
	if utf8.RuneCountInString(sequence) == 0 {
		return false
	}

	for _, r := range []rune(sequence) {
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

// OnlyURL strip all symbols non allowed in URLs and returns the sanitized url.
func OnlyURL(url string) string {
	allowedSymbols := []rune("$-_.+!*'(),{}|\\^~[]`<>#%\";/?:@&=.")
	tmp := []rune(url)
	var target []rune

	for _, r := range tmp {
		if InArray(r, allowedSymbols) || unicode.IsLetter(r) || unicode.IsNumber(r) {
			target = append(target, r)
		}
	}

	return string(tmp)
}

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
		} else {
			return CheckPersonNameResultOK
		}
	}

	// Person names doesn't accept other than letters, spaces and single quotes
	for _, r := range []rune(name) {
		if !unicode.IsLetter(r) && r != ' ' && r != '\'' {
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
