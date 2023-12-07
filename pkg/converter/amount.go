package converter

import (
	"math"

	"github.com/shopspring/decimal"
)

var decimals = math.Pow10(18)

type converter struct {
	decimals int64
}

func New() *converter {
	return &converter{
		decimals: int64(decimals),
	}
}
func (c *converter) EvmDecimal(amount float64) (int64, error) {

	dec := decimal.NewFromFloat(amount)
	rat := dec.Rat()
	denom := rat.Denom()

	amount1 := rat.Num().Int64() * (c.decimals / denom.Int64())
	return amount1, nil
}
