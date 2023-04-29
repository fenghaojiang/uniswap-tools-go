package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

type NFTPosition struct {
	Nonce                    decimal.Decimal `json:"nonce"`
	Operator                 common.Address  `json:"operator"`
	Token0                   common.Address  `json:"token0"`
	Token1                   common.Address  `json:"token1"`
	Fee                      int64           `json:"fee"`
	TickLower                int64           `json:"tickLower"`
	TickUpper                int64           `json:"tickUpper"`
	Liquidity                decimal.Decimal `json:"liquidity"`
	FeeGrowthInside0LastX128 decimal.Decimal `json:"feeGrowthInside0LastX128"`
	FeeGrowthInside1LastX128 decimal.Decimal `json:"feeGrowthInside1LastX128"`
	TokensOwed0              decimal.Decimal `json:"tokensOwed0"`
	TokensOwed1              decimal.Decimal `json:"tokensOwed1"`
}

type Position struct {
	Owner   common.Address `json:"owner"`
	TokenID *big.Int       `json:"tokenId"`
	Name    string         `json:"name"`

	Token0 common.Address `json:"token0"`
	Token1 common.Address `json:"token1"`

	Symbol0 string `json:"symbol0"`
	Symbol1 string `json:"symbol1"`

	LockedAmount0     *decimal.Decimal `json:"lockedAmount0"`
	LockedAmount1     *decimal.Decimal `json:"lockedAmount1"`
	LockedValue0InUSD *decimal.Decimal `json:"lockedValue0InUSD"`
	LockedValue1InUSD *decimal.Decimal `json:"lockedValue1InUSD"`

	FeeRewards0InUSD *decimal.Decimal `json:"feeRewards0"`
	FeeRewards1InUSD *decimal.Decimal `json:"feeRewards1"`

	FeeRewards0Amount *decimal.Decimal `json:"feeRewards0Amount"`
	FeeRewards1Amount *decimal.Decimal `json:"feeRewards1Amount"`

	TotalRewardsUSD *decimal.Decimal `json:"totalRewardsUSD"`
}
