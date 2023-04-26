package client

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/fenghaojiang/uniswap-tools-go/constants"
	"github.com/fenghaojiang/uniswap-tools-go/onchain/generated-go/multicall3"
	"github.com/samber/lo"
)

func (c *Clients) aggregatedCalls(ctx context.Context, calls []multicall3.Multicall3Call3) ([]multicall3.Multicall3Result, error) {
	cli := c.Client()
	if cli == nil {
		return nil, fmt.Errorf("no client available")
	}

	calldata, err := c.contractAbis.Multicall.Pack(constants.Aggregate3Method, calls)
	if err != nil {
		return nil, err
	}

	callMsg := ethereum.CallMsg{
		To:   lo.ToPtr(constants.Multicall3Address()),
		Data: calldata,
	}

	res, err := cli.ETHClient().CallContract(ctx, callMsg, nil)
	if err != nil {
		return nil, err
	}

	var results []multicall3.Multicall3Result
	if err := c.contractAbis.Multicall.UnpackIntoInterface(&results, constants.Aggregate3Method, res); err != nil {
		return nil, err
	}

	return results, nil
}
