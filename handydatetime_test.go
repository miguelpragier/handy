package handy

import (
	"testing"
	"time"
)

func TestToday(t *testing.T) {

	today := time.Now()
	y := today.Year()
	m := today.Month()
	d := today.Day()

	dt := time.Date(y, m, d, 0, 0, 0, 0, time.Local)

	format := golangDateTimeFormat("yyyy-mm-dd")

	ds := today.Format(format)

	if xd, xs := Today(); !xd.Equal(dt) || ds != xs {
		t.Errorf("Expected: %v and %s, Get %v and %s", dt, ds, xd, xs)
	}
}

func TestMonthLastDay(t *testing.T) {
	type testStructInput struct {
		year  int
		month int
	}

	type testStruct struct {
		summary        string
		input          testStructInput
		expectedOutput int
	}

	ts := []testStruct{
		{summary: "negative year june", input: testStructInput{-1, 6}, expectedOutput: 30},
		{summary: "zero year june", input: testStructInput{0, 6}, expectedOutput: 30},
		{summary: "negative year february", input: testStructInput{-1, 2}, expectedOutput: 0},
		{summary: "zero year february", input: testStructInput{0, 2}, expectedOutput: 0},
		{summary: "negative month", input: testStructInput{2019, -6}, expectedOutput: 0},
		{summary: "zero month", input: testStructInput{2019, 0}, expectedOutput: 0},
		{summary: "thirteen", input: testStructInput{2019, 13}, expectedOutput: 0},
		{summary: "january", input: testStructInput{2019, 1}, expectedOutput: 31},
		{summary: "february", input: testStructInput{2019, 2}, expectedOutput: 28},
		{summary: "february leap", input: testStructInput{2020, 2}, expectedOutput: 29},
		{summary: "march", input: testStructInput{2020, 3}, expectedOutput: 31},
		{summary: "april", input: testStructInput{2020, 4}, expectedOutput: 30},
		{summary: "may", input: testStructInput{2020, 5}, expectedOutput: 31},
		{summary: "june", input: testStructInput{2020, 6}, expectedOutput: 30},
		{summary: "july", input: testStructInput{2020, 7}, expectedOutput: 31},
		{summary: "august", input: testStructInput{2020, 8}, expectedOutput: 31},
		{summary: "september", input: testStructInput{2020, 9}, expectedOutput: 30},
		{summary: "october", input: testStructInput{2020, 10}, expectedOutput: 31},
		{summary: "november", input: testStructInput{2020, 11}, expectedOutput: 30},
		{summary: "december", input: testStructInput{2020, 12}, expectedOutput: 31},
	}

	for _, tc := range ts {
		t.Run(tc.summary, func(t *testing.T) {
			x := MonthLastDay(tc.input.year, tc.input.month)

			if tc.expectedOutput != x {
				t.Errorf("[%s] Test has failed with year/month %d/%d! Expected: %d and got: %d", tc.summary, tc.input.year, tc.input.month, tc.expectedOutput, x)
			}
		})
	}
}
