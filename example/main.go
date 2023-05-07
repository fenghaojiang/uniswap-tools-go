package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/fenghaojiang/uniswap-tools-go/client"
	"github.com/fenghaojiang/uniswap-tools-go/constants"
)

func main() {
	polygonClis, err := client.NewClientsWithEndpoints([]string{
		"https://rpc.ankr.com/polygon",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	results, err := polygonClis.WithLimitRPC(100).WithNetwork(constants.PolygonNetwork).AggregatedPosition(context.Background(), []*big.Int{
		new(big.Int).SetInt64(869899),
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	for i := range results {
		fmt.Printf("%+v\n", results[i])
	}
}
