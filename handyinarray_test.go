package handy

import "testing"

func TestInArrayFlex(t *testing.T) {
	tenThousandOcurrences := RandomIntArray(0, 100000, 10000)
	const greatestUint64 uint64 = 18446744073709551615

	testlist := []struct {
		summary        string
		item           interface{}
		array          interface{}
		expectedOutput bool
	}{
		{"nil array and item", nil, nil, false},
		{"nil array", 0, nil, false},
		{"nil item", 1, nil, false},
		{"item doesn't exist", 0, []int{-1, -2, 1, 2, 3, 4, 5, 6}, false},
		{"int16 checked", 32767, []int16{-32767, 0, 32767}, true},
		{"string array failing", 0, []string{"-1", "-0", "7455454", "xxx", "0xAF"}, false},
		{"ten thousand items", -100, tenThousandOcurrences, false},
		{"uint64 doesn't exist", greatestUint64, []uint64{1, 2, 3, 4, 5, 6}, false},
		{"greatest uint64 value exists", greatestUint64, []uint64{0, greatestUint64}, true},
		{"int exists", uint8(55), []int{-1, -2, 1, 2, 3, 4, 5, 6, 55}, true},
		{"int8 exists", uint64(255), []int8{-1, -2, 1, 2, 3, 4, 5, 6, 126}, false},
		{"int32 exists", 0, []int32{-1, -2, 1, 2, 3, 4, 5, 6}, false},
		{"int64 empty", 1, []int64{}, false},
		{"uint exists", 13, []uint{13}, true},
		{"uint8 doesn't exist", greatestUint64, []uint8{1, 2, 3, 4, 5, 6}, false},
		{"uint16 exist", 5555, []uint16{1, 2, 3, 4, 5, 6, 5555}, true},
		{"uint32 doesn't exist", -5555, []uint32{0, 0, 0, 0, 0, 5, 6, 5555}, false},
	}

	for _, tst := range testlist {
		t.Run(tst.summary, func(t *testing.T) {
			tr := InArrayIntFlex(tst.item, tst.array)

			if tr != tst.expectedOutput {
				t.Errorf("Failed! Send %d for %#v, expected %t got %t", tst.item, tst.array, tst.expectedOutput, tr)
			}
		})
	}
}
