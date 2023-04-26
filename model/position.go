package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

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
