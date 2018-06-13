package handy

import (
	"frontline/jsonanswer"
	"net/http"
	"strconv"
	"strings"
)

func RequestAsString(r *http.Request, key string, maxLenght int) string {
	s := r.FormValue(key)

	if len([]rune(s)) >= maxLenght {
		return s[0:maxLenght]
	}

	return s
}

func RequestAsInteger(r *http.Request, key string) int {
	s := strings.TrimSpace(r.FormValue(key))

	neg := s[0:1] == "-"

	const BillionLength = 12

	if len([]rune(s)) > BillionLength {
		s = s[0:BillionLength]
	}

	s = OnlyDigits(s)

	if i, err := strconv.Atoi(s); err == nil {
		if neg {
			return i * (-1)
		}

		return i
	}

	return 0
}

func RequestAsFloat64(r *http.Request, key string, decimalSeparator rune) float64 {
	s := strings.TrimSpace(r.FormValue(key))

	thousandSeparator := Tif(decimalSeparator == ',', '.', ',').(rune)

	return StringAsFloat(s, decimalSeparator, thousandSeparator)
}

// Check Person Name
// GET /handy/check/personname/
func WebCheckPersonName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name := RequestAsString(r, "name", 200)

	if !CheckPersonName(name, false) {
		jsonanswer.Fail200(w, "Nomes devem conter ao menos 5 caracteres em duas palavras")
	}

	jsonanswer.Success(w, "Nome Válido")
}

// Check Email
// GET /handy/check/personname/
func WebCheckEmail(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := RequestAsString(r, "email", 200)

	if !CheckEmail(email) {
		jsonanswer.Fail200(w, "Email Inválido")
	}

	jsonanswer.Success(w, "Email Válido")
}

// Check Phone
// GET /handy/check/personname/
func WebCheckPhone(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	phone := RequestAsString(r, "phone", 15)

	if !CheckPhone(phone, false) {
		jsonanswer.Fail200(w, "Fone Inválido")
	}

	jsonanswer.Success(w, "Fone Válido")
}

// Check CPF
// GET /handy/check/personname/
func WebCheckCPF(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	cpf := RequestAsString(r, "cpf", 15)

	if !CheckCPF(cpf) {
		jsonanswer.Fail200(w, "CPF Inválido")
	}

	jsonanswer.Success(w, "CPF Válido")
}

// Get Age
// GET /handy/check/personname/
func WebCheckAge(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	yyyymmdd := RequestAsString(r, "birthdate", 200)

	if !CheckDate(yyyymmdd) {
		jsonanswer.Fail200(w, "Data de Nascimento Inválida")
	}

	if birthdate, err := YMDasDate(yyyymmdd); err == nil {
		age := YearsAge(birthdate)
		jsonanswer.Answer(w, true, map[string]int{"idade": age}, "ok", http.StatusOK)
	} else {
		jsonanswer.Fail200(w, "Não foi possível calcular a idade")
	}
}

// Get Amount in Words
// GET /handy/check/personname/
func WebCheckAmountInWords(w http.ResponseWriter, r *http.Request) {

}
