package handy

func DateStrCheckErrMessage(idiom string, errCode DateStrCheck) string {
	if idiom == "bra" {
		switch errCode {
		case DateStrCheckErrInvalid:
			return "data ou formato inválido"
		case DateStrCheckErrOutOfRange:
			return "data fora do intervalo permitido"
		case DateStrCheckErrEmpty:
			return "data não definida"
		default:
			return ""
		}
	}

	switch errCode {
	case DateStrCheckErrInvalid:
		return "date or format invalid"
	case DateStrCheckErrOutOfRange:
		return "date out of range"
	case DateStrCheckErrEmpty:
		return "date undefined"
	default:
		return ""
	}
}
