package handy

import (
	"fmt"
	"os"
	"testing"
)

func TestEnvCheck(t *testing.T) {
	type test struct {
		title        string
		defaultValue string
		define       string
		mandatory    bool
		debugLog     bool
		assertion    error
	}

	tests := []test{
		{"ENVALRIGHT", "X", "xyz", true, false, nil},
		{"ENVDOESNTEXIST", "X", "", false, true, nil},
		{"ENVEMPTYOPTIONAL", "", "xyz", false, true, nil},
		{"ENVEMPTYMANDATORY", "", "", true, true, fmt.Errorf(`required environment variable "%s" isn't set`, "ENVEMPTYMANDATORY")},
		{"ENVMANDATORYWITHDEFAULT", "xyz", " ", true, true, nil},
	}

	for _, tx := range tests {
		if err := os.Setenv(tx.title, tx.define); err != nil {
			t.Fatal(err)
		}

		if err := EnvCheck(tx.title, tx.defaultValue, tx.mandatory, tx.debugLog); err != tx.assertion {
			if err != nil && tx.assertion != nil && tx.assertion.Error() != err.Error() {
				t.Logf("expected value %v, got %v\n", tx.assertion, err)
				t.Fail()
			}
		}
	}
}

func TestEnvStr(t *testing.T) {
	type test struct {
		envVar       string
		value        string
		defaultValue string
		assertion    string
	}

	tests := []test{
		{"ENV_XYZ", "xyz", "", "xyz"},
		{"ENV_EMPTY", "", "", ""},
		{"ENV_EMPTYDEFAULT", "", "xyz", "xyz"},
	}

	for _, tx := range tests {
		if err := os.Setenv(tx.envVar, tx.value); err != nil {
			t.Fatal(err)
		}

		if v := EnvStr(tx.envVar, tx.defaultValue); v != tx.assertion {
			t.Logf("expected value %v, got %v\n", tx.assertion, v)
			t.Fail()
		}
	}
}

//func TestEnvStrS(t *testing.T) {
//	if os.Getenv(key) != `` {
//		return strings.Split(os.Getenv(key), separator)
//	}
//
//	if len(defaultValue) > 0 {
//		return defaultValue
//	}
//
//	return []string{}
//}

func TestEnvInt(t *testing.T) {
	type test struct {
		envVar       string
		value        string
		defaultValue int
		assertion    int
	}

	tests := []test{
		{"ENV_111", "111", 0, 111},
		{"ENV_EMPTY", "", 0, 0},
		{"ENV_ZERO", "0", 0, 0},
		{"ENV_ZERODEFAULT", "", 0, 0},
		{"ENV_555", "", 555, 555},
	}

	for _, tx := range tests {
		if err := os.Setenv(tx.envVar, tx.value); err != nil {
			t.Fatal(err)
		}

		if v := EnvInt(tx.envVar, tx.defaultValue); v != tx.assertion {
			t.Logf("expected value %v, got %v\n", tx.assertion, v)
			t.Fail()
		}
	}
}

//func TestEnvInt64(t *testing.T) {
//	if os.Getenv(key) != `` {
//		if i, err := strconv.ParseInt(os.Getenv(key), 10, 64); err == nil {
//			return i
//		}
//	}
//
//	return defaultValue
//}
//
//func TestEnvIntS(t *testing.T) {
//	if os.Getenv(key) != `` {
//		a := strings.Split(os.Getenv(key), separator)
//
//		is := make([]int, len(a))
//
//		for i, x := range a {
//			is[i], _ = strconv.Atoi(x)
//		}
//
//		return is
//	}
//
//	if len(defaultValue) > 0 {
//		return defaultValue
//	}
//
//	return []int{}
//}
//
//func TestEnvFloat64(t *testing.T) {
//	if os.Getenv(key) != `` {
//		if f, err := strconv.ParseFloat(os.Getenv(key), 64); err == nil {
//			return f
//		}
//	}
//
//	return defaultValue
//}
//
//func TestEnvBool(t *testing.T) {
//	if os.Getenv(key) != `` {
//		if b, err := strconv.ParseBool(os.Getenv(key)); err == nil {
//			return b
//		}
//	}
//
//	return defaultValue
//}
