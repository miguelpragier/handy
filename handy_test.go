package handy

import (
	"strings"
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
	testlist := []struct {
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
		{"test require uppercase success", "1234Ab", "1234Ab", 4, CheckNewPasswordComplexityRequireUpperCase|CheckNewPasswordComplexityRequireLetter, CheckNewPasswordResultOK},
		{"test require uppercase error", "1234ab", "1234ab", 4, CheckNewPasswordComplexityRequireUpperCase|CheckNewPasswordComplexityRequireLetter, CheckNewPasswordResultTooSimple},
		{"test require number success", "abc1", "abc1", 4, CheckNewPasswordComplexityRequireNumber, CheckNewPasswordResultOK},
		{"test require number error", "abcd", "abcd", 4, CheckNewPasswordComplexityRequireNumber, CheckNewPasswordResultTooSimple},
		{"test require space success", "abc d", "abc d", 4, CheckNewPasswordComplexityRequireSpace, CheckNewPasswordResultOK},
		{"test require space error", "abcd", "abcd", 4, CheckNewPasswordComplexityRequireSpace, CheckNewPasswordResultTooSimple},
		{"test require symbol success", "abc#", "abc#", 4, CheckNewPasswordComplexityRequireSymbol, CheckNewPasswordResultOK},
		{"test require symbol error", "abcd", "abcd", 4, CheckNewPasswordComplexityRequireSymbol, CheckNewPasswordResultTooSimple},
	}

	for _, tst := range testlist {
		t.Run(tst.summary, func(t *testing.T) {
			tr := CheckNewPassword(tst.password, tst.checkpassword, tst.minimumlength, tst.flag)

			if tr != tst.expectedOutput {
				t.Errorf("Test has failed!\n\tInput: %s,\n\tExpected: %d, \n\tGot: %d", tst.password, tst.expectedOutput, tr)
			}
		})
	}
}

func TestStringHash(t *testing.T) {
	testcases := []TestDefaultTestStruct {
		{"Normal Test", "Handy", "E80649A6418B6C24FCCB199DAB7CB5BD6EC37593EA0285D52D717FCC7AEE5FB3"},
		{"string with number", "123456", "8D969EEF6ECAD3C29A3A629280E686CF0C3F5D5A86AFF3CA12020C923ADC6C92"},
		{"mashup", "Handy12345", "C82333DB3A6D91F98BE188C6C7B928DF4960B9EC3F3EB8CB50293368C673BE3D"},
		{"with symbols", "#handy_12Ax", "507512071AAEA24A94ECBB0F32EE74169FD59160EE9232819C504F39656E61F7"},
	}

	for _, tc := range testcases {
		t.Run(tc.summary, func(t *testing.T) {
			r := StringHash(tc.input.(string))

			if r != strings.ToLower(tc.expectedOutput.(string)) {
				t.Errorf("Test has failed!\n\tInput: %s,\n\tExpected: %d, \n\tGot: %d", tc.input, tc.expectedOutput, r)
			}
		})
	}
}

func TestOnlyLetters(t *testing.T)  {
	tcs := []TestDefaultTestStruct {
		{"empty", "", ""},
		{"only letters", "haoplhu", "haoplhu"},
		{"letters and numbers", "hlo1234", "hlo"},
		{"symbols", "$#@", ""},
		{"numbers", "1234", ""},
		{"with space", "with space", "withspace"},
		{"A full phrase", "Hello Sr! Tell me, how are you?", "HelloSrTellmehowareyou"},

	}

	for _, tc := range tcs {
		t.Run(tc.summary, func(t *testing.T) {
			r := OnlyLetters(tc.input.(string))

			if r != tc.expectedOutput {
				t.Errorf("Test has failed!\n\tInput: %s,\n\tExpected: %s, \n\tGot: %s", tc.input, tc.expectedOutput, r)
			}
		})
	}
}

func TestOnlyDigits(t *testing.T)  {
	tcs := []TestDefaultTestStruct {
		{"empty", "", ""},
		{"only letters", "haoplhu", ""},
		{"letters and numbers", "hlo1234", "1234"},
		{"symbols", "$#@", ""},
		{"numbers", "1234", "1234"},
		{"with space", "with space 10", "10"},
		{"A full phrase", "Hello Sr! I'm 24 Years Old!", "24"},

	}

	for _, tc := range tcs {
		t.Run(tc.summary, func(t *testing.T) {
			r := OnlyDigits(tc.input.(string))

			if r != tc.expectedOutput {
				t.Errorf("Test has failed!\n\tInput: %s,\n\tExpected: %s, \n\tGot: %s", tc.input, tc.expectedOutput, r)
			}
		})
	}
}

func TestOnlyLettersAndNumbers(t *testing.T) {
	tcs := []TestDefaultTestStruct {
		{"empty", "", ""},
		{"only letters", "haoplhu", "haoplhu"},
		{"letters and numbers", "hlo1234", "hlo1234"},
		{"symbols", "$#@", ""},
		{"numbers", "1234", "1234"},
		{"with space", "with space 10", "withspace10"},
		{"A full phrase", "Hello Sr! I'm 24 Years Old!", "HelloSrIm24YearsOld"},

	}

	for _, tc := range tcs {
		t.Run(tc.summary, func(t *testing.T) {
			r := OnlyLettersAndNumbers(tc.input.(string))

			if r != tc.expectedOutput {
				t.Errorf("Test has failed!\n\tInput: %s,\n\tExpected: %s, \n\tGot: %s", tc.input, tc.expectedOutput, r)
			}
		})
	}
}

func TestRandomInt(t *testing.T) {
	tcs := []struct{
		summary string
		min int
		max int
	}{
		{"normal test", int(10), int(20)},
		{"big range", int(10), int(1000)},
		{"negative", int(-10), int(1000)},
	}

	for _, tc := range tcs {
		t.Run(tc.summary, func(t *testing.T) {
			r := RandomInt(tc.min, tc.max)

			if r < tc.min || r > tc.max {
				t.Errorf("Test has failed!\n\tMin: %d, \n\tMax: %d, \n\tGot: %d", tc.min, tc.max, r)
			}
		})
	}
}

func TestCheckPhone(t *testing.T)  {
	tcs := []struct {
		summary        string
		input          string
		allowEmpty     bool
		expectedOutput bool
	}{
		{"Normal input", "948034118", false, true},
		{"Empty return false", "", false, false},
		{"Empty allowing empty", "", true, true},
		{"Normal input but allowing empty", "948034118", true, true},
		{"invalid input", "48034118", false, false},
	}
	
	for _, tc := range tcs {
		t.Run(tc.summary, func(t *testing.T) {
			tr := CheckPhone(tc.input, tc.allowEmpty)
			if tr != tc.expectedOutput {
				t.Errorf("Test has failed!\n\tExpected: %t, \n\tGot: %t, \n\tInput: %s\n\tAllowEmpty: %t", tc.expectedOutput, tr, tc.input, tc.allowEmpty)
			}
		})
	}
}

func TestStringAsFloat(t *testing.T) {
	tcs := []struct {
		summary string
		input string
		decimalSeparator rune
		thousandSeparator rune
		expectedOutput	float64
	}{
		{"Normal Test", "60.42", '.', ',', 60.42},
		{"Negative Test", "-60.42", '.', ',', -60.42},
		{"Virgula como decimal Test", "60.42", ',', '.', 6042.000000},
		{"ERROR TEST", "bla", '.', ',', 00.00},

	}

	for _, tc := range tcs {
		t.Run(tc.summary, func(t *testing.T) {
			tr := StringAsFloat(tc.input, tc.decimalSeparator, tc.thousandSeparator)
			 if tr != tc.expectedOutput {
				 t.Errorf("Test has failed!\n\tExpected: %f, \n\tGot: %f, \n\tInput: %s\n\tDecimalSeparator: %c\n\tThousandSeparator: %c", tc.expectedOutput, tr, tc.input, tc.decimalSeparator, tc.thousandSeparator)
			 }
		})
	}
}

func TestStringAsInteger(t *testing.T) {
	tcs := []TestDefaultTestStruct {
		{"default test", "30", 30},
		{"negative", "-30", -30},
		{"double", "30.5", 0},
		{"text", "text", 0},
	}

	for _, tc := range tcs {
		t.Run(tc.summary, func(t *testing.T) {
			tr := StringAsInteger(tc.input.(string))

			if tr != tc.expectedOutput {
				t.Errorf("Test has failed!\n\tExpected: %d, \n\tGot: %d, \n\tInput: %s", tc.expectedOutput, tr, tc.input)
			}
		})
	}
}

func TestBetween(t *testing.T) {
	tcs := []struct{
		summary        string
		number         int
		min            int
		max            int
		expectedOutput bool
	}{
		{"normal test", 5, 0, 10, true},
		{"with negative", -5, -10, 0, true},
		{"mix with negative", 5, -10, 100, true},
		{"Large Numbers", 1000, -1000, 10000, true},
		{"Give me a false!", -5, -4, 0, false},
	}

	for _, tc := range tcs {
		t.Run(tc.summary, func(t *testing.T) {
			tr := Between(tc.number, tc.min, tc.max)

			if tr != tc.expectedOutput {
				t.Errorf("Test has failed!\n\tExpected: %t, \n\tGot: %t, \n\tInput: %d, \n\tMin: %d, \n\tMax: %d", tc.expectedOutput, tr, tc.number, tc.min, tc.max)
			}
		})
	}
}

func TestTif(t *testing.T){

	tcs := []struct{
		summary        string
		condition      bool
		tifThen        interface{}
		tifElse        interface{}
		expectedOutput interface{}
	} {
		{"Normal Test", 5 < 10, "true", "false", "true"},
		{"False", 5 > 10, "true", "false", "false"},
		{"False with numbers", 5 > 10, 10, 15, 15},
		{"True with numbers", 5 < 10, 10, 15, 10},
		{"True with bool", 5 < 10, 5<10, 5>10, 5<10},
		{"False with bool", 5 > 10, 5<10, 5>10, 5>10},
	}

	for _, tc := range tcs {
		t.Run(tc.summary, func(t *testing.T) {
			tr := Tif(tc.condition, tc.tifThen, tc.tifElse)

			if tr != tc.expectedOutput {
				t.Errorf("Test has failed!\n\tExpected: %t, \n\tGot: %v, \n\tInput: %t, \n\tThen: %v, \n\tElse: %v", tc.expectedOutput, tr, tc.condition, tc.tifThen, tc.tifElse)

			}
		})
	}
}

func TestTruncate(t *testing.T) {
	tcs := []struct{
		summary 	   string
		input   	   string
		limit   	   int
		trim    	   bool
		expectedOutput string

	}{
		{"normal Test", "The Go programming language is an open source project to make programmers more productive.", 25, false, "The Go programming langua"},
		{"normal Test with trim", "   The Go programming language is an open source project to make programmers more productive.", 45, true, "The Go programming language is an open sou"},
		{"zero", "The Go programming language is an open source project to make programmers more productive.", 0, true, ""},
		{"zero zero", "", 45, true, ""},
	}

	for _, tc := range tcs {
		t.Run(tc.summary, func(t *testing.T) {
			tr := Truncate(tc.input, tc.limit, tc.trim)
			 if tr != tc.expectedOutput {
				 t.Errorf("Test has failed!\n\tExpected: %s, \n\tGot: %s, \n\tInput: %s, \n\tlimit: %d, \n\ttrim: %t", tc.expectedOutput, tr, tc.input, tc.limit, tc.trim)
			 }
		})
	}
}

func TestTransform(t *testing.T) {
	tcs := []struct{
		summary        string
		input          string
		max            int
		flags          uint8
		expectedOutput string
	}{
		{"without flags", "The Go programming language is an open source project to make programmers more productive.", 20, TransformNone, "The Go programming l"},
	}

	for _, tc := range tcs {
		t.Run(tc.summary, func(t *testing.T) {
			tr := Transform(tc.input, tc.max, tc.flags)

			if tr != tc.expectedOutput {
				t.Errorf("Test has failed!\n\tExpected: %s, \n\tGot: %s, \n\tInput: %s, \n\tlimit: %d, \n\tflags: %d", tc.expectedOutput, tr, tc.input, tc.max, tc.flags)
			}
		})
	}
}

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

//func CheckMinLen(value string, minLength int) bool {
//	value = strings.TrimSpace(value)
//
//	return TrimLen(value) >= minLength
//}

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

//func Bit(x interface{}) uint8 {
//	if IsNumericType(x) && x != 0 {
//		return 1
//	}
//
//	return 0
//}

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
