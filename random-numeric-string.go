package handy

import (
	"strconv"
	"strings"
)

// RandomNumericString returns a string with length between given lengthMin and lengthMax
// Any digit within forbiddenDigits param will be ignored
// Example: https://play.golang.org/p/phF-y9ZsUIP
func RandomNumericString(forbiddenDigits []int, lengthMin, lengthMax int) string {
	length := 0

	switch {
	case lengthMax < lengthMin:
		//panic(fmt.Sprintf("handy.RandomNumericString() lengthMax %d is smaller than lengthMix %d", lengthMax, lengthMin))
		// Instead of panic, return empty
		return ""

	case lengthMax == lengthMin:
		length = lengthMax

	default:
		length = RandomInt(lengthMin, lengthMax)
	}

	if length == 0 {
		return ""
	}

	allowedDigits := map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true}

	if len(forbiddenDigits) > 0 {
		for _, k := range forbiddenDigits {
			allowedDigits[k] = false
		}

		// Check if any digit remains allowed
		var remainingAllowed []int

		for k, b := range allowedDigits {
			if b {
				remainingAllowed = append(remainingAllowed, k)
			}
		}

		if len(remainingAllowed) == 0 {
			return ""
		}

		if len(remainingAllowed) == 1 {
			return strings.Repeat(strconv.Itoa(remainingAllowed[0]), length)
		}
	}

	s := make([]string, length)

	for i := 0; i < length; i++ {
		for {
			x := RandomInt(0, 9)

			if allowedDigits[x] {
				s[i] = strconv.Itoa(x)
				break
			}
		}
	}

	return strings.Join(s, "")
}
