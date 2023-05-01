package client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ethereum/go-ethereum/common"
)

func TestOnERC20Token(t *testing.T) {
	t.Run("run on erc20 token", func(t *testing.T) {
		clis, err := NewClientsWithEndpoints([]string{
			"https://rpc.ankr.com/eth",
		})
		if err != nil {
			t.Fatal(err.Error())
		}
		usdtAddress := common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7")

		info, err := clis.ERC20Token(context.Background(), usdtAddress)
		if err != nil {
			t.Fatal(err.Error())
		}
		assert.Equal(t, usdtAddress, info.ContractAddress)
		assert.Equal(t, "USDT", info.Symbol)
		assert.Equal(t, uint8(6), info.Decimals)
	})
}
