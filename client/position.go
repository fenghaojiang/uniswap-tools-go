package client

import (
	"context"
	"math/big"
	"strings"
	"sync"

	"github.com/fenghaojiang/uniswap-tools-go/constants"
	"github.com/fenghaojiang/uniswap-tools-go/model"
	"github.com/fenghaojiang/uniswap-tools-go/onchain/generated-go/multicall3"
	"golang.org/x/sync/errgroup"
)

func (c *Clients) Position(ctx context.Context, tokenIDs []*big.Int) ([]model.Position, error) {
	calls := make([]multicall3.Multicall3Call3, 0)
	for _, tokenID := range tokenIDs {
		call := multicall3.Multicall3Call3{
			Target:       constants.UniswapV3NFTPositionManagerAddress(),
			AllowFailure: false,
		}
		callData, err := c.contractAbis.NftPositionManager.Pack(constants.NFTPositionManagerPositionsMethod, tokenID)
		if err != nil {
			return nil, err
		}
		call.CallData = callData
		calls = append(calls, call)
	}

	results, err := c.AggregatedCalls(ctx, calls)
	if err != nil {
		return nil, err
	}

	var positions []model.NFTPosition

	for _, result := range results {
		var position model.NFTPosition
		if err := c.contractAbis.NftPositionManager.UnpackIntoInterface(&position, constants.NFTPositionManagerPositionsMethod, result.ReturnData); err != nil {
			return nil, err
		}
		positions = append(positions, position)
	}

	var eg errgroup.Group

	tokenMap := make(map[string]*model.ERC20Token)
	var tokenMu sync.Mutex

	for _, position := range positions {
		if c.limitChan != nil {
			c.limitChan <- struct{}{}
		}
		_p := position
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
				return err
			}
			tokenMu.Lock()
			tokenMap[strings.ToLower(_p.Token0.String())] = token0
			tokenMu.Unlock()

			token1, err = c.AggregatedERC20Token(ctx, _p.Token1)
			if err != nil {
				return err
			}
			tokenMu.Lock()
			tokenMap[strings.ToLower(_p.Token1.String())] = token1
			tokenMu.Unlock()
			return nil
		})
	}

	err = eg.Wait()
	if err != nil {
		return nil, err
	}

	return nil, nil
}
