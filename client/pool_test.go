package client

import (
	"context"
	"fmt"
	"math/big"
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

	t.Run("test on slot0", func(t *testing.T) {
		slot0, err := clis.Slot0(ctx, common.HexToAddress("0x1d42064Fc4Beb5F8aAF85F4617AE8b3b5B8Bd801"))

		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(slot0)
	})

	polClis, err := NewClientsWithEndpoints([]string{
		"https://rpc.ankr.com/polygon",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Run("test on pool batch request", func(t *testing.T) {
		info, err := polClis.Pool(ctx,
			common.HexToAddress("0xA374094527e1673A86dE625aa59517c5dE346d32"), new(big.Int).SetInt64(-276370), new(big.Int).SetInt64(-276350))
		if err != nil {
			t.Fatal(err)
		}

		fmt.Printf("%+v\n", info.FeeGrowthGlobal0X128)
		fmt.Printf("%+v\n", info.FeeGrowthGlobal1X128)
		fmt.Printf("%+v\n", info.Slot0)
		fmt.Printf("%+v\n", info.Liquidity)
		fmt.Printf("%+v\n", info.TickLowerTicks)
		fmt.Printf("%+v\n", info.TickUpperTicks)
	})
}
