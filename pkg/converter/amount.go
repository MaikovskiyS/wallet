package converter

import (
	"math"

	"github.com/shopspring/decimal"
)

var decimals = math.Pow10(18)

type converter struct {
	decimals uint64
}

func New() *converter {
	return &converter{
		decimals: uint64(decimals),
	}
}
func (c *converter) EvmDecimal(amount float64) (uint64, error) {

	dec := decimal.NewFromFloat(amount)
	rat := dec.Rat()
	denom := rat.Denom()

	amount1 := rat.Num().Uint64() * (c.decimals / denom.Uint64())
	return amount1, nil
}

func (c *converter) ToFloat(amount uint64) float64 {
	fAmount := float64(amount) / float64(c.decimals)

	return fAmount
}
