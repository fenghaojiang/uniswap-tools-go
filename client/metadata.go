package client

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/fenghaojiang/uniswap-tools-go/onchain/generated-go/erc20"
	"github.com/fenghaojiang/uniswap-tools-go/onchain/generated-go/multicall3"
	"github.com/fenghaojiang/uniswap-tools-go/onchain/generated-go/uniswapv3_factory"
	"github.com/fenghaojiang/uniswap-tools-go/onchain/generated-go/uniswapv3_nft_position_manager"
	"github.com/fenghaojiang/uniswap-tools-go/onchain/generated-go/uniswapv3_pool"
	"github.com/fenghaojiang/uniswap-tools-go/onchain/generated-go/uniswapv3_router"
)

type ContractABIs struct {
	Factory            *abi.ABI
	Pool               *abi.ABI
	Router             *abi.ABI
	NftPositionManager *abi.ABI
	Multicall          *abi.ABI
	ERC20              *abi.ABI
}

func NewContractAbis() *ContractABIs {
	factory, _ := uniswapv3_factory.Uniswapv3FactoryMetaData.GetAbi()
	pool, _ := uniswapv3_pool.Uniswapv3PoolMetaData.GetAbi()
	router, _ := uniswapv3_router.Uniswapv3RouterMetaData.GetAbi()
	nftManager, _ := uniswapv3_nft_position_manager.Uniswapv3NftPositionManagerMetaData.GetAbi()
	multilcall, _ := multicall3.Multicall3MetaData.GetAbi()
	erc20, _ := erc20.Erc20MetaData.GetAbi()

	return &ContractABIs{
		Factory:            factory,
		Pool:               pool,
		Router:             router,
		NftPositionManager: nftManager,
		Multicall:          multilcall,
		ERC20:              erc20,
	}
}
