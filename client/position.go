package client

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fenghaojiang/uniswap-tools-go/constants"
	"github.com/fenghaojiang/uniswap-tools-go/model"
	"github.com/fenghaojiang/uniswap-tools-go/onchain/generated-go/multicall3"
	"github.com/fenghaojiang/uniswap-tools-go/utils"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"golang.org/x/sync/errgroup"
)

func (c *Clients) AggregatedPosition(ctx context.Context, tokenIDs []*big.Int) ([]model.Position, error) {
	calls := make([]multicall3.Multicall3Call3, 0)
	for _, tokenID := range tokenIDs {
		call := multicall3.Multicall3Call3{
			Target:       constants.UniswapV3NFTPositionManagerAddress(),
			AllowFailure: false,
		}
		callData, err := c.contractAbis.NftPositionManager.Pack(constants.NFTPositionManagerPositionsMethod, tokenID)
		if err != nil {
			return nil, fmt.Errorf("failed to pack nft position manager, err: %w", err)
		}
		call.CallData = callData
		calls = append(calls, call)
	}

	results, err := c.AggregatedCalls(ctx, calls)
	if err != nil {
		return nil, fmt.Errorf("failed to aggregated call, err: %w", err)
	}

	var positions []model.NFTPosition
	for i, result := range results {
		if !result.Success {
			return nil, fmt.Errorf("failed to call uniswap nft manager position at %d th call", i)
		}
		var position model.NFTPosition
		if err := c.contractAbis.NftPositionManager.UnpackIntoInterface(&position, constants.NFTPositionManagerPositionsMethod, result.ReturnData); err != nil {
			return nil, fmt.Errorf("failed to unpack nft position, err:%w", err)
		}
		positions = append(positions, position)
	}

	var eg errgroup.Group

	tokenMap := make(map[string]*model.ERC20Token)

	var mu sync.Mutex
	pools := make([]model.Pool, 0)

	for _, position := range positions {
		if c.limitChan != nil {
			c.limitChan <- struct{}{}
		}
		_p := position
		pools = append(pools, model.Pool{
			Token0: _p.Token0,
			Token1: _p.Token1,
			Fee:    _p.Fee,
		})

		var token0, token1 *model.ERC20Token
		var err error
		eg.Go(func() error {
			defer func() {
				if c.limitChan != nil {
					<-c.limitChan
				}
			}()
			token0, err = c.AggregatedERC20Token(ctx, _p.Token0)
			if err != nil {
				return fmt.Errorf("failed to call aggregated erc20 token, err: %w", err)
			}
			mu.Lock()
			tokenMap[strings.ToLower(_p.Token0.String())] = token0
			mu.Unlock()

			token1, err = c.AggregatedERC20Token(ctx, _p.Token1)
			if err != nil {
				return fmt.Errorf("failed to call aggregated erc20 token, err: %w", err)
			}

			mu.Lock()
			tokenMap[strings.ToLower(_p.Token1.String())] = token1
			mu.Unlock()

			return nil
		})
	}

	_tokenAddresses := make([]common.Address, 0)
	tokenPrice := make(map[string]*decimal.Decimal)
	for _, pool := range pools {
		_tokenAddresses = append(_tokenAddresses, pool.Token0)
		_tokenAddresses = append(_tokenAddresses, pool.Token1)
	}

	if c.limitChan != nil {
		c.limitChan <- struct{}{}
	}
	eg.Go(func() error {
		defer func() {
			if c.limitChan != nil {
				<-c.limitChan
			}
		}()

		pools, err = c.AggregatedGetPools(ctx, pools)
		if err != nil {
			return fmt.Errorf("failed to aggregated get pools, err:%w", err)
		}

		return nil
	})

	if c.limitChan != nil {
		c.limitChan <- struct{}{}
	}

	eg.Go(func() error {
		defer func() {
			if c.limitChan != nil {
				<-c.limitChan
			}
		}()
		prices, err := c.AggregatedTokenPriceInUSD(ctx, _tokenAddresses)
		if err != nil {
			return fmt.Errorf("failed to aggregated get token price, err: %w", err)
		}
		for i, p := range prices {
			tokenPrice[strings.ToLower(_tokenAddresses[i].String())] = p
		}
		return nil
	})

	err = eg.Wait()
	if err != nil {
		return nil, err
	}

	resultPositions := make([]model.Position, 0)

	for i, position := range positions {
		_i := i
		_position := position
		if c.limitChan != nil {
			c.limitChan <- struct{}{}
		}
		eg.Go(func() error {
			defer func() {
				if c.limitChan != nil {
					<-c.limitChan
				}
			}()
			poolInfo, err := c.Pool(ctx, pools[_i].Pool, _position.TickLower, _position.TickUpper)
			if err != nil {
				return fmt.Errorf("failed to get pool, err: %w", err)
			}

			token0Addr := strings.ToLower(_position.Token0.String())
			token1Addr := strings.ToLower(_position.Token1.String())

			_tPriceLower := utils.TickToPrice(_position.TickLower)
			_tPriceUpper := utils.TickToPrice(_position.TickUpper)
			priceRangeInToken0_0 := utils.AdjustedPrice(_tPriceLower,
				tokenMap[token0Addr].Decimals, tokenMap[token1Addr].Decimals)
			priceRangeInToken0_1 := utils.AdjustedPrice(_tPriceUpper,
				tokenMap[token0Addr].Decimals, tokenMap[token1Addr].Decimals)

			priceRangeInToken1_0 := utils.Invert(priceRangeInToken0_0)
			priceRangeInToken1_1 := utils.Invert(priceRangeInToken0_1)

			currentTick := decimal.NewFromBigInt(poolInfo.Slot0.Tick, 0)
			tickUpper := decimal.NewFromBigInt(_position.TickUpper, 0)
			tickLower := decimal.NewFromBigInt(_position.TickLower, 0)

			var status constants.Status = constants.StatusInRange
			if currentTick.Cmp(tickUpper) > 0 {
				status = constants.StatusOutOfUpRange
			}
			if currentTick.Cmp(tickLower) < 0 {
				status = constants.StatusOutOfLowRange
			}
			if poolInfo.Liquidity.Cmp(new(big.Int).SetInt64(0)) <= 0 {
				status = constants.StatusClose
			}
			var lockToken0Amount decimal.Decimal
			var lockToken1Amount decimal.Decimal

			var feeReward0Amount decimal.Decimal
			var feeReward1Amount decimal.Decimal

			switch status {
			case constants.StatusClose:
				mu.Lock()
				resultPositions = append(resultPositions, model.Position{
					TokenID: tokenIDs[_i],
					Name: fmt.Sprintf("%s - %s, fee: %.2f",
						tokenMap[token0Addr].Symbol, tokenMap[token1Addr].Symbol, float64(pools[_i].Fee.Int64())/1000),
					Status: constants.StatusClose,
				})
				mu.Unlock()
				return nil

			case constants.StatusOutOfLowRange:
				lockToken1Amount = decimal.NewFromInt(0)
				lockToken0Amount = utils.TickPriceToToken0Balance(tokenMap[token0Addr].Decimals,
					_tPriceLower, _tPriceUpper, _position.Liquidity)

			case constants.StatusOutOfUpRange:
				lockToken0Amount = decimal.NewFromInt(0)
				lockToken1Amount = utils.TickPriceToToken1Balance(tokenMap[token1Addr].Decimals,
					_tPriceLower, _tPriceUpper, _position.Liquidity)

			}

			totalRewardValue := decimal.NewFromInt(1).Mul(feeReward0Amount).Mul(*tokenPrice[token0Addr]).Add(
				decimal.NewFromInt(1).Mul(feeReward1Amount).Mul(*tokenPrice[token1Addr]))

			totalValueInUSD := decimal.NewFromInt(1).Mul(lockToken0Amount).Mul(*tokenPrice[token0Addr]).
				Add(decimal.NewFromInt(1).Mul(lockToken1Amount).Mul(*tokenPrice[token1Addr])).
				Add(totalRewardValue)

			mu.Lock()
			resultPositions = append(resultPositions, model.Position{
				Status:  status,
				TokenID: tokenIDs[_i],
				Name: fmt.Sprintf("%s - %s, fee: %.2f",
					tokenMap[token0Addr].Symbol, tokenMap[token1Addr].Symbol, float64(pools[_i].Fee.Int64())/1000),
				PriceRangeInToken0: [2]*decimal.Decimal{
					lo.ToPtr[decimal.Decimal](priceRangeInToken0_0),
					lo.ToPtr[decimal.Decimal](priceRangeInToken0_1),
				},
				PriceRangeInToken1: [2]*decimal.Decimal{
					lo.ToPtr[decimal.Decimal](priceRangeInToken1_0),
					lo.ToPtr[decimal.Decimal](priceRangeInToken1_1),
				},

				Token0: common.HexToAddress(token0Addr),
				Token1: common.HexToAddress(token1Addr),

				Symbol0: tokenMap[token0Addr].Symbol,
				Symbol1: tokenMap[token1Addr].Symbol,

				LockedAmount0: lo.ToPtr[decimal.Decimal](lockToken0Amount),
				LockedAmount1: lo.ToPtr[decimal.Decimal](lockToken1Amount),

				LockedValue0InUSD: lo.ToPtr[decimal.Decimal](decimal.NewFromInt(1).Mul(lockToken0Amount).Mul(*tokenPrice[token0Addr])),
				LockedValue1InUSD: lo.ToPtr[decimal.Decimal](decimal.NewFromInt(1).Mul(lockToken1Amount).Mul(*tokenPrice[token1Addr])),

				TotalRewardsUSD: lo.ToPtr[decimal.Decimal](totalRewardValue),
				TotalValueUSD:   lo.ToPtr[decimal.Decimal](totalValueInUSD),
			})
			mu.Unlock()

			return nil
		})

	}

	err = eg.Wait()
	if err != nil {
		return nil, err
	}

	// TODO reward calculation

	return resultPositions, nil
}
