# Handy Go utilities
GO Golang Utilities and helpers like validators and string formatters


# Dependencies
None. It relies on standard library.


# ToDo list

* Add configuration engine for internationalization of messages - A json or yml file with mnemonics and message translations

* Translate all API documentaion/comments to english

* To consider some engine to - conditionally - preCompile all regex once at init() and store them in variables. The tradeoff would be the initial loadWeight, specially painful if you do not use those regex functions often

* Add config for rules, to customize passwords and other strings checking/validation

* Assure that all functions are covered by automated tests

* Add those cool badges to repository


# Functions
```golang
// CheckPersonName returns true if the name contains at least two words, one >= 3 chars and one >=2 chars.
// I understand that this is a particular criteria, but this is the OpenSourceMagic, where you can change and adapt to your own specs.
func CheckPersonName(name string, acceptEmpty bool) bool {}

// CheckCompanyName returns true if the name contains at least two words, digits allowed, one >= 3 chars and one >=2 chars.
// The main difference from CheckpersonName is that CheckCompanyName accept numbers/digits.
// I understand that this is a particular criteria, but this is the OpenSourceMagic, where you can change and adapt to your own specs.
func CheckCompanyName(name string, acceptEmpty bool) bool {}

// CheckCPF returns true if the given sequence is a valid cpf
// CPF is the Brazilian TAXPayerID document for persons
func CheckCPF(cpf string) bool {}

// CheckCNPJ returns true if the cnpj is valid
// Thanks to https://gopher.net.br/validacao-de-cpf-e-cnpj-em-go/
// CNPJ is the Brazilian TAXPayerID document for companies
func CheckCNPJ(cnpj string) bool {}

// CheckEmail returns true if the given sequence is a valid email address
// See https://tools.ietf.org/html/rfc2822#section-3.4.1 for details about email address anatomy
func CheckEmail(email string) bool {}

// CheckDate returns true if given sequence is a valid date in format yyyymmdd
// The function removes non-digit characteres like "yyyy/mm/dd" or "yyyy-mm-dd", filtering to "yyyymmdd"
func CheckDate(yyyymmdd string) bool {}

// AmountAsWord receives an int64 e returns the value as its text representation
// Today I have only the PT-BR text.
// Ex: AmountAsWord(129) => "cento e vinte e nove"
// Supports up to one trillion and does not add commas.
func AmountAsWord(n int64) string {}

// Run some basic checks on new password strings
// My rule requires at least six chars, with at least one letter and at least one number.
func CheckNewPassword(password, passwordConfirmation string) (bool, string) {}

// StringHash simply generates a SHA256 hash from the given string
func StringHash( sequence string ) string {}

// OnlyAlpha returns only the letters from the given string, after strip all the rest ( numbers, spaces, etc. )
func OnlyAlpha(sequence string) string {}

// OnlyDigits returns only the numbers from the given string, after strip all the rest ( letters, spaces, etc. )
func OnlyDigits(sequence string) string {}

// OnlyLettersAndNumbers returns only the letters and numbers from the given string, after strip all the rest, like spaces and special symbols.
func OnlyLettersAndNumbers(sequence string) string {}

// RandomInt returns a rondom integer within the given (inclusive) range
func RandomInt(min, max int) int {}

// CheckPhone returns true if a given sequence has between 9 and 14 digits
func CheckPhone(phone string, acceptEmpty bool) bool {}

// YMDasDateUTC returns a valid UTC time from the given yyymmdd-formatted sequence
func YMDasDateUTC(yyyymmdd string, utc bool) (time.Time, error) {}

// YMDasDate returns a valid time from the given yyymmdd-formatted sequence
func YMDasDate(yyyymmdd string) (time.Time, error) {}

// ElapsedMonths returns the number of elapsed months between two given dates
// I have to re-check the both versions and choose one to stay. If you have a good suggestion, just tell me!
func ElapsedMonths(fromDate, toDate time.Time) int {}

// ElapsedMonths2 returns the number of elapsed months between two given dates
// I have to re-check the both versions and choose one to stay. If you have a good suggestion, just tell me!
func ElapsedMonths2(fromDate, toDate time.Time) int {}

// StringAsFloat tries to convert a string to float, and if it can't, just returns zero
// It's limited to one billion
func StringAsFloat(s string, decimalSeparator, thousandsSeparator rune) float64 {}

// StringAsInteger returns the integer value extracted from string, or zero
func StringAsInteger(s string) int {}

// Between checks if param n is greather or equal to param low and lower than or equal param high
func Between(n, low, high int) bool {}

// Tif is a simple implementation of the dear ternary IF operator
func Tif(condition bool, tifThen, tifElse interface{}) interface{}

// ElapsedYears returns the number of elapsed years between two given dates
func ElapsedYears(from, to time.Time) int {}

// YearsAge returns the number of years past since a given date
func YearsAge(birthdate time.Time) int {}

// Truncate limits the length of a given string, trimming or not, according parameters
func Truncate(s string, maxLen int, trim bool) string {}

// Transform handles a string according given flags/parametrization, as follows:
// Available Flags to be used alone or combined:
//	TransformNone - Does nothing. It's only for truncation.
//	TransformFlagTrim - Trim spaces before and after proccess the input
//	TransformFlagLowerCase - Change string case to lower
//	TransformFlagUpperCase - Change string case to upper
//	TransformFlagOnlyDigits - Filter/strip all but digits
//	TransformFlagOnlyLetters - Filter/strip all but letters
//	TransformFlagOnlyLettersAndDigits - Filter/strip all but numbers and letters. Removes spaces, punctuation and special symbols
func Transform(s string, maxLen int, transformFlags uint8) string {}

// MatchesAny returns true if any of the given items matches ( equals ) the subject ( search parameter )
func MatchesAny(search interface{}, items ...interface{}) bool {}

// HasOnlyNumbers returns true if the sequence is entirely numeric
func HasOnlyNumbers(sequence string) bool {}

// HasOnlyNumbers returns true if the sequence is entirely composed by letters
func HasOnlyLetters(sequence string) bool {}

// TrimLen returns the runes count after trim the spaces
func TrimLen(text string) int {}

// CheckMinLen verifies if the rune-count is greater then or equal the given minimum
// It returns true if the given string has length greater than or equal than minLength parameter
func CheckMinLen(value string, minLength int) bool {}

