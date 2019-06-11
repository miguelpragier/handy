package handy

import (
	"math"
	"strings"
	"time"
)

var lastDayArray = []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

// golangDateFormat translate handy's arbitrary date format to Go's eccentric format
func golangDateTimeFormat(format string) string {
	if format == "" {
		return ""
	}

	newFormat := strings.ToLower(format)

	newFormat = strings.Replace(newFormat, "yyyy", "2006", -1)
	newFormat = strings.Replace(newFormat, "yy", "06", -1)
	newFormat = strings.Replace(newFormat, "mmmm", "January", -1)
	newFormat = strings.Replace(newFormat, "mmm", "Jan", -1)
	newFormat = strings.Replace(newFormat, "mm", "01", -1)
	newFormat = strings.Replace(newFormat, "m", "1", -1)
	newFormat = strings.Replace(newFormat, "dd", "02", -1)
	newFormat = strings.Replace(newFormat, "d", "2", -1)
	newFormat = strings.Replace(newFormat, "hh24", "15", -1)
	newFormat = strings.Replace(newFormat, "hh", "03 PM", -1)
	newFormat = strings.Replace(newFormat, "h", "3 PM", -1)
	newFormat = strings.Replace(newFormat, "nn", "04", -1)
	newFormat = strings.Replace(newFormat, "n", "4", -1)
	newFormat = strings.Replace(newFormat, "ss", "05", -1)
	newFormat = strings.Replace(newFormat, "s", "5", -1)
	newFormat = strings.Replace(newFormat, "ww", "Monday", -1)
	newFormat = strings.Replace(newFormat, "w", "Mon", -1)

	return newFormat
}

// DateTimeAsString formats time.Time variables as strings, considering the format directive
func DateTimeAsString(dt time.Time, format string) string {
	newFormat := golangDateTimeFormat(format)

	return dt.Format(newFormat)
}

// NowAsString formats time.Now() as string, considering the format directive
func NowAsString(format string) string {
	newFormat := golangDateTimeFormat(format)

	return time.Now().Format(newFormat)
}

// Today returns today's date at zero hours, minutes, seconds, etc.
// It returns a time and a yyyy-mm-dd formated string
func Today() (time.Time, string) {
	newFormat := golangDateTimeFormat("yyyy-mm-dd")

	t := time.Now()
	y := t.Year()
	m := t.Month()
	d := t.Day()

	return time.Date(y, m, d, 0, 0, 0, 0, time.Local), t.Format(newFormat)
}

// Todayf returns today's date at zero hours, minutes, seconds, etc.
// It returns a time and a custom formated string
func Todayf(format string) (time.Time, string) {
	newFormat := golangDateTimeFormat(format)

	t := time.Now()
	y := t.Year()
	m := t.Month()
	d := t.Day()

	return time.Date(y, m, d, 0, 0, 0, 0, time.Local), t.Format(newFormat)
}

// YMD returns today's date tokenized as year, month and day of month
func YMD() (int, int, int) {
	t := time.Now()

	return t.Year(), int(t.Month()), t.Day()
}

// StringAsDateTime formats time.Time variables as strings, considering the format directive
func StringAsDateTime(s string, format string) time.Time {
	goFormat := golangDateTimeFormat(format)

	t, _ := time.Parse(goFormat, s)

	return t
}

// CheckDate validates a date using the given format
func CheckDate(format, dateTime string) bool {
	f := golangDateTimeFormat(format)

	if f == "" {
		return false
	}

	_, err := time.Parse(f, dateTime)

	return err == nil
}

// CheckDateYMD returns true if given sequence is a valid date in format yyyymmdd
// The function removes non-digit characters like "yyyy/mm/dd" or "yyyy-mm-dd", filtering to "yyyymmdd"
func CheckDateYMD(yyyymmdd string) bool {
	return CheckDate("yyyymmdd", yyyymmdd)
}

// YMDasDateUTC returns a valid UTC time from the given yyymmdd-formatted sequence
func YMDasDateUTC(yyyymmdd string, utc bool) (time.Time, error) {
	yyyymmdd = OnlyDigits(yyyymmdd)

	t, err := time.Parse("20060102", yyyymmdd)

	if err != nil {
		return time.Time{}, err
	}

	if utc {
		return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC), nil
	}

	return t, nil
}

// YMDasDate returns a valid time from the given yyymmdd-formatted sequence
func YMDasDate(yyyymmdd string) (time.Time, error) {
	return YMDasDateUTC(yyyymmdd, false)
}

// ElapsedMonths returns the number of elapsed months between two given dates
func ElapsedMonths(from, to time.Time) int {
	// To produce calculations, "to" must be greater than "from"
	if to.Before(from) || (from.Year() == to.Year() && from.Month() == to.Month()) {
		return 0
	}

	diff := to.Sub(from)

	hours := diff.Hours()

	days := hours / 24

	return int(math.Abs(days / 30))
}

// ElapsedYears returns the number of elapsed years between two given dates
func ElapsedYears(from, to time.Time) int {
	months := float64(ElapsedMonths(from, to))
	return int(math.Abs(months / 12))
}

// YearsAge returns the number of years past since a given date
func YearsAge(birthdate time.Time) int {
	return ElapsedYears(birthdate, time.Now())
}

// MonthLastDay returns the last day of month, considering the year for cover february in leap years
func MonthLastDay(year int, month int) int {
	if month != 2 {
		return lastDayArray[month]
	}

	if ((year%4 == 0) && (year%100 != 0)) || (year%400 == 0) {
		return 29
	}

	return 28
}
