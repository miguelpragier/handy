package handy

import (
	"strings"
	"time"
)

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

// CheckDateYMD returns true if given input is a valid date in format yyyymmdd
// The function removes non-digit characters like "yyyy/mm/dd" or "yyyy-mm-dd", filtering to "yyyymmdd"
func CheckDateYMD(yyyymmdd string) bool {
	return CheckDate("yyyymmdd", yyyymmdd)
}

// YMDasDateUTC returns a valid UTC time from the given yyymmdd-formatted input
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

// YMDasDate returns a valid time from the given yyymmdd-formatted input
func YMDasDate(yyyymmdd string) (time.Time, error) {
	return YMDasDateUTC(yyyymmdd, false)
}

// ElapsedTime returns difference between two dates in years, months, days, hours, monutes and seconds
// Thanks to icza@https://stackoverflow.com/a/36531443/1301019
func ElapsedTime(dtx, dty time.Time) (int, int, int, int, int, int) {
	// If locations are different, convert one to make them the same
	if dtx.Location() != dty.Location() {
		dty = dty.In(dtx.Location())
	}

	// if the dates are equal,
	if dtx.Equal(dty) {
		return 0, 0, 0, 0, 0, 0
	}

	// For correct calculations, assure dtx is before or equal
	if dtx.After(dty) {
		dtx, dty = dty, dtx
	}

	y1, M1, d1 := dtx.Date()

	y2, M2, d2 := dty.Date()

	h1, m1, s1 := dtx.Clock()

	h2, m2, s2 := dty.Clock()

	year := y2 - y1
	month := int(M2 - M1)
	day := d2 - d1
	hour := h2 - h1
	min := m2 - m1
	sec := s2 - s1

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}

	if min < 0 {
		min += 60
		hour--
	}

	if hour < 0 {
		hour += 24
		day--
	}

	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)

		day += 32 - t.Day()

		month--
	}

	if month < 0 {
		month += 12

		year--
	}

	return year, month, day, hour, min, sec
}

// ElapsedMonths returns the number of elapsed months between two given dates
func ElapsedMonths(from, to time.Time) int {
	// To produce calculations, "to" must be greater than "from"
	if to.Before(from) || (from.Year() == to.Year() && from.Month() == to.Month()) {
		return 0
	}

	_, months, _, _, _, _ := ElapsedTime(from, to)

	return months
}

// ElapsedYears returns the number of elapsed years between two given dates
func ElapsedYears(from, to time.Time) int {
	// To produce calculations, "to" must be greater than "from"
	if to.Before(from) || (from.Year() == to.Year()) {
		return 0
	}

	years, _, _, _, _, _ := ElapsedTime(from, to)

	return years
}

// YearsAge returns the number of years past since a given date
func YearsAge(birthdate time.Time) int {
	return ElapsedYears(birthdate, time.Now())
}

// MonthLastDay returns the last day of month, considering the year for cover february in leap years
// Month is one-based, then january=1
// If month is different of 2 (february) year is ignored
// Of month is february, year have to be valid
func MonthLastDay(year int, month int) int {
	if month < 1 || month > 12 {
		return 0
	}

	var lastDayArrayZeroBased = []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if month != 2 {
		return lastDayArrayZeroBased[month-1]
	}

	if year <= 0 {
		return 0
	}

	if ((year%4 == 0) && (year%100 != 0)) || (year%400 == 0) {
		return 29
	}

	return 28
}

// DateReformat gets a date string in a given currentFormat, and transform it according newFormat
func DateReformat(dt1 string, currentFormat, newFormat string) string {
	if dx := StringAsDateTime(dt1, currentFormat); !dx.IsZero() {
		return DateTimeAsString(dx, newFormat)
	}

	return ""
}

type DateStrCheck uint8

const (
	DateStrCheckOk            DateStrCheck = 0
	DateStrCheckErrInvalid    DateStrCheck = 1
	DateStrCheckErrOutOfRange DateStrCheck = 2
	DateStrCheckErrEmpty      DateStrCheck = 3
)

// DateStrCheckAge checks a date string considering minimum and maximum age
// The resulting code can be translated to text, according prefered idiom, with DateStrCheckErrMessage
func DateStrCheckAge(date, format string, yearsAgeMin, yearsAgeMax int, acceptEmpty bool) DateStrCheck {
	if date == "" {
		if !acceptEmpty {
			return DateStrCheckErrEmpty
		}

		return DateStrCheckOk
	}

	dt := StringAsDateTime(date, format)

	if dt.IsZero() {
		return DateStrCheckErrInvalid
	}

	yearsAge := YearsAge(dt)

	if !Between(yearsAge, yearsAgeMin, yearsAgeMax) {
		return DateStrCheckErrOutOfRange
	}

	return DateStrCheckOk
}

// DateStrCheckRange checks a date string considering minimum and maximum date range
// The resulting code can be translated to text, according prefered idiom, with DateStrCheckErrMessage
func DateStrCheckRange(date, format string, dateMin, dateMax time.Time, acceptEmpty bool) DateStrCheck {
	if date == "" {
		if !acceptEmpty {
			return DateStrCheckErrEmpty
		}

		return DateStrCheckOk
	}

	dt := StringAsDateTime(date, format)

	if dt.IsZero() {
		return DateStrCheckErrInvalid
	}

	if !dateMin.IsZero() {
		if dt.Before(dateMin) {
			return DateStrCheckErrOutOfRange
		}
	}

	if !dateMax.IsZero() {
		if dt.After(dateMax) {
			return DateStrCheckErrOutOfRange
		}
	}

	return DateStrCheckOk
}
