package handy

// InArray searches for "item" in "array" and returns true if it's found
// This func resides here alone only because its long size.
// TODO Embrace/comprise all native scalar/primitive types
func InArray(array interface{}, item interface{}) bool {
	switch array.(type) {
	case []int:
		a, _ := array.([]int)

		if len(a) < 1 {
			return false
		}

		i, ok := item.(int)

		if !ok {
			return false
		}

		for _, x := range a {
			if x == i {
				return true
			}
		}

	case []int8:
		a, _ := array.([]int8)

		if len(a) < 1 {
			return false
		}

		i, ok := item.(int8)

		if !ok {
			return false
		}

		for _, x := range a {
			if x == i {
				return true
			}
		}

	case []int16:
		a, _ := array.([]int16)

		if len(a) < 1 {
			return false
		}

		i, ok := item.(int16)

		if !ok {
			return false
		}

		for _, x := range a {
			if x == i {
				return true
			}
		}

	case []int32: // Works for int32 and rune types
		a, _ := array.([]int32)

		if len(a) < 1 {
			return false
		}

		i, ok := item.(int32)

		if !ok {
			return false
		}

		for _, x := range a {
			if x == i {
				return true
			}
		}

	case []int64:
		a, _ := array.([]int64)

		if len(a) < 1 {
			return false
		}

		i, ok := item.(int64)

		if !ok {
			return false
		}

		for _, x := range a {
			if x == i {
				return true
			}
		}

	case []uint:
		a, _ := array.([]uint)

		if len(a) < 1 {
			return false
		}

		i, ok := item.(uint)

		if !ok {
			return false
		}

		for _, x := range a {
			if x == i {
				return true
			}
		}

	case []uint8: // Works for uint8 and byte types
		a, _ := array.([]uint8)

		if len(a) < 1 {
			return false
		}

		i, ok := item.(uint8)

		if !ok {
			return false
		}

		for _, x := range a {
			if x == i {
				return true
			}
		}

	case []uint16:
		a, _ := array.([]uint16)

		if len(a) < 1 {
			return false
		}

		i, ok := item.(uint16)

		if !ok {
			return false
		}

		for _, x := range a {
			if x == i {
				return true
			}
		}

	case []uint32:
		a, _ := array.([]uint32)

		if len(a) < 1 {
			return false
		}

		i, ok := item.(uint32)

		if !ok {
			return false
		}

		for _, x := range a {
			if x == i {
				return true
			}
		}

	case []uint64: // Works for uint8 and byte types
		a, _ := array.([]uint64)

		if len(a) < 1 {
			return false
		}

		i, ok := item.(uint64)

		if !ok {
			return false
		}

		for _, x := range a {
			if x == i {
				return true
			}
		}

	case []float32:
		a, _ := array.([]float32)

		if len(a) < 1 {
			return false
		}

		f, ok := item.(float32)

		if !ok {
			return false
		}

		for _, x := range a {
			if x == f {
				return true
			}
		}

	case []float64:
		a, _ := array.([]float64)

		if len(a) < 1 {
			return false
		}

		f, ok := item.(float64)

		if !ok {
			return false
		}

		for _, x := range a {
			if x == f {
				return true
			}
		}

	case []string:
		a, _ := array.([]string)

		if len(a) < 1 {
			return false
		}

		s, ok := item.(string)

		if !ok {
			return false
		}

		for _, x := range a {
			if x == s {
				return true
			}
		}
	}

	return false
}
