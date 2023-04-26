package client

import (
	"context"
	"testing"
)

func TestOnAccountHoldings(t *testing.T) {
	cli, err := NewClientsWithEndpoints([]string{
		"wss://eth-mainnet.g.alchemy.com/v2/RZpB2x6G6Ls0m3lgRTh66fLwHgeseYSj",
	})
	if err != nil {
		t.Fatal(err)
	}
	holdings, err := cli.GetAccountHoldings(context.Background(), "0xc101c69340FEB4d0c474BF8fC34f5266F3de8A15")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(holdings)
}
