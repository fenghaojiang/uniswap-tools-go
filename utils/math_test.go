package utils

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/shopspring/decimal"
)

func TestOnMath(t *testing.T) {
	t.Run("test on tick to price", func(t *testing.T) {
		fmt.Println(TickToPrice(new(big.Int).SetInt64(200240)))
		fmt.Println(TickToPrice(new(big.Int).SetInt64(200700)))
	})

	t.Run("test on adjusted price", func(t *testing.T) {
		fmt.Println(AdjustedPrice(TickToPrice(new(big.Int).SetInt64(200240)), 6, 18))
		fmt.Println(AdjustedPrice(TickToPrice(new(big.Int).SetInt64(200700)), 6, 18))
	})

	t.Run("test on invert", func(t *testing.T) {
		fmt.Println(Invert(decimal.NewFromFloat(0.00049645274801)))
		fmt.Println(Invert(decimal.NewFromFloat(0.00051982177317)))
	})

}
