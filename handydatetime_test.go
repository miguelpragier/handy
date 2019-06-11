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

//func TestTodayf(t *testing.T){
//	newFormat := golangDateTimeFormat(format)
//
//	t := time.Now()
//	y := t.Year()
//	m := t.Month()
//	d := t.Day()
//
//	return time.Date(y, m, d, 0, 0, 0, 0, time.Local), t.Format(newFormat)
//}
//
//// YMD returns today's date tokenized as year, month and day of month
//func TestYMD(t *testing.T){
//	t := time.Now()
//
//	return t.Year(), int(t.Month()), t.Day()
//}
//
