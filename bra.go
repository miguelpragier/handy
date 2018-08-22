// These file contains routines to handle Brazilian-specific rules
package handy

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// CheckCPF returns true if the given sequence is a valid cpf
// CPF is the Brazilian TAXPayerID document for persons
func CheckCPF(cpf string) bool {
	// Se já chegar vazio, falha
	if cpf == "" {
		return false
	}

	// Sanitiza a string de modo agressivo, retirando qualquer runa que não seja dígito
	re, _ := regexp.Compile("[\\D]")

	cpf = re.ReplaceAllString(cpf, "")

	// Se o comprimento da string estiver diferente de 11, falhar
	if len(cpf) != 11 {
		return false
	}

	// Testa seqüências de 11 dígitos iguais, cujo cálculo é válido mas são inaceitáveis como documento.
	for i := 0; i <= 9; i++ {
		if cpf == strings.Repeat(fmt.Sprintf("%d", i), 11) {
			return false
		}
	}

	intval := func(b byte) int {
		i, _ := strconv.Atoi(string(b))

		return i
	}

	soma1 := intval(cpf[0])*10 + intval(cpf[1])*9 + intval(cpf[2])*8 + intval(cpf[3])*7 + intval(cpf[4])*6 + intval(cpf[5])*5 + intval(cpf[6])*4 + intval(cpf[7])*3 + intval(cpf[8])*2

	resto1 := (soma1 * 10) % 11

	if resto1 == 10 {
		resto1 = 0
	}

	soma2 := intval(cpf[0])*11 + intval(cpf[1])*10 + intval(cpf[2])*9 + intval(cpf[3])*8 + intval(cpf[4])*7 + intval(cpf[5])*6 + intval(cpf[6])*5 + intval(cpf[7])*4 + intval(cpf[8])*3 + intval(cpf[9])*2

	resto2 := (soma2 * 10) % 11

	if resto2 == 10 {
		resto2 = 0
	}

	return resto1 == intval(cpf[9]) && resto2 == intval(cpf[10])
}

// CheckCNPJ returns true if the cnpj is valid
// Thanks to https://gopher.net.br/validacao-de-cpf-e-cnpj-em-go/
// CNPJ is the Brazilian TAXPayerID document for companies
func CheckCNPJ(cnpj string) bool {
	cnpj = OnlyDigits(cnpj)

	if len(cnpj) != 14 {
		return false
	}

	algs := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	algProdCpfDig1 := make([]int, 12, 12)

	for key, val := range algs {
		intParsed, _ := strconv.Atoi(string(cnpj[key]))
		sumTmp := val * intParsed
		algProdCpfDig1[key] = sumTmp
	}

	sum := 0

	for _, val := range algProdCpfDig1 {
		sum += val
	}

	digit1 := sum % 11

	if digit1 < 2 {
		digit1 = 0
	} else {
		digit1 = 11 - digit1
	}

	char12, _ := strconv.Atoi(string(cnpj[12]))

	if char12 != digit1 {
		return false
	}

	algs = append([]int{6}, algs...)

	var algProdCpfDig2 = make([]int, 13, 13)

	for key, val := range algs {
		intParsed, _ := strconv.Atoi(string(cnpj[key]))
		sumTmp := val * intParsed
		algProdCpfDig2[key] = sumTmp
	}

	sum = 0

	for _, val := range algProdCpfDig2 {
		sum += val
	}

	digit2 := sum % 11

	if digit2 < 2 {
		digit2 = 0
	} else {
		digit2 = 11 - digit2
	}

	char13, _ := strconv.Atoi(string(cnpj[13]))

	if char13 != digit2 {
		return false
	}

	return true
}
