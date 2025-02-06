package money

type Currency struct {
	code      string
	precision byte
}

// is returned when the currency to parse isn't a 3 letter code
const ErrInvalidCurrencyCode = Error("invalida currency code")

func ParseCurrency(code string) (Currency, error) {
	if len(code) != 3 {
		return Currency{}, ErrInvalidCurrencyCode
	}

	for _, char := range code {
		if char < 'A' || char > 'Z' {
			return Currency{}, ErrInvalidCurrencyCode
		}
	}

	switch code {
	case "IRR":
		return Currency{code: code, precision: 0}, nil
	case "CNY", "VND":
		return Currency{code: code, precision: 1}, nil
	case "BHD", "IQD", "KWD", "LYD", "OMR", "TND":
		return Currency{code: code, precision: 3}, nil
	default:
		return Currency{code: code, precision: 2}, nil
	}
}
