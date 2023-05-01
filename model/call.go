package model

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

type CallContractParam struct {
	To       string `json:"to"`
	CallData string `json:"data"`
}

func NewCallContractParam(address string, calldata []byte) (CallContractParam, error) {
	if len(address) == 0 {
		return CallContractParam{}, fmt.Errorf("contract address can not be empty")
	}
	data := hexutil.Encode(calldata)
	return CallContractParam{
		To:       address,
		CallData: data,
	}, nil
}
