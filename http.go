package handy

import (
	"encoding/json"
	"log"
	"net/http"
)

// HTTPRequestAsString gets a parameter coming from a http request as string, truncated to maxLength
// Only maxLength >= 1 is considered. Otherwise, it's ignored
func HTTPRequestAsString(r *http.Request, key string, maxLength int, transformOptions ...uint) string {
	if err := r.ParseForm(); err != nil {
		log.Println(r.RequestURI, err)
		return ""
	}

	s := r.FormValue(key)

	if s == "" {
		s = r.URL.Query().Get(key)

		if s == "" {
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
		s = r.URL.Query().Get(key)

		if s == "" {
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
		s = r.URL.Query().Get(key)

		if s == "" {
			return 0
		}
	}

	thousandSeparator := Tif(decimalSeparator == ',', '.', ',').(rune)

	return StringAsFloat(s, decimalSeparator, thousandSeparator)
}

// HTTPJSONBodyToStruct decode json to a given anatomically compatible struct
// THIS ROUTINE IS BEEN DEPRECATED. Use HTTPJSONToStruct() instead.
func HTTPJSONBodyToStruct(r *http.Request, targetStruct interface{}) bool {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(targetStruct)

	return err == nil
}

// HTTPJSONToStruct decode json to a given anatomically compatible struct
// the differences to HTTPJSONBodyToStruct is that:
// - HTTPJSONToStruct can condittionally close body after unmarshalling
// - HTTPJSONToStruct returns an error instead of a bool
func HTTPJSONToStruct(r *http.Request, targetStruct interface{}, closeBody bool) error {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(targetStruct)

	if closeBody {
		defer func() {
			if err0 := r.Body.Close(); err0 != nil {
				log.Println(err0)
			}
		}()
	}

	return err
}

// HTTPAnswerJSON converts the given data as json, set the content-type header and write it to requester
func HTTPAnswerJSON(w http.ResponseWriter, data interface{}) error {
	var jb []byte

	if j1, ok := data.(string); ok {
		jb = []byte(j1)
	} else {
		if j2, err := json.Marshal(data); err != nil {
			return err
		} else {
			jb = j2
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if _, errw := w.Write(jb); errw != nil {
		return errw
	}

	return nil
}
