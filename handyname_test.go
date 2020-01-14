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

//// StringSlicesAreEqual compares two string slices and returns true if they have the same elements, in same order
//func StringSlicesAreEqual(x, y []string) bool {
//	if ((x == nil) != (y == nil)) || (len(x) != len(y)) {
//		return false
//	}
//
//	for i := range y {
//		if x[i] != y[i] {
//			return false
//		}
//	}
//
//	return true
//}

func TestNameFirstAndLast(t *testing.T) {
	type TestNameFirstAndLastStruct struct {
		summary         string
		name            string
		transformFlags  uint
		expectedOutputS string
	}

	testlist := []TestNameFirstAndLastStruct{
		{"Only two letters", "x Y", TransformNone, `x Y`},
		{"one word name", "namë", TransformNone, `namë`},
		{"all non-ascii runes", "çá öáã àÿ", TransformNone, `çá àÿ`},
		{"all non-ascii runes to upper", "çá öáã àÿ", TransformFlagUpperCase, `ÇÁ ÀŸ`},
		{"mixing letters and numbers and then filtering digits off", "W0RDS W1TH NUMB3RS", TransformFlagRemoveDigits, `WRDS NUMBRS`},
		{"empty string", "", TransformNone, ``},
		{"only spaces", "     ", TransformNone, ``},
		{"with spaces and tabs", " FIRST NAME - MIDDLENAME 	LAST	 ", TransformNone, `FIRST LAST`},
		{"last name single rune", "NAME X", TransformNone, `NAME X`},
		{"only symbols", "5454#@$", TransformNone, `5454#@$`},
		{"single letter", "x", TransformNone, `x`},
		{"only spaces empty return", " 		 ", TransformNone, ``},
		{"regular name to upper", "name lastname", TransformFlagUpperCase, `NAME LASTNAME`},
		{"regular name to title", "name LASTNAME", TransformFlagTitleCase, `Name Lastname`},
		{"REGULAR Name to lOwEr", "name LASTNAME", TransformFlagLowerCase, `name lastname`},
	}

	for _, tst := range testlist {
		t.Run(tst.summary, func(t *testing.T) {
			s := NameFirstAndLast(tst.name, tst.transformFlags)

			if s != tst.expectedOutputS {
				t.Errorf(`[%s] Test has failed! Given name: "%s", Expected string: "%s", Got: "%s"`, tst.summary, tst.name, tst.expectedOutputS, s)
			}
		})
	}
}

func TestNameInitials(t *testing.T) {
	type tStruct struct {
		summary        string
		name           string
		transformFlags uint
		expectedOutput string
	}

	testlist := []tStruct{
		{`simplest 2 words name`, `miguel pragier`, TransformNone, `m p`},
		{`3 words name separated`, `ivan alexandrovitch kleshtakov`, TransformNone, `i a k`},
		{`3 words with unicode`, `Ívän Âlexandrovitch Çzelyatchenko`, TransformNone, `Í Â Ç`},
		{`3 words with unicode title-case`, `ívän âlexandrovitch çzelyatchenko`, TransformFlagTitleCase, `Í Â Ç`},
		{`empty string`, ``, TransformNone, ``},
		{`dot`, `.`, TransformNone, `.`},
		{`spaces and tabs`, "  \t\t \n", TransformNone, ``},
		{`name with tabs`, "richard\t\tstallmann", TransformNone, `r s`},
		{`noble name with 1`, `dom pedro 1`, TransformNone, `d p 1`},
		{`noble name with I uppercase`, `dom pedro I`, TransformFlagUpperCase, `D P I`},
		{`3 letters`, `x y z`, TransformNone, `x y z`},
		{`one word`, `asingleword`, TransformNone, `a`},
		{`comma separators`, `name,with,comma,separators`, TransformNone, `n`},
	}

	for _, tst := range testlist {
		t.Run(tst.summary, func(t *testing.T) {
			s := NameInitials(tst.name, tst.transformFlags)

			if s != tst.expectedOutput {
				t.Errorf(`[%s] Test has failed! Given name: "%s", Expected string: "%s", Got: "%s"`, tst.summary, tst.name, tst.expectedOutput, s)
			}
		})
	}
}
