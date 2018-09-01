package handy

import (
	"math"
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

// CheckDatef validates a date using the given format
func CheckDatef(format, dateTime string) bool {
	f := golangDateTimeFormat(format)

	if f == "" {
		return false
	}

	_, err := time.Parse(f, dateTime)

	return err == nil
}

// CheckDate returns true if given sequence is a valid date in format yyyymmdd
// The function removes non-digit characteres like "yyyy/mm/dd" or "yyyy-mm-dd", filtering to "yyyymmdd"
func CheckDateYMD(yyyymmdd string) bool {
	// Se já chegar vazio, falha
	if yyyymmdd == "" {
		return false
	}

	// Allow 9 digits in order to check if consoumer sent a wrong date
	// It means I won't truncate a badlyFormated string
	yyyymmdd = Transform(yyyymmdd, 9, TransformFlagOnlyDigits)

	if len(yyyymmdd) != 8 {
		return false
	}

	_, err := time.Parse("20060102", yyyymmdd)

	return err == nil
}

// YMDasDateUTC returns a valid UTC time from the given yyymmdd-formatted sequence
func YMDasDateUTC(yyyymmdd string, utc bool) (time.Time, error) {
	yyyymmdd = OnlyDigits(yyyymmdd)

	if t, err := time.Parse("20060102", yyyymmdd); err == nil {
		if utc {
			return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC), nil
		} else {
			return t, nil
		}
	} else {
		return time.Time{}, err
	}
}

// YMDasDate returns a valid time from the given yyymmdd-formatted sequence
func YMDasDate(yyyymmdd string) (time.Time, error) {
	return YMDasDateUTC(yyyymmdd, false)
}

// ElapsedMonths returns the number of elapsed months between two given dates
func ElapsedMonths(from, to time.Time) int {
	// To produce calculations, todate must be greater than from
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
