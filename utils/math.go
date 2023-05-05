package utils

import (
	"math/big"

	"github.com/shopspring/decimal"
)

const (
	base = 1.0001
)

func TickToPrice(tick *big.Int) decimal.Decimal {
	_tick := decimal.NewFromBigInt(tick, 0)
	_base := decimal.NewFromFloat(base)
	return _base.Pow(_tick).Round(2)
}

func AdjustedPrice(price decimal.Decimal, token0Decimals, token1Decimals uint8) decimal.Decimal {
	return price.Mul(decimal.New(1, int32(token0Decimals)-int32(token1Decimals)))
}

func Invert(number decimal.Decimal) decimal.Decimal {
	if number.IsZero() {
		return decimal.NewFromInt(0)
	}
	return decimal.NewFromInt(1).Div(number).Round(2)
}
