package constants

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func ERC721TransferHash() common.Hash {
	return crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))
}
