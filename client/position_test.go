package client

import (
	"context"
	"math/big"
	"testing"
)

func TestOnAccountHoldings(t *testing.T) {
	clis, err := NewClientsWithEndpoints([]string{
		"https://rpc.ankr.com/eth",
	})
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	_, err = clis.AggregatedPosition(ctx, []*big.Int{
		new(big.Int).SetInt64(100000),
	})
	if err != nil {
		t.Fatal(err)
	}

}
