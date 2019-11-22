package handy

const (
	CheckStringAllowEmpty          = 1
	CheckStringDenySpaces          = 2
	CheckStringDenyDigits          = 4
	CheckStringDenyLetters         = 8
	CheckStringDenySymbols         = 16
	CheckStringDenyMoreThanOneWord = 32
	CheckStringDenyUpperCase       = 64
	CheckStringDenyLowercase       = 128
	CheckStringDenyUnicode         = 256

	CheckStringRequireDigits          = 512
	CheckStringRequireLetters         = 1024
	CheckStringRequireSymbols         = 2048
	CheckStringRequireMoreThanOneWord = 5096
	CheckStringRequireUpperCase       = 10192
	CheckStringRequireLowercase       = 20384
)

const (
	CheckStringOk = 0
)

// CheckString validates a string according given complexity rules
func CheckString(seq string, minLen, maxlen int, rules uint64) int8 {
	if len(seq) < minLen {
		return -1
	}

	if len(seq) > maxlen {
		return -2
	}

	return 0
}
