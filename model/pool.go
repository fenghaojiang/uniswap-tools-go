package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Pool struct {
	Token0 common.Address `json:"token0"`
	Token1 common.Address `json:"token1"`
	Fee    *big.Int       `json:"fee"`
	Pool   common.Address `json:"pool"`
}
