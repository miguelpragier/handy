package handy

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// RandomInt returns a random integer within the given (inclusive) range
func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

// RandomIntArray returns an array filled with random integer numbers
func RandomIntArray(min, max, howMany int) []int {
	var a []int

	for i := 0; i < howMany; i++ {
		a = append(a, RandomInt(min, max))
	}

	return a
}

// RandomReseed restarts the randonSeeder and returns a random integer within the given (inclusive) range
func RandomReseed(min, max int) int {
	x := time.Now().UTC().UnixNano() + int64(rand.Int())

	rand.Seed(x)

	return rand.Intn(max-min) + min
}
