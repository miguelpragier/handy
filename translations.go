package handy

// CheckPersonNameResult returns a meaningful message describing the code generated bu CheckPersonName
// The routine considers the given idiom. The fallback is in english
func CheckPersonNameResult(idiom string, r uint8) string {
	switch idiom {
	case "bra":
		switch r {
		case CheckPersonNameResultOK:
			return "Nome Válido"
		case CheckPersonNameResultPolluted:
			return "O campo nome permite apenas letras e espaços"
		case CheckPersonNameResultTooFewWords:
			return "O nome deve ser composto de ao menos duas palavras"
		case CheckPersonNameResultTooShort:
			return "O nome deve ser composto de ao menos duas palavras, sendo uma com três e outra com ao menos duas letras"
		case CheckPersonNameResultTooSimple:
			return "Nome muito curto ou vazio"
		default:
			return "Erro desconhecido"
		}
	default:
		switch r {
		case CheckPersonNameResultOK:
			return "Name is well formed"
		case CheckPersonNameResultPolluted:
			return "Name accepts only letters and spaces"
		case CheckPersonNameResultTooFewWords:
			return "Name should be composed by at least two words"
		case CheckPersonNameResultTooShort:
			return "Name should be composed by at least two words, been one with 2 and the other with at least 3 letters"
		case CheckPersonNameResultTooSimple:
			return "Name too short or empty"
		default:
			return "Unknow error"
		}
	}
}

// CheckNewPasswordResult returns a meaningful message describing the code generated bu CheckNewPassword()
// The routine considers the given idiom. The fallback is in english
func CheckNewPasswordResult(idiom string, r uint8) string {
	if idiom == "bra" {
		switch r {
		case CheckNewPasswordResultOK:
			return "Senha Válida"
		case CheckNewPasswordResultDivergent:
			return "Senha diferente da confirmação"
		case CheckNewPasswordResultTooShort:
			return "Senha deve conter ao menos 6 caracteres, entre números e letras"
		case CheckNewPasswordResultTooSimple:
			return "Senha deve conter ao menos 6 caracteres, entre números e letras"
		default:
			return "Erro desconhecido"
		}
	} else {
		switch r {
		case CheckNewPasswordResultOK:
			return "Password Validated"
		case CheckNewPasswordResultDivergent:
			return "Password and confirmation doesn't match"
		case CheckNewPasswordResultTooShort:
			return "Password should contains at least 6 characters, mixing letters and numbers"
		case CheckNewPasswordResultTooSimple:
			return "Password should contains at least 6 characters, mixing letters and numbers"
		default:
			return "Unknow error"
		}
	}
}
