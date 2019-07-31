package handy

import "testing"

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
		{"Sending but not accepting empty string", " ", false, CheckPersonNameResultTooShort},
		{"Sending spaces-only string and refusing empty", "     ", false, CheckPersonNameResultTooShort},
		{"Sending numbers, expecting false", " 5454 ", true, CheckPersonNameResultPolluted},
		{"OneWorded string", "ONEWORD", false, CheckPersonNameResultTooFewWords},
		{"Minimum acceptable", "AB CDE", false, CheckPersonNameResultOK},
		{"Non-ascii stuff", "ÑÔÑÀSÇÏÏ ÇÃO ÀË", false, CheckPersonNameResultOK},
		{"Words with symbols. Expecting true", "WORDS-WITH SYMBOLS'", false, CheckPersonNameResultOK},
		{"Words with symbols. Expecting false", "WORDS WITH SYMBOLS`", false, CheckPersonNameResultPolluted},
		{"less than two letters", "a", false, CheckPersonNameResultTooFewWords},
		{"Sending numbers, expecting false", "5454", false, CheckPersonNameResultPolluted},
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

func stringSlicesAreEqual(x, y []string) bool {
	if ((x == nil) != (y == nil)) || (len(x) != len(y)) {
		return false
	}

	for i := range y {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}

func TestNameFirstAndLast(t *testing.T) {
	type TestNameFirstAndLastStruct struct {
		summary         string
		name            string
		transformFlags  uint
		expectedOutputS string
		expectedOutputA []string
	}

	testlist := []TestNameFirstAndLastStruct{
		{"Only two letters", "x Y", TransformNone, `x Y`, []string{"x", "Y"}},
		{"one word name", "namë", TransformNone, `namë`, []string{`namë`}},
		{"all non-ascii runes", "çá öáã àÿ", TransformNone, `çá àÿ`, []string{`çá`, `àÿ`}},
		{"all non-ascii runes to upper", "çá öáã àÿ", TransformFlagUpperCase, `ÇÁ ÀŸ`, []string{`ÇÁ`, `ÀŸ`}},
		{"mixing letters and numbers and then filtering digits off", "W0RDS W1TH NUMB3RS", TransformFlagRemoveDigits, `WRDS NUMBRS`, []string{`WRDS`, `NUMBRS`}},
		{"empty string", "", TransformNone, ``, []string{}},
		{"only spaces", "     ", TransformNone, ``, []string{}},
		{"with spaces and tabs", " FIRST NAME - MIDDLENAME 	LAST	 ", TransformNone, `FIRST LAST`, []string{`FIRST`, `LAST`}},
		{"last name single rune", "NAME X", TransformNone, `NAME X`, []string{`NAME`, `X`}},
		{"only symbols", "5454#@$", TransformNone, `5454#@$`, []string{`5454#@$`}},
		{"single letter", "x", TransformNone, `x`, []string{`x`}},
		{"only spaces empty return", " 		 ", TransformNone, ``, []string{}},
		{"regular name to upper", "name lastname", TransformFlagUpperCase, `NAME LASTNAME`, []string{`NAME`, `LASTNAME`}},
		{"regular name to title", "name LASTNAME", TransformFlagTitleCase, `Name Lastname`, []string{`Name`, `Lastname`}},
		{"REGULAR Name to lOwEr", "name LASTNAME", TransformFlagLowerCase, `name lastname`, []string{`name`, `lastname`}},
	}

	for _, tst := range testlist {
		t.Run(tst.summary, func(t *testing.T) {
			s, a := NameFirstAndLast(tst.name, tst.transformFlags)

			if s != tst.expectedOutputS {
				t.Errorf(`[%s] Test has failed! Given name: "%s", Expected string: "%s", Got: "%s"`, tst.summary, tst.name, tst.expectedOutputS, s)
			}

			if !stringSlicesAreEqual(a, tst.expectedOutputA) {
				t.Errorf(`[%s] Test has failed! Given name: "%s", Expected array: %#v, Got: %#v`, tst.summary, tst.name, tst.expectedOutputA, a)
			}
		})
	}
}
