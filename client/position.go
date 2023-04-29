package client

import (
	"context"
	"math/big"

	"github.com/fenghaojiang/uniswap-tools-go/constants"
	"github.com/fenghaojiang/uniswap-tools-go/model"
	"github.com/fenghaojiang/uniswap-tools-go/onchain/generated-go/multicall3"
	"golang.org/x/sync/errgroup"
)

const (
	NFTPositionManagerPositionMethod = "position"
)

func (c *Clients) Position(ctx context.Context, tokenIDs []*big.Int) ([]model.Position, error) {
	calls := make([]multicall3.Multicall3Call3, 0)
	for _, tokenID := range tokenIDs {
		call := multicall3.Multicall3Call3{
			Target:       constants.UniswapV3NFTPositionManagerAddress(),
			AllowFailure: false,
		}
		callData, err := c.contractAbis.NftPositionManager.Pack(NFTPositionManagerPositionMethod, tokenID)
		if err != nil {
			return nil, err
		}
		call.CallData = callData
		calls = append(calls, call)
	}

	results, err := c.aggregatedCalls(ctx, calls)
	if err != nil {
		return nil, err
	}

	var positions []model.NFTPosition

	for _, result := range results {
		var position model.NFTPosition
		if err := c.contractAbis.NftPositionManager.UnpackIntoInterface(&position, NFTPositionManagerPositionMethod, result.ReturnData); err != nil {
			return nil, err
		}
		positions = append(positions, position)
	}

	var eg errgroup.Group
	for _, position := range positions {
		c.limitChan <- struct{}{}
		_p := position
		eg.Go(func() error {
			defer func() {
				<-c.limitChan
			}()
			_, err := c.ERC20Token(ctx, _p.Token0)
			if err != nil {
				return err
			}

			_, err = c.ERC20Token(ctx, _p.Token1)
			if err != nil {
				return err
			}
			return nil
		})

	}

	return nil, nil
}
