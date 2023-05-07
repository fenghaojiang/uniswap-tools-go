package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fenghaojiang/uniswap-tools-go/constants"
	"github.com/shopspring/decimal"
)

type NFTPosition struct {
	Nonce                    *big.Int       `json:"nonce"`
	Operator                 common.Address `json:"operator"`
	Token0                   common.Address `json:"token0"`
	Token1                   common.Address `json:"token1"`
	Fee                      *big.Int       `json:"fee"`
	TickLower                *big.Int       `json:"tickLower"`
	TickUpper                *big.Int       `json:"tickUpper"`
	Liquidity                *big.Int       `json:"liquidity"`
	FeeGrowthInside0LastX128 *big.Int       `json:"feeGrowthInside0LastX128"`
	FeeGrowthInside1LastX128 *big.Int       `json:"feeGrowthInside1LastX128"`
	TokensOwed0              *big.Int       `json:"tokensOwed0"`
	TokensOwed1              *big.Int       `json:"tokensOwed1"`
	Pool                     common.Address `json:"-"`
}

type Position struct {
	// Owner   common.Address `json:"owner"`
	Status  constants.Status `json:"status"`
	TokenID *big.Int         `json:"tokenId"`
	Name    string           `json:"name"`

	PriceRangeInToken0 [2]*decimal.Decimal `json:"priceRangeInToken0"`
	PriceRangeInToken1 [2]*decimal.Decimal `json:"priceRangeInToken1"`

	Token0 common.Address `json:"token0"`
	Token1 common.Address `json:"token1"`

	Symbol0 string `json:"symbol0"`
	Symbol1 string `json:"symbol1"`

	LockedAmount0 *decimal.Decimal `json:"lockedAmount0"`
	LockedAmount1 *decimal.Decimal `json:"lockedAmount1"`

	LockedValue0InUSD *decimal.Decimal `json:"lockedValue0InUSD"`
	LockedValue1InUSD *decimal.Decimal `json:"lockedValue1InUSD"`

	FeeRewards0Amount *decimal.Decimal `json:"feeRewards0Amount"`
	FeeRewards1Amount *decimal.Decimal `json:"feeRewards1Amount"`

	FeeRewards0InUSD *decimal.Decimal `json:"feeRewards0"`
	FeeRewards1InUSD *decimal.Decimal `json:"feeRewards1"`

	TotalValueUSD   *decimal.Decimal `json:"tokenValueUSD"`
	TotalRewardsUSD *decimal.Decimal `json:"totalRewardsUSD"`

	Token0Price *decimal.Decimal `json:"token0Price"`
	Token1Price *decimal.Decimal `json:"token1Price"`
}
