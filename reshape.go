package handy

import (
	"strings"
	"unicode/utf8"
)

const reshapePlaceholderDefault = '#'

// ReshapePH formats a string according positional directives given with an arbitrary placeholder
// It's the core of Reshape
func ReshapePH(placeholder rune, format, sequence string) string {
	if placeholder == 0 || format == "" || sequence == "" {
		return sequence
	}

	hashes := strings.Count(format, string(placeholder))

	positions := utf8.RuneCountInString(sequence)

	var (
		seqRunes = []rune(sequence)
		j        = 0
		sb       strings.Builder
	)

	for _, r := range format {
		if j >= positions {
			break
		}

		if r != placeholder {
			sb.WriteRune(r)
		} else {
			sb.WriteRune(seqRunes[j])
			j++
		}
	}

	if hashes < positions {
		sb.WriteString(sequence[hashes:])
	}

	return sb.String()
}

// Reshape formats a string according "hash" positional directives
// Example: Reshape("#####-###","98765432") returns "98765-432"
// Example: Reshape("####.####-####/##","98765432101234") returns "9876.5432-1012/34"
// It tries to adapt the format to the sequence when the sizes doesn't match
// Example: Reshape("###.###.##","9876") returns "987.6"
// Example: Reshape("###.###","123456789") returns "123.456789"
// Run: https://play.golang.org/p/hlCyVZehTn3
func Reshape(format, sequence string) string {
	return ReshapePH(reshapePlaceholderDefault, format, sequence)
}
