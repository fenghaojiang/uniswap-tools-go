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
	"golang.org/x/sync/errgroup"
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

func (c *Clients) Slot0(ctx context.Context, address common.Address) (*model.Slot0, error) {
	_calldata, err := c.contractAbis.Pool.Pack(constants.Slot0Method)
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
	var slot0 model.Slot0
	err = c.contractAbis.Pool.UnpackIntoInterface(&slot0, constants.Slot0Method, data)
	if err != nil {
		return nil, err
	}
	return lo.ToPtr[model.Slot0](slot0), nil
}

func (c *Clients) Pool(ctx context.Context, poolAddress common.Address, tickLower, tickUpper *big.Int) (*model.PoolAggregated, error) {
	calls := make([]multicall3.Multicall3Call3, 0)

	_feeGlobal0X128Call, err := c.contractAbis.Pool.Pack(constants.FeeGrowthGlobal0X128Method)
	if err != nil {
		return nil, err
	}
	calls = append(calls, multicall3.Multicall3Call3{
		Target:       poolAddress,
		AllowFailure: false,
		CallData:     _feeGlobal0X128Call,
	})

	_feeGlobal1X128Call, err := c.contractAbis.Pool.Pack(constants.FeeGrowthGlobal1X128Method)
	if err != nil {
		return nil, err
	}
	calls = append(calls, multicall3.Multicall3Call3{
		Target:       poolAddress,
		AllowFailure: false,
		CallData:     _feeGlobal1X128Call,
	})

	_slotCall, err := c.contractAbis.Pool.Pack(constants.Slot0Method)
	if err != nil {
		return nil, err
	}
	calls = append(calls, multicall3.Multicall3Call3{
		Target:       poolAddress,
		AllowFailure: false,
		CallData:     _slotCall,
	})

	_tickLowerCall, err := c.contractAbis.Pool.Pack(constants.TicksMethod, tickLower)
	if err != nil {
		return nil, err
	}
	calls = append(calls, multicall3.Multicall3Call3{
		Target:       poolAddress,
		AllowFailure: false,
		CallData:     _tickLowerCall,
	})

	_tickUpperCall, err := c.contractAbis.Pool.Pack(constants.TicksMethod, tickUpper)
	if err != nil {
		return nil, err
	}
	calls = append(calls, multicall3.Multicall3Call3{
		Target:       poolAddress,
		AllowFailure: false,
		CallData:     _tickUpperCall,
	})

	_liquidityCall, err := c.contractAbis.Pool.Pack(constants.LiquidityMethod)
	if err != nil {
		return nil, err
	}
	calls = append(calls, multicall3.Multicall3Call3{
		Target:       poolAddress,
		AllowFailure: false,
		CallData:     _liquidityCall,
	})

	results, err := c.AggregatedCalls(ctx, calls)
	if err != nil {
		return nil, err
	}

	if len(results) != 6 {
		return nil, fmt.Errorf("call pool contract result do not match the request")
	}

	for i, result := range results {
		if !result.Success {
			return nil, fmt.Errorf("failed to call pool contract on %d th call", i)
		}
	}

	var feeGrowthGlobal0X128, feeGrowthGlobal1X128 *big.Int
	var _slot0 model.Slot0
	var _tickLower, _tickUpper model.Tick
	var _liquidity *big.Int

	var eg errgroup.Group

	eg.Go(func() error {
		return c.contractAbis.Pool.UnpackIntoInterface(&feeGrowthGlobal0X128, constants.FeeGrowthGlobal0X128Method, results[0].ReturnData)
	})
	eg.Go(func() error {
		return c.contractAbis.Pool.UnpackIntoInterface(&feeGrowthGlobal1X128, constants.FeeGrowthGlobal1X128Method, results[1].ReturnData)
	})
	eg.Go(func() error {
		return c.contractAbis.Pool.UnpackIntoInterface(&_slot0, constants.Slot0Method, results[2].ReturnData)
	})
	eg.Go(func() error {
		return c.contractAbis.Pool.UnpackIntoInterface(&_tickLower, constants.TicksMethod, results[3].ReturnData)
	})
	eg.Go(func() error {
		return c.contractAbis.Pool.UnpackIntoInterface(&_tickUpper, constants.TicksMethod, results[4].ReturnData)
	})
	eg.Go(func() error {
		return c.contractAbis.Pool.UnpackIntoInterface(&_liquidity, constants.LiquidityMethod, results[5].ReturnData)
	})

	err = eg.Wait()
	if err != nil {
		return nil, err
	}

	return &model.PoolAggregated{
		FeeGrowthGlobal0X128: feeGrowthGlobal0X128,
		FeeGrowthGlobal1X128: feeGrowthGlobal1X128,
		Slot0:                lo.ToPtr[model.Slot0](_slot0),
		TickLowerTicks:       lo.ToPtr[model.Tick](_tickLower),
		TickUpperTicks:       lo.ToPtr[model.Tick](_tickUpper),
		Liquidity:            _liquidity,
	}, nil
}
