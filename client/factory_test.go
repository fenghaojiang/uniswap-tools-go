package client

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fenghaojiang/uniswap-tools-go/model"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestOnGetUniswapV3Pool(t *testing.T) {
	clis, err := NewClientsWithEndpoints([]string{
		"https://rpc.ankr.com/eth",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Run("test on get uniswap v3 pool", func(t *testing.T) {
		res, err := clis.UniswapV3GetPool(context.Background(),
			common.HexToAddress("0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984"),
			common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
			new(big.Int).SetInt64(3000))
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, common.HexToAddress("0x1d42064Fc4Beb5F8aAF85F4617AE8b3b5B8Bd801"), lo.FromPtr[common.Address](res))
	})

	t.Run("test on aggregated getPools", func(t *testing.T) {
		res, err := clis.AggregatedGetPools(context.Background(), []model.Pool{
			{
				Token0: common.HexToAddress("0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984"),
				Token1: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
				Fee:    new(big.Int).SetInt64(3000),
			},
		})
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, 1, len(res))
		assert.Equal(t, common.HexToAddress("0x1d42064Fc4Beb5F8aAF85F4617AE8b3b5B8Bd801"), res[0].Pool)
	})

}
