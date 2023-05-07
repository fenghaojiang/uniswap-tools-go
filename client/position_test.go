package client

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/fenghaojiang/uniswap-tools-go/constants"
)

func TestOnAccountHoldings(t *testing.T) {
	clis, err := NewClientsWithEndpoints([]string{
		"https://rpc.ankr.com/polygon",
	})
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	position, err := clis.WithNetwork(constants.PolygonNetwork).AggregatedPosition(ctx, []*big.Int{
		new(big.Int).SetInt64(869899),
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v", position)
}
