package handy

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// HTTPRequestAsString gets a parameter coming from a http request as string, truncated to maxLength
// Only maxLength >= 1 is considered. Otherwise, it's ignored
func HTTPRequestAsString(r *http.Request, key string, maxLength int, transformOptions ...uint8) string {
	if err := r.ParseForm(); err != nil {
		return ""
	}

	s := r.FormValue(key)

	if s == "" {
		vars := mux.Vars(r)

		var ok bool

		if s, ok = vars[key]; !ok {
			return ""
		}
	}

	if len(transformOptions) > 0 {
		s = Transform(s, maxLength, transformOptions[0])
	}

	if s == "" {
		return ""
	}

	if (maxLength > 0) && (len([]rune(s)) >= maxLength) {
		return s[0:maxLength]
	}

	return s
}

// HTTPRequestAsInteger gets a parameter coming from a http request as an integer
// It tries to guess if it's a signed/negative integer
func HTTPRequestAsInteger(r *http.Request, key string) int {
	if err := r.ParseForm(); err != nil {
		return 0
	}

	s := r.FormValue(key)

	if s == "" {
		vars := mux.Vars(r)

		var ok bool

		if s, ok = vars[key]; !ok {
			return 0
		}
	}

	neg := s[0:1] == "-"

	i := StringAsInteger(s)

	if neg && (i > 0) {
		return i * -1
	}

	return i
}

// HTTPRequestAsFloat64 gets a parameter coming from a http request as float64 number
// You have to inform the decimal separator symbol.
// If decimalSeparator is period, engine considers thousandSeparator is comma, and vice-versa.
func HTTPRequestAsFloat64(r *http.Request, key string, decimalSeparator rune) float64 {
	if err := r.ParseForm(); err != nil {
		return 0
	}

	s := r.FormValue(key)

	if s == "" {
		vars := mux.Vars(r)

		var ok bool

		if s, ok = vars[key]; !ok {
			return 0
		}
	}

	thousandSeparator := Tif(decimalSeparator == ',', '.', ',').(rune)

	return StringAsFloat(s, decimalSeparator, thousandSeparator)
}

// HTTPJSONBodyToStruct decode json to a given anatomically compatible struct
func HTTPJSONBodyToStruct(r *http.Request, targetStruct interface{}) bool {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(targetStruct)

	return err == nil
}
