package handy

import (
	"testing"
	"unicode/utf8"
)

func TestRandomString(t *testing.T) {
	emptyString := ""

	testBunch := []struct {
		title                                                 string
		minLen, maxLen                                        int
		allowUnicode, allowNumbers, allowSymbols, allowSpaces bool
		expectedString                                        *string
	}{
		{
			title:          "zeroed len",
			minLen:         0,
			maxLen:         0,
			allowUnicode:   false,
			allowNumbers:   false,
			allowSymbols:   false,
			allowSpaces:    false,
			expectedString: &emptyString,
		},
		{
			title:        "zeroed minLen",
			minLen:       0,
			maxLen:       1,
			allowUnicode: false,
			allowNumbers: false,
			allowSymbols: false,
			allowSpaces:  false,
		},
		{
			title:          "zeroedMaxLen",
			minLen:         1,
			maxLen:         0,
			allowUnicode:   false,
			allowNumbers:   false,
			allowSymbols:   false,
			allowSpaces:    false,
			expectedString: &emptyString,
		},
		{
			title:          "minLen>maxLen",
			minLen:         2,
			maxLen:         1,
			allowUnicode:   false,
			allowNumbers:   false,
			allowSymbols:   false,
			allowSpaces:    false,
			expectedString: &emptyString,
		},
		{
			title:        "len=1",
			minLen:       1,
			maxLen:       1,
			allowUnicode: false,
			allowNumbers: false,
			allowSymbols: false,
			allowSpaces:  false,
		},
		{
			title:        "full possibilities fixed len=32",
			minLen:       32,
			maxLen:       32,
			allowUnicode: true,
			allowNumbers: true,
			allowSymbols: true,
			allowSpaces:  true,
		},
		{
			title:        "fixed len=32, disallow unicode",
			minLen:       32,
			maxLen:       32,
			allowUnicode: false,
			allowNumbers: true,
			allowSymbols: true,
			allowSpaces:  true,
		},
		{
			title:        "fixed len=32, disallow unicode and numbers",
			minLen:       32,
			maxLen:       32,
			allowUnicode: false,
			allowNumbers: false,
			allowSymbols: true,
			allowSpaces:  true,
		},
		{
			title:        "fixed len=32, disallow unicode, numbers and symbols",
			minLen:       32,
			maxLen:       32,
			allowUnicode: false,
			allowNumbers: false,
			allowSymbols: false,
			allowSpaces:  false,
		},
		{
			title:        "fixed len=32, disallow all",
			minLen:       32,
			maxLen:       32,
			allowUnicode: false,
			allowNumbers: false,
			allowSymbols: false,
			allowSpaces:  false,
		},
	}

	for _, tx := range testBunch {
		t.Run(tx.title, func(t *testing.T) {
			s := RandomString(tx.minLen, tx.maxLen, tx.allowUnicode, tx.allowNumbers, tx.allowSymbols, tx.allowSpaces)

			l := utf8.RuneCountInString(s)

			switch {
			case tx.maxLen == 0:
				if l > 0 {
					t.Errorf("expected len is 0, but got %d", l)
				}

			case tx.maxLen > tx.minLen:
				if l < tx.minLen || l > tx.maxLen {
					t.Errorf("expected len between %d and %d, but got %d", tx.minLen, tx.maxLen, l)
				}

			case tx.maxLen == tx.minLen:
				if l != tx.maxLen {
					t.Errorf("expected len is %d, but got %d", tx.maxLen, l)
				}
			}

			if tx.expectedString != nil {
				if *tx.expectedString != s {
					t.Errorf(`expected string "%s", but got %s`, *tx.expectedString, s)
				}
			}
		})
	}
}
