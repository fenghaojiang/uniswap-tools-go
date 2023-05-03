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
	"github.com/samber/lo"
)

func (c *Clients) UniswapV3GetPool(ctx context.Context, token0 common.Address, token1 common.Address, feeLevel *big.Int) (*common.Address, error) {
	_calldata, err := c.contractAbis.Factory.Pack(constants.GetPoolMethod, token0, token1, feeLevel)
	if err != nil {
		return nil, err
	}

	data, err := c.Call(ctx, model.CallContractParam{
		To:       constants.UniswapV3FacotryAddress().String(),
		CallData: hexutil.Encode(_calldata),
	})
	if err != nil {
		return nil, err
	}

	var pool common.Address
	err = c.contractAbis.Factory.UnpackIntoInterface(&pool, constants.GetPoolMethod, data)
	if err != nil {
		return nil, err
	}

	return lo.ToPtr[common.Address](pool), nil
}

func (c *Clients) AggregatedGetPools(ctx context.Context, getPoolsReq []model.Pool) ([]model.Pool, error) {
	_calls := make([]multicall3.Multicall3Call3, 0)
	for _, pool := range getPoolsReq {
		_calldata, err := c.contractAbis.Factory.Pack(constants.GetPoolMethod, pool.Token0, pool.Token1, pool.Fee)
		if err != nil {
			return nil, err
		}
		_calls = append(_calls, multicall3.Multicall3Call3{
			Target:       constants.UniswapV3FacotryAddress(),
			AllowFailure: false,
			CallData:     _calldata,
		})
	}

	results, err := c.AggregatedCalls(ctx, _calls)
	if err != nil {
		return nil, err
	}
	if len(getPoolsReq) != len(results) {
		return nil, fmt.Errorf("response of call getPool do not match the request")
	}
	for i, res := range results {
		if !res.Success {
			return nil, fmt.Errorf("failed to handle %d th request, %+v", i, getPoolsReq[i])
		}
		var poolAddress common.Address
		err = c.contractAbis.Factory.UnpackIntoInterface(&poolAddress, constants.GetPoolMethod, res.ReturnData)
		if err != nil {
			return nil, err
		}
		getPoolsReq[i].Pool = poolAddress
	}
	return getPoolsReq, nil
}
