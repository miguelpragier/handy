package handy

import (
	"math/big"
)

func intToBigint(i interface{}) *big.Int {
	bi := big.NewInt(0)

	switch x := i.(type) {
	case int:
		bi.SetInt64(int64(x))
	case int8:
		bi.SetInt64(int64(x))
	case int16:
		bi.SetInt64(int64(x))
	case int32:
		bi.SetInt64(int64(x))
	case int64:
		bi.SetInt64(x)
	case uint:
		bi.SetUint64(uint64(x))
	case uint8:
		bi.SetUint64(uint64(x))
	case uint16:
		bi.SetUint64(uint64(x))
	case uint32:
		bi.SetUint64(uint64(x))
	case uint64:
		bi.SetUint64(x)
	}

	return bi
}

// InArrayIntFlex returns true if "item" exists in "array"
// item and array can be of different kinds, since they're integer types
// array param should be an array of any integer type (int,int8,int16,int32,int64,uint,uint8,uint16,uint32 and uint64)
// The function uses bigInt type convertions previous to values comparison
func InArrayIntFlex(item interface{}, array interface{}) bool {
	if array == nil || item == nil {
		return false
	}

	b1 := intToBigint(item)

	b2 := big.NewInt(0)

	switch a := array.(type) {
	case []int:
		for _, x := range a {
			b2.SetInt64(int64(x))

			if b1.Cmp(b2) == 0 {
				return true
			}
		}

	case []int8:
		for _, x := range a {
			b2.SetInt64(int64(x))

			if b1.Cmp(b2) == 0 {
				return true
			}
		}

	case []int16:
		for _, x := range a {
			b2.SetInt64(int64(x))

			if b1.Cmp(b2) == 0 {
				return true
			}
		}

	case []int32:
		for _, x := range a {
			b2.SetInt64(int64(x))

			if b1.Cmp(b2) == 0 {
				return true
			}
		}

	case []int64:
		for _, x := range a {
			b2.SetInt64(x)

			if b1.Cmp(b2) == 0 {
				return true
			}
		}

	case []uint:
		for _, x := range a {
			b2.SetUint64(uint64(x))

			if b1.Cmp(b2) == 0 {
				return true
			}
		}

	case []uint8:
		for _, x := range a {
			b2.SetUint64(uint64(x))

			if b1.Cmp(b2) == 0 {
				return true
			}
		}

	case []uint16:
		for _, x := range a {
			b2.SetUint64(uint64(x))

			if b1.Cmp(b2) == 0 {
				return true
			}
		}

	case []uint32:
		for _, x := range a {
			b2.SetUint64(uint64(x))

			if b1.Cmp(b2) == 0 {
				return true
			}
		}

	case []uint64:
		for _, x := range a {
			b2.SetUint64(x)

			if b1.Cmp(b2) == 0 {
				return true
			}
		}
	}

	return false
}

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

	case []bool:
		a, _ := array.([]bool)

		if len(a) < 1 {
			return false
		}

		s, ok := item.(bool)

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
