package handy

import (
	"testing"
)

//func TestCheckPersonName(t *testing.T) {
//	type TestStructForCheckPersonName struct {
//		summary        string
//		name           string
//		acceptEmpty    bool
//		expectedOutput bool
//	}
//
//	testlist := []TestStructForCheckPersonName{
//		{"Only two letters", "T S", false, false},
//		{"only four letters", "AB CD", false, false},
//		{"five letters with non-ascii runes", "ça vá", false, false},
//		{"mixing letters and numbers", "W0RDS W1TH NUMB3RS", false, false},
//		{"Sending and accepting empty string", "", true, true},
//		{"Sending spaces-only string and accepting empty", "     ", true, true},
//		{"Sending but not accepting empty string", "", false, false},
//		{"Sending spaces-only string and refusing empty", "     ", false, false},
//		{"Sending numbers, expecting false", " 5454 ", true, false},
//		{"OneWorded string", "ONEWORD", false, false},
//		{"Minimum acceptable", "AB CDE", false, true},
//		{"Non-ascii stuff", "ÑÔÑ-ÀSÇÏÏ ÇÃO ÀË", false, true},
//		{"Words with symbols. Expecting true", "WORDS-WITH SYMBOLS'", false, true},
//		{"Words with symbols. Expecting false", "WORDS WITH SYMBOLS`", false, true},
//	}
//
//	for _, tst := range testlist {
//		tr := CheckPersonName(tst.name, tst.acceptEmpty)
//
//		if tr != tst.expectedOutput {
//			t.Error("Test has failed!\n", "\n\tName: ", tst.name, "\n\tAcceptEmpty: ", tst.acceptEmpty, "\n\tExpected: ", tst.expectedOutput, "\n\tSummary: ", tst.summary)
//		}
//	}
//}

func TestCheckCPF(t *testing.T) {
	type TestStructForCheckCPF struct {
		summary        string
		cpf            string
		expectedOutput bool
	}

	testlist := []TestStructForCheckCPF{
		{"send empty string", "", false},
		{"send wrong length string (10)", "153.255.555.4", false},
		{"send wrong length string (12)", "153.255.555.455", false},
		{"send cheating cpf", "55555555555", false},
		{"send invalid string", "153.278.966.A6", false},
		{"send alright string", "03818534110", true},
	}

	for _, tst := range testlist {
		tr := CheckCPF(tst.cpf)

		if tr != tst.expectedOutput {
			t.Error("Test has failed!\n", "\n\tCPF: ", tst.cpf, "\n\tExpected: ", tst.expectedOutput, "\n\tSummary: ", tst.summary)
		}
	}
}

func TestCheckCNPJ(t *testing.T) {
	type TestStructForCheckCNPJ struct {
		summary        string
		cnpj           string
		expectedOutput bool
	}

	testlist := []TestStructForCheckCNPJ{
		{"send empty string", "", false},
		{"send wrong length string (13)", "88.015.315/0001-5", false},
		{"send wrong length string (15)", "88.015.315/0001-5003", false},
		{"send cheating cnpj", "55555555555555", false},
		{"send invalid string", "88.015.315/0001-5A", false},
		{"send alright string with punctuation", "88.015.315/0001-53", true},
		{"send alright string", "88015315000153", true},
	}

	for _, tst := range testlist {
		tr := CheckCNPJ(tst.cnpj)

		if tr != tst.expectedOutput {
			t.Error("Test has failed!\n", "\n\tCNPJ: ", tst.cnpj, "\n\tExpected: ", tst.expectedOutput, "\n\tSummary: ", tst.summary)
		}
	}
}

func TestCheckEmail(t *testing.T) {
	type TestStructForCheckEmail struct {
		summary        string
		email          string
		expectedOutput bool
	}

	testlist := []TestStructForCheckEmail{
		{"send empty string", "", false},
		{"send invalid address", "email-gmail.com", false},
		{"send valid address", "email@gmail.com", true},
	}

	for _, tst := range testlist {
		tr := CheckEmail(tst.email)

		if tr != tst.expectedOutput {
			t.Error("Test has failed!\n", "\n\tEmail: ", tst.email, "\n\tExpected: ", tst.expectedOutput, "\n\tSummary: ", tst.summary)
		}
	}
}

func TestCheckDate(t *testing.T) {
	type TestStructForCheckDate struct {
		summary        string
		date           string
		expectedOutput bool
	}

	testlist := []TestStructForCheckDate{
		{"empty string", "", false},
		{"invalid date", "2018-02-29", false},
		{"invalid date", "2018-13-01", false},
		{"invalid date", "2018-12-32", false},
		{"valid date 1", "2018-12-31", true},
		{"valid date 2", "2018-01-01", true},
	}

	for _, tst := range testlist {
		tr := CheckDateYMD(tst.date)

		if tr != tst.expectedOutput {
			t.Error("Test has failed!\n", "\n\tDate: ", tst.date, "\n\tExpected: ", tst.expectedOutput, "\n\tSummary: ", tst.summary)
		}
	}
}

func TestPorExtenso(t *testing.T) {
	type TestStructForPorExtenso struct {
		summary        string
		value          int64
		expectedOutput string
	}

	testlist := []TestStructForPorExtenso{
		{"zero", 0, "zero"},
		{"-125", -125, "menos cento e vinte e cinco"},
		{"-987654321", -987654321, "menos novecentos e oitenta e sete milhões seicentos e cinquenta e quatro mil trezentos e vinte e um"},
	}

	for _, tst := range testlist {
		tr := AmountAsWord(tst.value)

		if tr != tst.expectedOutput {
			t.Error("Test has failed!\n", "\n\tValue: ", tst.value, "\n\tExpected: ", tst.expectedOutput, "\n\tReceived: ", tr, "\n\tSummary: ", tst.summary)
		}
	}
}

//func CheckNewPassword(password, passwordConfirmation string) (bool, string) {
//	const minPasswordLength = 6
//
//	if utf8.RuneCountInString(strings.TrimSpace(password)) < minPasswordLength {
//		return false, fmt.Sprintf("Senha deve conter ao menos %d caracteres, com ao menos uma letra ou ao menos um número", minPasswordLength)
//	}
//
//	if password != passwordConfirmation {
//		return false, "Senha e confirmação diferentes entre si"
//	}
//
//	re, _ := regexp.Compile("[\\d]")
//
//	letters := re.ReplaceAllString(password, "")
//
//	re, _ = regexp.Compile("[\\D]")
//
//	digits := re.ReplaceAllString(password, "")
//
//	if letters == "" || digits == "" {
//		return false, fmt.Sprintf("Senha deve conter ao menos %d caracteres, com ao menos uma letra ou ao menos um número", minPasswordLength)
//	}
//
//	return true, ""
//}
//
//func PasswordHash(password string) string {
//	h := sha256.New()
//
//	h.Write([]byte(password))
//
//	sum := h.Sum(nil)
//
//	return fmt.Sprintf("%x", sum)
//}
//
//func OnlyAlpha(sequence string) string {
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
//func RandomInt(min, max int) int {
//	rand.Seed(time.Now().UTC().UnixNano())
//
//	return rand.Intn(max) + min
//}
//
//func CheckPhone(phone string, acceptEmpty bool) bool {
//	phone = OnlyDigits(phone)
//
//	return (acceptEmpty && (phone == "")) || ((len([]rune(phone)) >= 9) && (len([]rune(phone)) <= 14))
//}
//
//func YMDasDateUTC(yyyymmdd string, utc bool) (time.Time, error) {
//	yyyymmdd = OnlyDigits(yyyymmdd)
//
//	if t, err := time.Parse("20060102", yyyymmdd); err == nil {
//		if utc {
//			return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC), nil
//		} else {
//			return t, nil
//		}
//	} else {
//		return time.Time{}, err
//	}
//}
//
//func YMDasDate(yyyymmdd string) (time.Time, error) {
//	return YMDasDateUTC(yyyymmdd, false)
//}
//
//func ElapsedMonths(fromDate, toDate time.Time) int {
//	// Se toDate for anterior a fromDate, retornar 0
//	if !fromDate.Before(toDate) {
//		return 0
//	}
//
//	// Se as datas estiverem no mesmo ano, retornar a diferença em meses
//	if fromDate.Year() == toDate.Year() {
//		// Se for o mesmo mês, retornar 0
//		if fromDate.Month() == toDate.Month() {
//			return 0
//		}
//
//		// Se o mês for posterior, mas o dia for anterior ( ex: 30/01 --> 15/02 ) então descontar um mês
//		if (toDate.Month() <= fromDate.Month()) && (fromDate.Day() > toDate.Day()) {
//			return int(toDate.Month()-fromDate.Month()) - 1
//		}
//
//		return int(toDate.Month() - fromDate.Month())
//	}
//
//	years := toDate.Year() - fromDate.Year()
//
//	if years == 1 {
//		// Se o mês for posterior, mas o dia for anterior ( ex: 30/01 --> 15/02 ) então descontar um mês
//		if fromDate.Day() > toDate.Day() {
//			return (12 - fromDate.Year()) + int(toDate.Month()) - 1
//		}
//
//		return (12 - fromDate.Year()) + int(toDate.Month())
//	}
//
//	months := (years - 1) * 12
//
//	months += int(toDate.Month())
//
//	if toDate.Day() > fromDate.Day() {
//		return months - 1
//	}
//
//	return months
//}
//
//func ElapsedMonths2(fromDate, toDate time.Time) int {
//	// Se toDate for anterior a fromDate, retornar 0
//	if !fromDate.Before(toDate) {
//		return 0
//	}
//
//	months := 0
//
//	if fromDate.Year() == toDate.Year() {
//		if fromDate.Month() == toDate.Month() {
//			return 0
//		}
//
//		if toDate.Day() > fromDate.Day() {
//			return int(toDate.Month() - fromDate.Month())
//		}
//
//		return int(toDate.Month()-fromDate.Month()) - 1
//	}
//
//	for y := fromDate.Year(); y <= toDate.Year(); y++ {
//		// Se for o primeiro ano, apenas soma os meses restantes
//		if y == fromDate.Year() {
//			months += 12 - int(fromDate.Month())
//			continue
//		}
//
//		// Soma 12 meses para cada ano.
//		if y < toDate.Year() && y > fromDate.Year() {
//			months += 12
//			continue
//		}
//
//		if y == toDate.Year() {
//			months += int(toDate.Month())
//		}
//	}
//
//	if fromDate.Day() > toDate.Day() {
//		return months - 1
//	}
//
//	return months
//}
//
///*func ElapsedYears(birthdate, today time.Time) int {
//	bday := birthdate.YearDay()
//
//	years := today.Year() - birthdate.Year()
//
//	leapy := func(d time.Time) bool {
//		year := d.Year()
//
//		return (year%400 == 0) || (year%100 != 0) || (year%4 == 0)
//	}
//
//	if (bday >= 60) && leapy(birthdate) && !leapy(today) {
//		return bday - 1
//	}
//
//	if (today.YearDay() >= 60) && !leapy(birthdate) && leapy(today) {
//		return bday + 1
//	}
//
//	if today.YearDay() < bday {
//		return years - 1
//	}
//
//	return years
//}
//*/
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
//// Between checks if param n is greather or equal to param low and lower than or equal param high
//func Between(n, low, high int) bool {
//	return n >= low && n <= high
//}
//
//func Tif(condition bool, tifThen, tifElse interface{}) interface{} {
//	if condition {
//		return tifThen
//	} else {
//		return tifElse
//	}
//}
//
//func ElapsedYears(from, to time.Time) int {
//	if from.IsZero() || to.IsZero() {
//		return 0
//	}
//
//	elapsedYears := to.Year() - from.Year()
//
//	if to.YearDay() < from.YearDay() {
//		elapsedYears--
//	}
//
//	return elapsedYears
//}
//
//func YearsAge(birthdate time.Time) int {
//	return ElapsedYears(birthdate, time.Now())
//}
//
//func Truncate( s string, maxLen int, trim bool) string {
//	if s=="" {
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
//	TransformFlagTrim                 = uint8(2)
//	TransformFlagLowerCase            = uint8(4)
//	TransformFlagUpperCase            = uint8(8)
//	TransformFlagOnlyDigits           = uint8(16)
//	TransformFlagOnlyLetters          = uint8(32)
//	TransformFlagOnlyLettersAndDigits = uint8(64)
//)
//
//func Transform( s string, maxLen int, transformFlags uint8) string {
//	if s=="" {
//		return s
//	}
//
//	if (transformFlags & TransformFlagOnlyLettersAndDigits) == TransformFlagOnlyLettersAndDigits {
//		s = OnlyLettersAndNumbers(s)
//
//		if s == "" {
//			return s
//		}
//	} else if (transformFlags & TransformFlagOnlyDigits) == TransformFlagOnlyDigits {
//		s = OnlyDigits(s)
//
//		if s == "" {
//			return s
//		}
//	} else if (transformFlags & TransformFlagOnlyLetters) == TransformFlagOnlyLetters {
//		s = OnlyAlpha(s)
//
//		if s == "" {
//			return s
//		}
//	}
//
//	// Have to trim before and after, to avoid issues with string truncation and new leading/trailing spaces
//	if (transformFlags & TransformFlagTrim) == TransformFlagTrim {
//		s = strings.TrimSpace(s)
//	}
//
//	if utf8.RuneCountInString(s) > maxLen {
//		s = string([]rune(s)[:maxLen])
//	}
//
//	// Have to trim before and after, to avoid issues with string truncation and new leading/trailing spaces
//	if ( transformFlags & TransformFlagTrim ) == TransformFlagTrim {
//		s = strings.TrimSpace(s)
//	}
//
//	if ( transformFlags & TransformFlagLowerCase) == TransformFlagLowerCase {
//		s = strings.ToLower(s)
//	}
//
//	if ( transformFlags & TransformFlagUpperCase) == TransformFlagUpperCase {
//		s = strings.ToUpper(s)
//	}
//
//	return s
//}
//
//func MatchesAny(search interface{}, items ...interface{}) bool {
//	for _, v := range items {
//		if fmt.Sprintf("%T", search) != fmt.Sprintf("%T", v) {
//			if search == v {
//				return true
//			}
//		}
//	}
//
//	return false
//}
//
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
