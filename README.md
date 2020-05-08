![GitHub](https://img.shields.io/github/license/mashape/apistatus.svg) ![GitHub](https://img.shields.io/badge/goDoc-Yes!-blue.svg) 
[![Go Report Card](https://goreportcard.com/badge/github.com/miguelpragier/handy?update)](https://goreportcard.com/report/github.com/miguelpragier/handy) 
[![Coverage Status](https://img.shields.io/badge/coverage-86.8%25-green.svg?style=flat)](https://gocover.io/github.com/miguelpragier/handy?version=1.10.x) [![Build Status](https://travis-ci.org/miguelpragier/handy.svg?branch=master)](https://travis-ci.org/miguelpragier/handy) [![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)
[![CircleCI](https://circleci.com/gh/miguelpragier/handy/tree/master.svg?style=svg)](https://circleci.com/gh/miguelpragier/handy/tree/master)

# Handy Go utilities
GO Golang Utilities and helpers like validators, sanitizers and string formatters


# Dependencies
None, since v1.1.1

___
Documentation :green_book:|
--------------------------|
[GoDocs](https://godoc.org/github.com/miguelpragier/handy)|
___

___
Quality :point_left:      |
--------------------------|
[GoCover](https://gocover.io/github.com/miguelpragier/handy)|
___


# Functions
```golang
// CheckEmail returns true if the given sequence is a valid email address
// See https://tools.ietf.org/html/rfc2822#section-3.4.1 for details about email address anatomy
func CheckEmail(email string) bool {}

// CheckNewPassword Run some basic checks on new password strings, based on given options
// This routine requires at least 4 (four) characters
// Example requiring only basic minimum lenght: CheckNewPassword("lalala", "lalala", 10, CheckNewPasswordComplexityLowest)
// Example requiring number and symbol: CheckNewPassword("lalala", "lalala", 10, CheckNewPasswordComplexityRequireNumber|CheckNewPasswordComplexityRequireSymbol)
func CheckNewPassword(password, passwordConfirmation string, minimumLenght uint, flagComplexity uint8) uint8 {

// StringHash simply generates a SHA256 hash from the given string
func StringHash(s string) string {}

// OnlyLetters returns only the letters from the given string, after strip all the rest ( numbers, spaces, etc. )
func OnlyLetters(sequence string) string {}

// OnlyDigits returns only the numbers from the given string, after strip all the rest ( letters, spaces, etc. )
func OnlyDigits(sequence string) string {}

// OnlyLettersAndNumbers returns only the letters and numbers from the given string, after strip all the rest, like spaces and special symbols.
func OnlyLettersAndNumbers(sequence string) string {}

// RandomInt returns a rondom integer within the given (inclusive) range
func RandomInt(min, max int) int {}

// StringAsFloat tries to convert a string to float, and if it can't, just returns zero
// It's limited to one billion
func StringAsFloat(s string, decimalSeparator, thousandsSeparator rune) float64 {}

// StringAsInteger returns the integer value extracted from string, or zero
func StringAsInteger(s string) int {}

// Between checks if param n in between low and high integer params
func Between(n, low, high int) bool {}

// Tif is a simple implementation of the dear ternary IF operator
func Tif(condition bool, tifThen, tifElse interface{}) interface{} {}

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
// 	TransformFlagHash - Apply handy.StringHash() routine to string
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

// IsNumericType checks if an interface's concrete type corresponds to some of golang native numeric types
func IsNumericType(x interface{}) bool {}

// Bit returns only uint8(0) or uint8(1).
// It receives an interface, and when it's a number, and when this number is 0 (zero) it returns 0. Otherwise it returns 1 (one)
// If the interface is not a number, it returns 0 (zero)
func Bit(x interface{}) uint8 {}

// Boolean returns the bool version/interpretation of some value;
// It receives an interface, and when this is a number, Boolean() returns flase of zero and true for different from zero.
// If it's a string, try to find "1", "T", "TRUE" to return true.
// Any other case returns false
func Boolean(x interface{}) bool {}

// Reverse returns the given string written backwards, with letters reversed.
func Reverse(s string) string {}

// CheckPersonName returns true if the name contains at least two words, one >= 3 chars and one >=2 chars.
// I understand that this is a particular criteria, but this is the OpenSourceMagic, where you can change and adapt to your own specs.
func CheckPersonName(name string, acceptEmpty bool) uint8 {}

// InArray searches for "item" in "array" and returns true if it's found
// This func resides here alone only because its long size.
// TODO Embrace/comprise all native scalar/primitive types
func InArray(array interface{}, item interface{}) bool {}

// ArrayDifferenceAtoB returns the items from A that doesn't exist in B
func ArrayDifferenceAtoB(a, b []int) []int {}

// ArrayDifference returns all items that doesn't exist in both given arrays
func ArrayDifference(a, b []int) []int {}


/* DateTime Routines */

// DateTimeAsString formats time.Time variables as strings, considering the format directive
func DateTimeAsString(dt time.Time, format string) string {}

// StringAsDateTime converts a date-time string using given format string and return it as time.Time
func StringAsDateTime(s string, format string) time.Time {}

// CheckDate validates a date using the given format
func CheckDate(format, dateTime string) bool {

// CheckDate returns true if given sequence is a valid date in format yyyymmdd
// The function removes non-digit characteres like "yyyy/mm/dd" or "yyyy-mm-dd", filtering to "yyyymmdd"
func CheckDateYMD(yyyymmdd string) bool {}

// YMDasDateUTC returns a valid UTC time from the given yyymmdd-formatted sequence
func YMDasDateUTC(yyyymmdd string, utc bool) (time.Time, error) {}

// YMDasDate returns a valid time from the given yyymmdd-formatted sequence
func YMDasDate(yyyymmdd string) (time.Time, error) {}

// ElapsedMonths returns the number of elapsed months between two given dates
func ElapsedMonths(from, to time.Time) int {}

// ElapsedYears returns the number of elapsed years between two given dates
func ElapsedYears(from, to time.Time) int {}

// YearsAge returns the number of years past since a given date
func YearsAge(birthdate time.Time) int {}

/* Brazilian specific routines */

// CheckCPF returns true if the given sequence is a valid cpf
// CPF is the Brazilian TAXPayerID document for persons
func CheckCPF(cpf string) bool {}

// CheckCNPJ returns true if the cnpj is valid
// Thanks to https://gopher.net.br/validacao-de-cpf-e-cnpj-em-go/
// CNPJ is the Brazilian TAXPayerID document for companies
func CheckCNPJ(cnpj string) bool {}

// AmountAsWord receives an int64 e returns the value as its text representation
// Today I have only the PT-BR text.
// Ex: AmountAsWord(129) => "cento e vinte e nove"
// Supports up to one trillion and does not add commas.
func AmountAsWord(n int64) string {}

/* Web/HTTP Specific Routines */

// HTTPRequestAsString gets a parameter coming from a http request as string, truncated to maxLenght
// Only maxLenght >= 1 is considered. Otherwise, it's ignored
func HTTPRequestAsString(r *http.Request, key string, maxLenght int, transformOptions ...uint8) string {}

// HTTPRequestAsInteger gets a parameter coming from a http request as an integer
// It tries to guess if it's a signed/negative integer
func HTTPRequestAsInteger(r *http.Request, key string) int {}

// HTTPRequestAsFloat64 gets a parameter coming from a http request as float64 number
// You have to inform the decimal separator symbol.
// If decimalSeparator is period, engine considers thousandSeparator is comma, and vice-versa.
func HTTPRequestAsFloat64(r *http.Request, key string, decimalSeparator rune) float64 {}

/* Other Routines  */

// CheckPersonNameResult returns a meaningful message describing the code generated bu CheckPersonName
// The routine considers the given idiom. The fallback is in english
func CheckPersonNameResult(idiom string, r uint8) string {}

// CheckNewPasswordResult returns a meaningful message describing the code generated bu CheckNewPassword()
// The routine considers the given idiom. The fallback is in english
func CheckNewPasswordResult(idiom string, r uint8) string {}

```
---

- [We're here as well at LIBHunt](https://go.libhunt.com/handy-alternatives)
- [We're also here at GOCover](https://gocover.io/github.com/miguelpragier/handy)
- [We're here too at GOLangCI](https://golangci.com/r/github.com/miguelpragier/handy)

