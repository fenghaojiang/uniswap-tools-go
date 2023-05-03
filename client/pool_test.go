package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestOnFeeGrowthGlobal(t *testing.T) {
	clis, err := NewClientsWithEndpoints([]string{
		"https://rpc.ankr.com/eth",
	})
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	t.Run("test on feeGrowthGlobal0X128", func(t *testing.T) {
		res, err := clis.FeeGrowthGlobal0X128(ctx, common.HexToAddress("0x1d42064Fc4Beb5F8aAF85F4617AE8b3b5B8Bd801"))
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(res)
	})

	t.Run("test on feeGrowthGlobal1X128", func(t *testing.T) {
		res, err := clis.FeeGrowthGlobal1X128(ctx, common.HexToAddress("0x1d42064Fc4Beb5F8aAF85F4617AE8b3b5B8Bd801"))
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(res)
	})

	t.Run("test on aggregatedFeeGrowthGlobal0X128", func(t *testing.T) {
		res, err := clis.AggregatedFeeGrowthGlobal0X128(ctx, []common.Address{
			common.HexToAddress("0x1d42064Fc4Beb5F8aAF85F4617AE8b3b5B8Bd801"),
		})
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(res)
	})

	t.Run("test on aggregatedFeeGrowthGlobal1X128", func(t *testing.T) {
		res, err := clis.AggregatedFeeGrowthGlobal1X128(ctx, []common.Address{
			common.HexToAddress("0x1d42064Fc4Beb5F8aAF85F4617AE8b3b5B8Bd801"),
		})
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(res)
	})
}
