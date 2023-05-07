package model

import (
	"github.com/ethereum/go-ethereum/common"
)

type ERC20Token struct {
	ContractAddress common.Address
	// Name            string
	Symbol   string
	Decimals uint8
}
