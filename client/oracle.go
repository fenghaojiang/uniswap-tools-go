package client

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fenghaojiang/uniswap-tools-go/constants"
	"github.com/fenghaojiang/uniswap-tools-go/onchain/generated-go/multicall3"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

func (c *Clients) OneInchOracleAddress() common.Address {
	switch c.network {
	case constants.EthereumNetwork:
		return constants.OneInchPriceOracleAddressEthereum()
	case constants.PolygonNetwork:
		return constants.OneInchPriceOracleAddressPolygon()
	case constants.ArbitrumNetwork:
		return constants.OneInchPriceOracleAddressArbitrum()
	case constants.BinanceSmartChainNetwork:
		return constants.OneInchPriceOracleAddressBSC()
	case constants.OptimismNetwork:
		return constants.OneInchPriceOracleAddressOptimism()
	default:
		return constants.OneInchPriceOracleAddressEthereum()
	}
}

func (c *Clients) USDAddress() common.Address {
	switch c.network {
	case constants.EthereumNetwork:
		return constants.USDAddressInEthereum()
	case constants.PolygonNetwork:
		return constants.USDAddressInPolygon()
	case constants.ArbitrumNetwork:
		return constants.USDAddressInArbitrum()
	case constants.BinanceSmartChainNetwork:
		return constants.USDAddressInBSC()
	case constants.OptimismNetwork:
		return constants.USDAddressInOptimism()
	default:
		return constants.USDAddressInEthereum()
	}
}

func (c *Clients) USDDecimals() uint8 {
	switch c.network {
	case constants.EthereumNetwork:
		return constants.USDTDecimalsInEthereum
	case constants.PolygonNetwork:
		return constants.USDTDecimalsInPolygon
	case constants.ArbitrumNetwork:
		return constants.USDTDecimalsInArbitrum
	case constants.BinanceSmartChainNetwork:
		return constants.BUSDDecimalsInBSC
	case constants.OptimismNetwork:
		return constants.USDTDecimalsInOptimism
	default:
		return constants.USDTDecimalsInEthereum
	}
}

func (c *Clients) TokenPriceInUSD(ctx context.Context, tokenAddress common.Address) (*decimal.Decimal, error) {
	_tokenCalldata, err := c.contractAbis.Oracle.Pack(constants.GetRateToETHMethod, tokenAddress, true)
	if err != nil {
		return nil, err
	}

	_usdCallData, err := c.contractAbis.Oracle.Pack(constants.GetRateToETHMethod, c.USDAddress(), true)
	if err != nil {
		return nil, err
	}

	_decimalsCallData, err := c.contractAbis.ERC20.Pack(constants.DecimalsMethod)
	if err != nil {
		return nil, err
	}

	results, err := c.AggregatedCalls(ctx, []multicall3.Multicall3Call3{
		{
			Target:       c.OneInchOracleAddress(),
			AllowFailure: false,
			CallData:     _tokenCalldata,
		},
		{
			Target:       c.OneInchOracleAddress(),
			AllowFailure: false,
			CallData:     _usdCallData,
		},
		{
			Target:       tokenAddress,
			AllowFailure: false,
			CallData:     _decimalsCallData,
		},
	})
	if err != nil {
		return nil, err
	}

	if len(results) != 3 {
		return nil, fmt.Errorf("failed to call on 1Inch oracle, result do not match")
	}

	var tokenPrice *big.Int
	var usdPrice *big.Int
	var decimals uint8

	for i, result := range results {
		if !result.Success {
			return nil, fmt.Errorf("failed to call 1Inch oracle on %d th call", i)
		}
	}

	err = c.contractAbis.Oracle.UnpackIntoInterface(&tokenPrice, constants.GetRateToETHMethod, results[0].ReturnData)
	if err != nil {
		return nil, err
	}

	err = c.contractAbis.Oracle.UnpackIntoInterface(&usdPrice, constants.GetRateToETHMethod, results[1].ReturnData)
	if err != nil {
		return nil, err
	}

	err = c.contractAbis.ERC20.UnpackIntoInterface(&decimals, constants.DecimalsMethod, results[2].ReturnData)
	if err != nil {
		return nil, err
	}

	tokenInETH := decimal.NewFromBigInt(tokenPrice, -int32(constants.EthereumDecimals))
	usdInETH := decimal.NewFromBigInt(usdPrice, -int32(constants.EthereumDecimals))

	zero := decimal.NewFromInt(0)
	if tokenInETH.Equal(zero) {
		return lo.ToPtr[decimal.Decimal](zero), nil
	}

	tokenPerETH := decimal.NewFromInt(1).Div(tokenInETH)
	usdPerETH := decimal.NewFromInt(1).Div(usdInETH)

	usdPerETH = usdPerETH.Shift(-int32(c.USDDecimals()))
	tokenPerETH = tokenPerETH.Shift(-int32(decimals))

	return lo.ToPtr[decimal.Decimal](usdPerETH.Div(tokenPerETH)), nil
}

func (c *Clients) AggregatedTokenPriceInUSD(ctx context.Context, tokenAddresses []common.Address) ([]*decimal.Decimal, error) {
	calls := make([]multicall3.Multicall3Call3, 0)
	var expectLength int
	for _, tokenAddress := range tokenAddresses {
		_tokenCalldata, err := c.contractAbis.Oracle.Pack(constants.GetRateToETHMethod, tokenAddress, true)
		if err != nil {
			return nil, err
		}
		calls = append(calls, multicall3.Multicall3Call3{
			Target:       c.OneInchOracleAddress(),
			AllowFailure: false,
			CallData:     _tokenCalldata,
		})
		_decimalsCallData, err := c.contractAbis.ERC20.Pack(constants.DecimalsMethod)
		if err != nil {
			return nil, err
		}
		calls = append(calls, multicall3.Multicall3Call3{
			Target:       tokenAddress,
			AllowFailure: false,
			CallData:     _decimalsCallData,
		})
		expectLength += 2
	}

	_usdCallData, err := c.contractAbis.Oracle.Pack(constants.GetRateToETHMethod, c.USDAddress(), true)
	if err != nil {
		return nil, err
	}

	calls = append(calls, multicall3.Multicall3Call3{
		Target:       c.OneInchOracleAddress(),
		AllowFailure: false,
		CallData:     _usdCallData,
	})
	expectLength += 1

	results, err := c.AggregatedCalls(ctx, calls)
	if err != nil {
		return nil, err
	}

	if len(results) != expectLength {
		return nil, fmt.Errorf("aggregated price call result do not match the call number")
	}

	for i, result := range results {
		if !result.Success {
			return nil, fmt.Errorf("failed on aggregated price call on %d th call", i)
		}
	}

	var usdPrice *big.Int
	err = c.contractAbis.Oracle.UnpackIntoInterface(&usdPrice, constants.GetRateToETHMethod, results[expectLength-1].ReturnData)
	if err != nil {
		return nil, err
	}
	usdInETH := decimal.NewFromBigInt(usdPrice, -int32(constants.EthereumDecimals))

	prices := make([]*decimal.Decimal, 0)
	zero := decimal.NewFromInt(0)

	for i := 0; i < expectLength-1; i += 2 {
		var tokenPrice *big.Int
		err := c.contractAbis.Oracle.UnpackIntoInterface(&tokenPrice, constants.GetRateToETHMethod, results[i].ReturnData)
		if err != nil {
			return nil, fmt.Errorf("failed on unpack oracle getRateToEth, %w", err)
		}
		var decimals uint8
		err = c.contractAbis.ERC20.UnpackIntoInterface(&decimals, constants.DecimalsMethod, results[i+1].ReturnData)
		if err != nil {
			return nil, fmt.Errorf("failed on unpack erc20 decimals, %w", err)
		}
		tokenInETH := decimal.NewFromBigInt(tokenPrice, -int32(constants.EthereumDecimals))
		if tokenInETH.Equal(zero) {
			prices = append(prices, lo.ToPtr[decimal.Decimal](zero))
			continue
		}

		tokenPerETH := decimal.NewFromInt(1).Div(tokenInETH)
		usdPerETH := decimal.NewFromInt(1).Div(usdInETH)

		usdPerETH = usdPerETH.Shift(-int32(c.USDDecimals()))
		tokenPerETH = tokenPerETH.Shift(-int32(decimals))

		prices = append(prices, lo.ToPtr[decimal.Decimal](usdPerETH.Div(tokenPerETH)))
	}
	return prices, nil
}
