package client

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/fenghaojiang/uniswap-tools-go/constants"
	"github.com/fenghaojiang/uniswap-tools-go/model"
	"github.com/fenghaojiang/uniswap-tools-go/onchain/generated-go/multicall3"
	"github.com/samber/lo"
)

func (c *Clients) AggregatedCalls(ctx context.Context, calls []multicall3.Multicall3Call3) ([]multicall3.Multicall3Result, error) {
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

func (c *Clients) Call(ctx context.Context, callParam model.CallContractParam) ([]byte, error) {
	cli := c.Client()
	if cli == nil {
		return nil, fmt.Errorf("no available client")
	}

	var _res string
	err := cli.RPCClient().CallContext(ctx, &_res, "eth_call", callParam, "latest")
	if err != nil {
		return nil, err
	}

	decodeRes, err := hexutil.Decode(_res)
	if err != nil {
		return nil, err
	}

	return decodeRes, nil
}
