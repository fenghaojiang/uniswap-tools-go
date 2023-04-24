package client

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/fenghaojiang/uniswap-tools-go/constants"
	"github.com/fenghaojiang/uniswap-tools-go/onchain/generated-go/multicall3"
	nftmanager "github.com/fenghaojiang/uniswap-tools-go/onchain/generated-go/uniswapv3_nft_position_manager"
)

func (c *Clients) GetAccountHoldings(ctx context.Context, accountAddress string) ([]*big.Int, error) {
	cli := c.Client()
	if cli == nil {
		return nil, fmt.Errorf("no client available")
	}

	logs, err := cli.FilterLogs(ctx, ethereum.FilterQuery{
		Addresses: []common.Address{
			constants.UniswapV3NFTPositionManagerAddress(),
		},
		Topics: [][]common.Hash{
			{
				constants.ERC721TransferHash(),
			},
			{},
			{
				common.HexToHash(accountAddress),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return c.FilterAccountHoldings(ctx, logs, accountAddress)
}

func (c *Clients) FilterAccountHoldings(ctx context.Context, logs []types.Log, accountAddress string) ([]*big.Int, error) {
	calls := make([]multicall3.Multicall3Call3, 0)
	tokenIDs := make([]*big.Int, 0)

	for i := range logs {
		var transferEvent nftmanager.Uniswapv3NftPositionManagerTransfer
		err := c.contractAbis.NftPositionManager.UnpackIntoInterface(&transferEvent, constants.TransferEvent, logs[i].Data)
		if err != nil {
			return nil, err
		}

		_calldata, err := c.contractAbis.NftPositionManager.Pack(constants.OwnerOfMethod, transferEvent.TokenId)
		if err != nil {
			return nil, err
		}
		calls = append(calls, multicall3.Multicall3Call3{
			Target:       constants.UniswapV3NFTPositionManagerAddress(),
			AllowFailure: false,
			CallData:     _calldata,
		})

		tokenIDs = append(tokenIDs, transferEvent.TokenId)
	}

	callResults, err := c.AggregatedCalls(ctx, calls)
	if err != nil {
		return nil, err
	}

	filterTokens := make([]*big.Int, 0)
	for i := range callResults {
		var ownerOf common.Address
		err = c.contractAbis.Multicall.UnpackIntoInterface(&ownerOf, constants.OwnerOfMethod, callResults[i].ReturnData)
		if err != nil {
			return nil, err
		}
		if !callResults[i].Success {
			return nil, fmt.Errorf("call failed")
		}
		if strings.EqualFold(ownerOf.String(), accountAddress) {
			filterTokens = append(filterTokens, tokenIDs[i])
		}
	}

	return filterTokens, nil
}

// TODO: implement this
func (c *Clients) Position() {

}
