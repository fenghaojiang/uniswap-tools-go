package client

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fenghaojiang/uniswap-tools-go/model"
)

func (c *Clients) ERC20Token(ctx context.Context, address common.Address) (model.ERC20Token, error) {

	// TODO
	return model.ERC20Token{}, nil
}
