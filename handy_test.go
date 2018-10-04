package handy

import (
	"testing"
)

type TestDefaultTestStruct struct {
	summary        string
	input          interface{}
	expectedOutput interface{}
}

func TestCheckEmail(t *testing.T) {

	testList := []TestDefaultTestStruct{
		{"send empty string", "", false},
		{"send invalid address", "email-gmail.com", false},
		{"send valid address", "email@gmail.com", true},
	}

	for _, tst := range testList {
		t.Run(tst.summary, func(t *testing.T) {
			tr := CheckEmail(tst.input.(string))

			if tr != tst.expectedOutput {
				t.Errorf("Test has failed!\n\tEmail: %s, \n\tExpected: %s", tst.input, tst.expectedOutput)
			}
		})
	}
}

func TestCheckPersonName(t *testing.T) {
	type TestStructForCheckPersonName struct {
		summary        string
		name           string
		acceptEmpty    bool
		expectedOutput uint8
	}

	testlist := []TestStructForCheckPersonName{
		{"Only two letters", "T S", false, CheckPersonNameResultTooSimple},
		{"only four letters", "AB CD", false, CheckPersonNameResultTooSimple},
		{"five letters with non-ascii runes", "ça vá", false, CheckPersonNameResultTooSimple},
		{"mixing letters and numbers", "W0RDS W1TH NUMB3RS", false, CheckPersonNameResultPolluted},
		{"Sending and accepting empty string", "", true, CheckPersonNameResultOK},
		{"Sending spaces-only string and accepting empty", "     ", true, CheckPersonNameResultOK},
		{"Sending but not accepting empty string", "", false, CheckPersonNameResultTooShort},
		{"Sending spaces-only string and refusing empty", "     ", false, CheckPersonNameResultTooShort},
		{"Sending numbers, expecting false", " 5454 ", true, CheckPersonNameResultPolluted},
		{"OneWorded string", "ONEWORD", false, CheckPersonNameResultTooFewWords},
		{"Minimum acceptable", "AB CDE", false, CheckPersonNameResultOK},
		{"Non-ascii stuff", "ÑÔÑÀSÇÏÏ ÇÃO ÀË", false, CheckPersonNameResultOK},
		{"Words with symbols. Expecting true", "WORDS-WITH SYMBOLS'", false, CheckPersonNameResultOK},
		{"Words with symbols. Expecting false", "WORDS WITH SYMBOLS`", false, CheckPersonNameResultPolluted},
	}

	for _, tst := range testlist {
		t.Run(tst.summary, func(t *testing.T) {
			tr := CheckPersonName(tst.name, tst.acceptEmpty)

			if tr != tst.expectedOutput {
				t.Errorf("Test has failed!\n\tName: %s\n\tAcceptEmpty: %t, \n\tExpected: %d, \n\tGot: %d,", tst.name, tst.acceptEmpty, tst.expectedOutput, tr)
			}
		})
	}
}

func TestCheckCPF(t *testing.T) {
	testlist := []TestDefaultTestStruct{
		{"send empty string", "", false},
		{"send wrong length string (10)", "153.255.555.4", false},
		{"send wrong length string (12)", "153.255.555.455", false},
		{"send cheating cpf", "55555555555", false},
		{"send invalid string", "153.278.966.A6", false},
		{"send alright string", "03818534110", true},
	}

	for _, tst := range testlist {
		t.Run(tst.summary, func(t *testing.T) {
			tr := CheckCPF(tst.input.(string))

			if tr != tst.expectedOutput {
				t.Errorf("Test has failed!\n\tCPF: %s,\n\tExpected: %t,\n\tGot: %t", tst.input, tst.expectedOutput, tr)
			}
		})
	}
}

func TestCheckCNPJ(t *testing.T) {
	testlist := []TestDefaultTestStruct{
		{"send empty string", "", false},
		{"send wrong length string (13)", "88.015.315/0001-5", false},
		{"send wrong length string (15)", "88.015.315/0001-5003", false},
		{"send cheating cnpj", "55555555555555", false},
		{"send invalid string", "88.015.315/0001-5A", false},
		{"send alright string with punctuation", "88.015.315/0001-53", true},
		{"send alright string", "88015315000153", true},
	}

	for _, tst := range testlist {
		t.Run(tst.summary, func(t *testing.T) {
			tr := CheckCNPJ(tst.input.(string))

			if tr != tst.expectedOutput {
				t.Errorf("Test has failed!\n\tCNPJ: %s,\n\tExpected: %t, \n\tGot: %t", tst.input, tst.expectedOutput, tr)
			}
		})
	}
}

func TestCheckDate(t *testing.T) {
	type TestStructForCheckDate struct {
		summary        string
		format         string
		date           string
		expectedOutput bool
	}

	testlist := []TestStructForCheckDate{
		{"empty string", "", "", false},
		{"invalid date", "2006-01-02", "2018-02-29", false},
		{"invalid date", "2006-01-02", "2018-13-01", false},
		{"invalid date", "2006-01-02", "2018-12-32", false},
		{"valid date 1", "2006-01-02", "2018-12-31", true},
		{"valid date 2", "20060102", "20180101", true},
		{"invalid date format 1", "20060102", "2018-01-01", false},
		{"invalid date format 1", "2006-01-02", "20180201", false},
	}

	for _, tst := range testlist {
		t.Run(tst.summary, func(t *testing.T) {
			tr := CheckDate(tst.format, tst.date)

			if tr != tst.expectedOutput {
				t.Errorf("Test has failed!\n\tDate: %s,\n\tExpected: %t, \n\tGot: %t, \n\tFormat: %s", tst.date, tst.expectedOutput, tr, tst.format)
			}
		})
	}
}

func TestAmountAsWord(t *testing.T) {
	testlist := []TestDefaultTestStruct{
		{"zero", 0, "zero"},
		{"-125", -125, "menos cento e vinte e cinco"},
		{"-987654321", -987654321, "menos novecentos e oitenta e sete milhões seicentos e cinquenta e quatro mil trezentos e vinte e um"},
	}
	for _, tst := range testlist {
		t.Run(tst.summary, func(t *testing.T) {
			tr := AmountAsWord(int64(tst.input.(int)))

			if tr != tst.expectedOutput {
				t.Errorf("Test has failed!\n\tInput: %d,\n\tExpected: %s, \n\tGot: %s", tst.input, tst.expectedOutput, tr)
			}
		})
	}
}

func TestCheckNewPassword(t *testing.T) {
	testCases := []struct {
		summary        string
		password       string
		checkpassword  string
		minimumlength  uint
		flag           uint8
		expectedOutput uint8
	}{
		{"test lowest flag", "1234AB", "1234AB", 6, CheckNewPasswordComplexityLowest, CheckNewPasswordResultOK},
		{"test check password", "1234AB", "1234A", 6, CheckNewPasswordComplexityLowest, CheckNewPasswordResultDivergent},
		{"Only Numbers with Default Flag", "1234", "1234", 4, CheckNewPasswordComplexityLowest, CheckNewPasswordResultOK},
		{"Only letters with Default Flag", "lala", "lala", 4, CheckNewPasswordComplexityLowest, CheckNewPasswordResultOK},
		{"testing minimum length", "1234", "1234", 2, CheckNewPasswordComplexityLowest, CheckNewPasswordResultOK},
		{"testing minimum length for password", "123", "123", 2, CheckNewPasswordComplexityLowest, CheckNewPasswordResultTooShort},
		{"test require letter success", "1234AB", "1234AB", 4, CheckNewPasswordComplexityRequireLetter, CheckNewPasswordResultOK},
		{"test require letter error", "1234", "1234", 4, CheckNewPasswordComplexityRequireLetter, CheckNewPasswordResultTooSimple},
		{"test require letuppercaseter success", "1234Ab", "1234Ab", 4, CheckNewPasswordComplexityRequireUpperCase, CheckNewPasswordResultOK},
		{"test require uppercase error", "1234ab", "1234ab", 4, CheckNewPasswordComplexityRequireUpperCase, CheckNewPasswordResultTooSimple},

	}
}

//// StringHash simply generates a SHA256 hash from the given string
//func StringHash(s string) string {
//	h := sha256.New()
//
//	h.Write([]byte(s))
//
//	sum := h.Sum(nil)
//
//	return fmt.Sprintf("%x", sum)
//}
//
//// OnlyLetters returns only the letters from the given string, after strip all the rest ( numbers, spaces, etc. )
//func OnlyLetters(sequence string) string {
//	if utf8.RuneCountInString(sequence) == 0 {
//		return ""
//	}
//
//	var letters []rune
//
//	for _, r := range []rune(sequence) {
//		if unicode.IsLetter(r) {
//			letters = append(letters, r)
//		}
//	}
//
//	return string(letters)
//}
//
//// OnlyDigits returns only the numbers from the given string, after strip all the rest ( letters, spaces, etc. )
//func OnlyDigits(sequence string) string {
//	if utf8.RuneCountInString(sequence) > 0 {
//		re, _ := regexp.Compile("[\\D]")
//
//		sequence = re.ReplaceAllString(sequence, "")
//	}
//
//	return sequence
//}
//
//// OnlyLettersAndNumbers returns only the letters and numbers from the given string, after strip all the rest, like spaces and special symbols.
//func OnlyLettersAndNumbers(sequence string) string {
//	if utf8.RuneCountInString(sequence) == 0 {
//		return ""
//	}
//
//	var aplhanumeric []rune
//
//	for _, r := range []rune(sequence) {
//		if unicode.IsLetter(r) || unicode.IsDigit(r) {
//			aplhanumeric = append(aplhanumeric, r)
//		}
//	}
//
//	return string(aplhanumeric)
//}
//
//// RandomInt returns a rondom integer within the given (inclusive) range
//func RandomInt(min, max int) int {
//	rand.Seed(time.Now().UTC().UnixNano())
//
//	return rand.Intn(max) + min
//}
//
//// CheckPhone returns true if a given sequence has between 9 and 14 digits
//func CheckPhone(phone string, acceptEmpty bool) bool {
//	phone = OnlyDigits(phone)
//
//	return (acceptEmpty && (phone == "")) || ((len([]rune(phone)) >= 9) && (len([]rune(phone)) <= 14))
//}
//
//// StringAsFloat tries to convert a string to float, and if it can't, just returns zero
//// It's limited to one billion
//func StringAsFloat(s string, decimalSeparator, thousandsSeparator rune) float64 {
//	if s == "" {
//		return 0.0
//	}
//
//	const BillionLength = 12
//
//	if len([]rune(s)) > BillionLength {
//		s = s[0:12]
//	}
//
//	s = strings.Replace(s, string(thousandsSeparator), "", -1)
//
//	s = strings.Replace(s, string(decimalSeparator), ".", -1)
//
//	if f, err := strconv.ParseFloat(s, 64); err == nil {
//		return f
//	}
//
//	return 0.0
//}
//
//// StringAsInteger returns the integer value extracted from string, or zero
//func StringAsInteger(s string) int {
//	if s == "" {
//		return 0
//	}
//
//	if i, err := strconv.ParseInt(s, 10, 32); err == nil {
//		return int(i)
//	}
//
//	return 0
//}
//
//// Between checks if param n in between low and high integer params
//func Between(n, low, high int) bool {
//	return n >= low && n <= high
//}
//
//// Tif is a simple implementation of the dear ternary IF operator
//func Tif(condition bool, tifThen, tifElse interface{}) interface{} {
//	if condition {
//		return tifThen
//	}
//
//	return tifElse
//}
//
//// Truncate limits the length of a given string, trimming or not, according parameters
//func Truncate(s string, maxLen int, trim bool) string {
//	if s == "" {
//		return s
//	}
//
//	if len(s) > maxLen {
//		s = s[0:maxLen]
//	}
//
//	if trim {
//		s = strings.TrimSpace(s)
//	}
//
//	return s
//}
//
//const (
//	// TransformNone No transformations are ordered. Only constraints maximum length
//	TransformNone = uint8(1)
//	// TransformFlagTrim Trims the string, removing leading and trailing spaces
//	TransformFlagTrim = uint8(2)
//	// TransformFlagLowerCase Makes the string lowercase
//	TransformFlagLowerCase = uint8(4)
//	// TransformFlagUpperCase Makes the string uppercase
//	TransformFlagUpperCase = uint8(8)
//	// TransformFlagOnlyDigits Removes all non-numeric characters
//	TransformFlagOnlyDigits = uint8(16)
//	// TransformFlagOnlyLetters Removes all non-letter characters
//	TransformFlagOnlyLetters = uint8(32)
//	// TransformFlagOnlyLettersAndDigits Leaves only letters and numbers
//	TransformFlagOnlyLettersAndDigits = uint8(64)
//	// TransformFlagHash After process all pther flags, applies SHA256 hashing on string for output
//	TransformFlagHash = uint8(128)
//)
//
//// Transform handles a string according given flags/parametrization, as follows:
//// Available Flags to be used alone or combined:
////	TransformNone Does nothing, and turns all other flags OFF.
////	TransformFlagTrim Trim spaces before and after process the input
////	TransformFlagLowerCase Change string case to lower. If it's combined with TransformFlagUpperCase, only uppercase remains, once is executed later.
////	TransformFlagUpperCase Change string case to upper. If it's combined with TransformFlagLowerCase, only uppercase remains, once it's executed later.
////	TransformFlagOnlyDigits Filter/strip all but digits
////	TransformFlagOnlyLetters Filter/strip all but letters
////	TransformFlagOnlyLettersAndDigits Filter/strip all but numbers and letters. Removes spaces, punctuation and special symbols
//// 	TransformFlagHash Apply handy.StringHash() routine to string
//func Transform(s string, maxLen int, transformFlags uint8) string {
//	if s == "" {
//		return s
//	}
//
//	if transformFlags&TransformNone != TransformNone {
//
//		if (transformFlags & TransformFlagOnlyLettersAndDigits) == TransformFlagOnlyLettersAndDigits {
//			s = OnlyLettersAndNumbers(s)
//
//			if s == "" {
//				return s
//			}
//		} else if (transformFlags & TransformFlagOnlyDigits) == TransformFlagOnlyDigits {
//			s = OnlyDigits(s)
//
//			if s == "" {
//				return s
//			}
//		} else if (transformFlags & TransformFlagOnlyLetters) == TransformFlagOnlyLetters {
//			s = OnlyLetters(s)
//
//			if s == "" {
//				return s
//			}
//		}
//
//		// Have to trim before and after, to avoid issues with string truncation and new leading/trailing spaces
//		if (transformFlags & TransformFlagTrim) == TransformFlagTrim {
//			s = strings.TrimSpace(s)
//		}
//
//		if utf8.RuneCountInString(s) > maxLen {
//			s = string([]rune(s)[:maxLen])
//		}
//
//		// Have to trim before and after, to avoid issues with string truncation and new leading/trailing spaces
//		if (transformFlags & TransformFlagTrim) == TransformFlagTrim {
//			s = strings.TrimSpace(s)
//		}
//
//		if (transformFlags & TransformFlagLowerCase) == TransformFlagLowerCase {
//			s = strings.ToLower(s)
//		}
//
//		if (transformFlags & TransformFlagUpperCase) == TransformFlagUpperCase {
//			s = strings.ToUpper(s)
//		}
//
//		if (transformFlags & TransformFlagHash) == TransformFlagHash {
//			s = StringHash(s)
//		}
//	}
//
//	return s
//}
//
//// MatchesAny returns true if any of the given items matches ( equals ) the subject ( search parameter )
//func MatchesAny(search interface{}, items ...interface{}) bool {
//	for _, v := range items {
//		if fmt.Sprintf("%T", search) == fmt.Sprintf("%T", v) {
//			if search == v {
//				return true
//			}
//		}
//	}
//
//	return false
//}
//
//// HasOnlyNumbers returns true if the sequence is entirely numeric
//func HasOnlyNumbers(sequence string) bool {
//	if utf8.RuneCountInString(sequence) == 0 {
//		return false
//	}
//
//	for _, r := range []rune(sequence) {
//		if !unicode.IsDigit(r) {
//			return false
//		}
//	}
//
//	return true
//}
//
//// HasOnlyLetters returns true if the sequence is entirely composed by letters
//func HasOnlyLetters(sequence string) bool {
//	if utf8.RuneCountInString(sequence) == 0 {
//		return false
//	}
//
//	for _, r := range []rune(sequence) {
//		if !unicode.IsLetter(r) {
//			return false
//		}
//	}
//
//	return true
//}
//
//// TrimLen returns the runes count after trim the spaces
//func TrimLen(text string) int {
//	if text == "" {
//		return 0
//	}
//
//	text = strings.TrimSpace(text)
//
//	if text == "" {
//		return 0
//	}
//
//	return utf8.RuneCountInString(text)
//}
//
//// CheckMinLen verifies if the rune-count is greater then or equal the given minimum
//// It returns true if the given string has length greater than or equal than minLength parameter
//func CheckMinLen(value string, minLength int) bool {
//	value = strings.TrimSpace(value)
//
//	return TrimLen(value) >= minLength
//}
//
//// IsNumericType checks if an interface's concrete type corresponds to some of golang native numeric types
//func IsNumericType(x interface{}) bool {
//	switch x.(type) {
//	case uint:
//		return true
//	case uint8: // Or byte
//		return true
//	case uint16:
//		return true
//	case uint32:
//		return true
//	case uint64:
//		return true
//	case int:
//		return true
//	case int8:
//		return true
//	case int16:
//		return true
//	case int32:
//		return true
//	case float32:
//		return true
//	case float64:
//		return true
//	case complex64:
//		return true
//	case complex128:
//		return true
//	default:
//		return false
//	}
//}
//
//// Bit returns only uint8(0) or uint8(1).
//// It receives an interface, and when it's a number, and when this number is 0 (zero) it returns 0. Otherwise it returns 1 (one)
//// If the interface is not a number, it returns 0 (zero)
//func Bit(x interface{}) uint8 {
//	if IsNumericType(x) && x != 0 {
//		return 1
//	}
//
//	return 0
//}
//
//// Boolean returns the bool version/interpretation of some value;
//// It receives an interface, and when this is a number, Boolean() returns flase of zero and true for different from zero.
//// If it's a string, try to find "1", "T", "TRUE" to return true.
//// Any other case returns false
//func Boolean(x interface{}) bool {
//	if IsNumericType(x) {
//		return x != 0
//	}
//
//	if s, ok := x.(string); ok {
//		s = Transform(s, 4, TransformFlagLowerCase|TransformFlagTrim)
//		return MatchesAny(s, "1", "true", "t")
//	}
//
//	return false
//}
//
//// Reverse returns the given string written backwards, with letters reversed.
//func Reverse(s string) string {
//	if utf8.RuneCountInString(s) < 2 {
//		return s
//	}
//
//	r := []rune(s)
//
//	buffer := make([]rune, len(r))
//
//	for i, j := len(r)-1, 0; i >= 0; i-- {
//		buffer[j] = r[i]
//		j++
//	}
//
//	return string(buffer)
//}
//
//// OnlyURL strip all symbols non allowed in URLs and returns the sanitized url.
//func OnlyURL(url string) string {
//	allowedSymbols := []rune("$-_.+!*'(),{}|\\^~[]`<>#%\";/?:@&=.")
//	tmp := []rune(url)
//	var target []rune
//
//	for _, r := range tmp {
//		if InArray(r, allowedSymbols) || unicode.IsLetter(r) || unicode.IsNumber(r) {
//			target = append(target, r)
//		}
//	}
//
//	return string(tmp)
//}
//
//const (
//	// CheckPersonNameResultOK means the name was validated
//	CheckPersonNameResultOK = 0
//	// CheckPersonNameResultPolluted The routine only accepts letters, single quotes and spaces
//	CheckPersonNameResultPolluted = 1
//	// CheckPersonNameResultTooFewWords The funcion requires at least 2 words
//	CheckPersonNameResultTooFewWords = 2
//	// CheckPersonNameResultTooShort the sum of all characters must be >= 6
//	CheckPersonNameResultTooShort = 3
//	// CheckPersonNameResultTooSimple The name rule requires that at least one word
//	CheckPersonNameResultTooSimple = 4
//)
//
//// CheckPersonName returns true if the name contains at least two words, one >= 3 chars and one >=2 chars.
//// I understand that this is a particular criteria, but this is the OpenSourceMagic, where you can change and adapt to your own specs.
//func CheckPersonName(name string, acceptEmpty bool) uint8 {
//	name = strings.TrimSpace(name)
//
//	// If name is empty, AND it's accepted, return ok. Else, cry!
//	if name == "" {
//		if !acceptEmpty {
//			return CheckPersonNameResultTooShort
//		}
//
//		return CheckPersonNameResultOK
//	}
//
//	// Person names doesn't accept other than letters, spaces and single quotes
//	for _, r := range []rune(name) {
//		if !unicode.IsLetter(r) && r != ' ' && r != '\'' {
//			return CheckPersonNameResultPolluted
//		}
//	}
//
//	// A complete name has to be at least 2 words.
//	a := strings.Fields(name)
//
//	if len(a) < 2 {
//		return CheckPersonNameResultTooFewWords
//	}
//
//	// At least two words, one with 3 chars and other with 2
//	found2 := false
//	found3 := false
//
//	for _, s := range a {
//		if !found3 && utf8.RuneCountInString(s) >= 3 {
//			found3 = true
//			continue
//		}
//
//		if !found2 && utf8.RuneCountInString(s) >= 2 {
//			found2 = true
//			continue
//		}
//	}
//
//	if !found2 || !found3 {
//		return CheckPersonNameResultTooSimple
//	}
//
//	return CheckPersonNameResultOK
//}
//
//// StringReplaceAll keeps replacing until there's no more ocurrences to replace.
//func StringReplaceAll(original string, replacementPairs ...string) string {
//	if original==""{
//		return original
//	}
//
//	r := strings.NewReplacer(replacementPairs...)
//
//	for {
//		result := r.Replace(original)
//
//		if original!=result {
//			original = result
//		} else {
//			break
//		}
//	}
//
//	return original
//}
