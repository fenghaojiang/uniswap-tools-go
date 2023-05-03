package client

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/fenghaojiang/uniswap-tools-go/constants"
	"github.com/fenghaojiang/uniswap-tools-go/model"
	"github.com/fenghaojiang/uniswap-tools-go/onchain/generated-go/multicall3"
)

func (c *Clients) FeeGrowthGlobal0X128(ctx context.Context, address common.Address) (*big.Int, error) {
	_calldata, err := c.contractAbis.Pool.Pack(constants.FeeGrowthGlobal0X128Method)
	if err != nil {
		return nil, err
	}

	data, err := c.Call(ctx, model.CallContractParam{
		To:       address.String(),
		CallData: hexutil.Encode(_calldata),
	})
	if err != nil {
		return nil, err
	}

	var res *big.Int
	err = c.contractAbis.Pool.UnpackIntoInterface(&res, constants.FeeGrowthGlobal0X128Method, data)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Clients) FeeGrowthGlobal1X128(ctx context.Context, address common.Address) (*big.Int, error) {
	_calldata, err := c.contractAbis.Pool.Pack(constants.FeeGrowthGlobal1X128Method)
	if err != nil {
		return nil, err
	}

	data, err := c.Call(ctx, model.CallContractParam{
		To:       address.String(),
		CallData: hexutil.Encode(_calldata),
	})
	if err != nil {
		return nil, err
	}

	var res *big.Int
	err = c.contractAbis.Pool.UnpackIntoInterface(&res, constants.FeeGrowthGlobal1X128Method, data)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Clients) AggregatedFeeGrowthGlobal0X128(ctx context.Context, addresses []common.Address) ([]*big.Int, error) {
	_calls := make([]multicall3.Multicall3Call3, 0)
	for _, address := range addresses {
		_calldata, err := c.contractAbis.Pool.Pack(constants.FeeGrowthGlobal0X128Method)
		if err != nil {
			return nil, err
		}

		_calls = append(_calls, multicall3.Multicall3Call3{
			Target:       address,
			AllowFailure: false,
			CallData:     _calldata,
		})
	}

	results, err := c.AggregatedCalls(ctx, _calls)
	if err != nil {
		return nil, err
	}

	fees := make([]*big.Int, 0)
	for i, result := range results {
		if !result.Success {
			return nil, fmt.Errorf("feeGrowthGlobal0X128 failed on %d th call, contract address: %s", i, addresses[i].String())
		}

		var fee = new(big.Int)
		err = c.contractAbis.Pool.UnpackIntoInterface(&fee, constants.FeeGrowthGlobal0X128Method, result.ReturnData)
		if err != nil {
			return nil, err
		}
		fees = append(fees, fee)
	}

	return fees, nil
}

func (c *Clients) AggregatedFeeGrowthGlobal1X128(ctx context.Context, addresses []common.Address) ([]*big.Int, error) {
	_calls := make([]multicall3.Multicall3Call3, 0)
	for _, address := range addresses {
		_calldata, err := c.contractAbis.Pool.Pack(constants.FeeGrowthGlobal1X128Method)
		if err != nil {
			return nil, err
		}

		_calls = append(_calls, multicall3.Multicall3Call3{
			Target:       address,
			AllowFailure: false,
			CallData:     _calldata,
		})
	}

	results, err := c.AggregatedCalls(ctx, _calls)
	if err != nil {
		return nil, err
	}

	fees := make([]*big.Int, 0)
	for i, result := range results {
		if !result.Success {
			return nil, fmt.Errorf("feeGrowthGlobal0X128 failed on %d th call, contract address: %s", i, addresses[i].String())
		}

		var fee = new(big.Int)
		err = c.contractAbis.Pool.UnpackIntoInterface(&fee, constants.FeeGrowthGlobal1X128Method, result.ReturnData)
		if err != nil {
			return nil, err
		}
		fees = append(fees, fee)
	}

	return fees, nil
}
