package money

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	ErrInvalidDecimal = Error("unable to convert the decimal")
	ErrTooLarge       = Error("quantity over 10^12 is too large")
	// maxDecimal value is a thousand billion, using the shortscale -- 10^12
	maxDecimal = 1e12
)

// Decimal represents a floating-point number with a fdixed precision
// example: 1.52 = 152 * 10^(-2) will be sotred as {152, 2}
type Decimal struct {
	// subunits is the amount of subunits. Multiply it by the precesion to get the real value
	subunits int64
	// Number if "subunits" in a unit, expressed as a power of 10
	precision byte
}

func ParseDecimal(value string) (Decimal, error) {
	intPart, fracPart, _ := strings.Cut(value, ".")

	subunits, err := strconv.ParseInt(intPart+fracPart, 10, 64)
	if err != nil {
		return Decimal{}, fmt.Errorf("%w: %s", ErrInvalidDecimal, err.Error())
	}

	if subunits > maxDecimal {
		return Decimal{}, ErrTooLarge
	}

	precision := byte(len(fracPart))

	dec := Decimal{subunits: subunits, precision: precision}

	dec.simplify()

	return dec, nil
}

func (d *Decimal) simplify() {
	// Using %10 returns the last digit in base 10 of a number.
	// If the precion is positive, that digit belongs to the right side of the decimal separator.
	for d.subunits%10 == 0 && d.precision > 0 {
		d.precision--
		d.subunits /= 10
	}
}

func pow10(power byte) int64 {
	switch power {
	case 0:
		return 1
	case 1:
		return 10
	case 2:
		return 100
	case 3:
		return 1000
	default:
		return int64(math.Pow(10, float64(power)))
	}
}
