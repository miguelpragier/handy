package handy

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// CheckCPF returns true if the given input is a valid cpf
// CPF is the Brazilian TAXPayerID document for persons
func CheckCPF(cpf string) bool {
	// Se já chegar vazio, falha
	if cpf == "" {
		return false
	}

	// Sanitiza a string de modo agressivo, retirando qualquer runa que não seja dígito
	re, _ := regexp.Compile(`[\D]`)

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

	algProdCpfDig1 := make([]int, 12)

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

	var algProdCpfDig2 = make([]int, 13)

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

	return char13 == digit2
}

// AmountAsWord receives an int64 e returns the value as its text representation
// Today I have only the PT-BR text.
// Ex: AmountAsWord(129) => "cento e vinte e nove"
// Supports up to one trillion and does not add commas.
func AmountAsWord(n int64) string {
	if n == 0 {
		return "zero"
	}

	negativo := n < 0
	menos := ""

	if negativo {
		n *= -1
		menos = "menos "
	}

	const (
		umTrilhao = 1000000000000
		umBilhao  = 1000000000
		umMilhao  = 1000000
		mil       = 1000
		cem       = 100
		vinte     = 20

		conjuncaoE = " e "
	)

	if n < vinte {
		a := [...]string{"", "um", "dois", "três", "quatro", "cinco", "seis", "sete", "oito", "nove", "dez", "onze", "doze", "treze", "quatorze", "quinze", "dezesseis", "dezessete", "dezoito", "dezenove"}

		return fmt.Sprintf("%s%s", menos, a[n])
	}

	if n < cem {
		quantos := n / 10
		resto := n % 10

		a := [...]string{"", "", "vinte", "trinta", "quarenta", "cinquenta", "sessenta", "setenta", "oitenta", "noventa"}
		e := " "

		if resto > 0 {
			e = conjuncaoE
		}

		return strings.Replace(fmt.Sprintf("%s%s%s%s", menos, a[quantos], e, AmountAsWord(resto)), "  ", " ", -1)
	}

	if n < mil {
		if n == cem {
			return "cem"
		}

		quantos := n / cem
		resto := n % cem

		a := [...]string{"", "cento", "duzentos", "trezentos", "quatrocentos", "quinhentos", "seicentos", "setecentos", "oitocentos", "novecentos"}
		e := " "

		if resto > 0 {
			e = conjuncaoE
		}

		return strings.Replace(fmt.Sprintf("%s%s%s%s", menos, a[quantos], e, AmountAsWord(resto)), "  ", " ", -1)
	}

	if n < umMilhao {
		if n == mil {
			return "mil"
		}

		quantos := n / mil
		resto := n % mil
		e := " "

		if (resto > 0) && (resto < cem) {
			e = conjuncaoE
		}

		return strings.Replace(fmt.Sprintf("%s%s mil %s%s", menos, AmountAsWord(quantos), e, AmountAsWord(resto)), "  ", " ", -1)
	}

	if n < umBilhao {
		if n == umMilhao {
			return "um milhão"
		}

		quantos := n / umMilhao
		resto := n % umMilhao
		e := " "
		u := " milhão "

		if quantos > 1 {
			u = " milhões "
		}

		if (resto > 0) && (resto < mil) {
			e = " e "
		}

		return strings.Replace(fmt.Sprintf("%s%s%s%s%s", menos, AmountAsWord(quantos), u, e, AmountAsWord(resto)), "  ", " ", -1)
	}

	if n < umTrilhao {
		if n == umBilhao {
			return "um bilhão"
		}

		quantos := n / umBilhao
		resto := n % umBilhao
		e := " "
		u := " bilhão "

		if quantos > 1 {
			u = " bilhões "
		}

		if (resto > 0) && (resto < mil) {
			e = conjuncaoE
		}

		return strings.Replace(fmt.Sprintf("%s%s%s%s%s", menos, AmountAsWord(quantos), u, e, AmountAsWord(resto)), "  ", " ", -1)
	}

	if n == umTrilhao {
		return "um trilhão"
	}

	return fmt.Sprintf("%s%d", menos, n)
}
