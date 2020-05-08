package handy

import (
	"bytes"
	"fmt"
	"math"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func Test_HTTPRequestAsString(t *testing.T) {
	const habemosPapa = "[9,8,7,5]"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Run("testing parsed value", func(t *testing.T) {
			tr := HTTPRequestAsString(r, "status", 50)

			if tr != habemosPapa {
				t.Errorf("Test has failed!\n\tExpected: %s, \n\tGot: %v", habemosPapa, tr)
			}
		})
	}))

	defer ts.Close()

	if _, err := http.PostForm(ts.URL, url.Values{"status": {habemosPapa}}); err != nil {
		t.Errorf("Test has failed! %v", err)
	}

	if _, err := http.Get(fmt.Sprintf("%s?status=%s", ts.URL, habemosPapa)); err != nil {
		t.Errorf("Test has failed! %v", err)
	}
}

func Test_HTTPRequestAsInteger(t *testing.T) {
	const xBurger = 1234567890

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Run("testing parsed value", func(t *testing.T) {
			tr := HTTPRequestAsInteger(r, "status")

			if tr != xBurger {
				t.Errorf("Test has failed!\n\tExpected: %v, \n\tGot: %v", xBurger, tr)
			}
		})
	}))

	defer ts.Close()

	if _, err := http.PostForm(ts.URL, url.Values{"status": {fmt.Sprintf("%d", xBurger)}}); err != nil {
		t.Errorf("Test has failed! %v", err)
	}
}

func Test_HTTPRequestAsFloat64(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Run("testing parsed value", func(t *testing.T) {
			tr := HTTPRequestAsFloat64(r, "status", '.')

			if tr != math.Pi {
				t.Errorf("Test has failed!\n\tExpected: %v, \n\tGot: %v", math.Pi, tr)
			}
		})
	}))

	defer ts.Close()

	if _, err := http.PostForm(ts.URL, url.Values{"status": {fmt.Sprintf("%.15f", math.Pi)}}); err != nil {
		t.Errorf("Test has failed! %v", err)
	}
}

func Test_HTTPJSONBodyToStruct(t *testing.T) {
	const (
		testName = "Forty Two"
		testAge  = 42
	)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Run("testing parsed value", func(t *testing.T) {
			var testDummy struct {
				Name string `json:"name"`
				Age  int    `json:"age"`
			}

			tr := HTTPJSONBodyToStruct(r, &testDummy)

			if !tr || ((testDummy.Name != testName) || (testDummy.Age != testAge)) {
				t.Errorf("Test has failed!\n\tExpected: (%v,%v)\n\tGot: %v", testName, testAge, testDummy)
			}
		})
	}))

	defer ts.Close()

	j := []byte(fmt.Sprintf(`{"name":"%s","age":%d}`, testName, testAge))

	if _, err := http.Post(ts.URL, "application/json;charset=utf-8", bytes.NewBuffer(j)); err != nil {
		t.Errorf("Test has failed! %v", err)
	}
}

func TestHTTPJSONToStruct(t *testing.T) {
	const (
		testName  = "Forty Two"
		testAge   = 42
		closeBody = true
	)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Run("testing parsed value", func(t *testing.T) {
			var testDummy struct {
				Name string `json:"name"`
				Age  int    `json:"age"`
			}

			if err := HTTPJSONToStruct(r, &testDummy, closeBody); err != nil {
				t.Error(err)
			}

			if (testDummy.Name != testName) || (testDummy.Age != testAge) {
				t.Errorf("Test has failed!\n\tExpected: (%v,%v)\n\tGot: %v", testName, testAge, testDummy)
			}
		})
	}))

	defer ts.Close()

	j := []byte(fmt.Sprintf(`{"name":"%s","age":%d}`, testName, testAge))

	if _, err := http.Post(ts.URL, "application/json;charset=utf-8", bytes.NewBuffer(j)); err != nil {
		t.Errorf("Test has failed! %v", err)
	}
}
