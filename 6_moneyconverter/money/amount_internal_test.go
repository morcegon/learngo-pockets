package money

import (
	"testing"
)

func TestNewAmount(t *testing.T) {
	tt := map[string]struct {
		quantity Decimal
		currency Currency
		want     Amount
		err      error
	}{
		"nominal": {
			quantity: Decimal{subunits: 150, precision: 2},
			currency: Currency{code: "BRL", precision: 2},
			want: Amount{
				quantity: Decimal{subunits: 150, precision: 2},
				currency: Currency{code: "BRL", precision: 2},
			},
		},
		"error too precise": {
			quantity: Decimal{subunits: 150, precision: 10},
			currency: Currency{code: "BRL", precision: 2},
			want:     Amount{},
			err:      ErrTooPrecise,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := NewAmount(tc.quantity, tc.currency)
			if tc.err != err {
				t.Errorf("expecting %s, got %s", tc.err.Error(), err.Error())
			}

			if got != tc.want {
				t.Errorf("expecting %v, got %v", tc.want, got)
			}
		})
	}
}
