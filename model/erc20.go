package model

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

type ERC20Token struct {
	ContractAddress common.Address
	// Name            string
	Symbol       string
	Decimals     uint8
	CurrentPrice decimal.Decimal
}
