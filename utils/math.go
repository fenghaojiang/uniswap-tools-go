package utils

import (
	"math"
	"math/big"

	"github.com/shopspring/decimal"
)

const (
	base = 1.0001
)

func TickToPrice(tick *big.Int) decimal.Decimal {
	_tick := decimal.NewFromBigInt(tick, 0)
	_base := decimal.NewFromFloat(base)
	return _base.Pow(_tick)
}

func AdjustedPrice(price decimal.Decimal, token0Decimals, token1Decimals uint8) decimal.Decimal {
	return price.Mul(decimal.New(1, int32(token0Decimals)-int32(token1Decimals)))
}

func Invert(number decimal.Decimal) decimal.Decimal {
	if number.IsZero() {
		return decimal.NewFromInt(0)
	}
	return decimal.NewFromInt(1).Div(number)
}

func TickPriceToToken0Balance(decimals uint8, tickPriceLower, tickPriceUpper decimal.Decimal, liquidity *big.Int) decimal.Decimal {
	_tickLower := tickPriceLower.BigFloat()
	_tickUpper := tickPriceUpper.BigFloat()
	sqrtLower := _tickLower.Sqrt(_tickLower)
	sqrtUpper := _tickUpper.Sqrt(_tickUpper)

	bias := SafeSub(sqrtUpper, sqrtLower)
	accum := new(big.Float).Mul(sqrtLower, sqrtUpper)

	_liquidity := new(big.Float).SetInt(liquidity)

	_balance := new(big.Float).Mul(_liquidity, new(big.Float).Quo(bias, accum))
	_balanceF64, _ := _balance.Float64()
	balance := decimal.NewFromFloat(_balanceF64).Shift(-int32(decimals))
	return balance
}

func TickPriceToToken1Balance(decimals uint8, tickPriceLower, tickPriceUpper decimal.Decimal, liquidity *big.Int) decimal.Decimal {
	_tickLower := tickPriceLower.BigFloat()
	_tickUpper := tickPriceUpper.BigFloat()
	sqrtLower := _tickLower.Sqrt(_tickLower)
	sqrtUpper := _tickUpper.Sqrt(_tickUpper)

	_liquidity := new(big.Float).SetInt(liquidity)
	bias := SafeSub(sqrtUpper, sqrtLower)

	_balance := new(big.Float).Mul(_liquidity, bias)
	_balanceF64, _ := _balance.Float64()
	balance := decimal.NewFromFloat(_balanceF64).Shift(-int32(decimals))
	return balance
}

func SafeSub(x *big.Float, y *big.Float) *big.Float {
	Q256 := big.NewFloat(math.Pow(2, 256))
	diff := new(big.Float).Sub(x, y)
	if x.Cmp(y) >= 0 {
		return diff
	} else {
		return diff.Add(diff, Q256)
	}
}
