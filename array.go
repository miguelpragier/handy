package handy

// ArrayDifferenceAtoB returns the items from A that doesn't exist in B
func ArrayDifferenceAtoB(a, b []int) []int {
	m := make(map[int]bool)

	for _, item := range b {
		m[item] = true
	}

	var diff []int

	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}

	return diff
}

// ArrayDifference returns all items that doesn't exist in both given arrays
func ArrayDifference(a, b []int) []int {
	diffA := ArrayDifferenceAtoB(a, b)

	diffB := ArrayDifferenceAtoB(b, a)

	return append(diffB, diffA...)
}
