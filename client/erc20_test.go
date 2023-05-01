package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ethereum/go-ethereum/common"
)

func TestOnERC20Token(t *testing.T) {

	clis, err := NewClientsWithEndpoints([]string{
		"https://rpc.ankr.com/eth",
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	usdtAddress := common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7")

	ctx := context.Background()
	t.Run("run on aggregate erc20 token", func(t *testing.T) {
		info, err := clis.AggregatedERC20Token(ctx, usdtAddress)
		if err != nil {
			t.Fatal(err.Error())
		}

		assert.Equal(t, usdtAddress, info.ContractAddress)
		assert.Equal(t, "USDT", info.Symbol)
		assert.Equal(t, uint8(6), info.Decimals)
	})

	t.Run("run on erc20 symbol", func(t *testing.T) {
		symbol, err := clis.ERC20Symbol(ctx, usdtAddress)
		if err != nil {
			t.Fatal(err.Error())
		}

		assert.Equal(t, "USDT", symbol)
	})

	t.Run("run on erc20 total supply", func(t *testing.T) {
		totalSupply, err := clis.ERC20TotalSupply(ctx, usdtAddress)
		if err != nil {
			t.Fatal(err.Error())
		}

		fmt.Println(totalSupply)
	})
}
