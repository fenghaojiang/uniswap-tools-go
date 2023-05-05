package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fenghaojiang/uniswap-tools-go/constants"
)

func TestOnClientOracle(t *testing.T) {
	clis, err := NewClientsWithEndpoints([]string{
		"https://rpc.ankr.com/eth",
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	ctx := context.Background()
	clis.WithNetwork(constants.EthereumNetwork)

	usdPrice, err := clis.TokenPriceInUSD(ctx, common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("usd price:", usdPrice)

	uniPrice, err := clis.TokenPriceInUSD(ctx, common.HexToAddress("0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("uni price:", uniPrice)

	bscClis, err := NewClientsWithEndpoints([]string{
		"https://bsc-dataseed1.binance.org/",
	})
	if err != nil {
		t.Fatal(err)
	}
	bscClis.WithNetwork(constants.BinanceSmartChainNetwork)

	pancakePrice, err := bscClis.TokenPriceInUSD(ctx, common.HexToAddress("0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("pancake price:", pancakePrice)
}
