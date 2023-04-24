package constants

import "github.com/ethereum/go-ethereum/common"

const (
	uniswapv3NFTPositionManagerAddress = "0xC36442b4a4522E871399CD717aBDD847Ab11FE88"
	uniswapv3NFTDesciptorAddress       = "0x42B24A95702b9986e82d421cC3568932790A48Ec"
	uniswapv3FactoryAddress            = "0x1F98431c8aD98523631AE4a59f267346ea31F984"
	uniswapv3SwapRouterAddress         = "0xE592427A0AEce92De3Edee1F18E0157C05861564"
	uniswapv3SwapRouter02Address       = "0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45"
)

const (
	multicall3Address = "0xcA11bde05977b3631167028862bE2a173976CA11"
)

func UniswapV3NFTPositionManagerAddress() common.Address {
	return common.HexToAddress(uniswapv3NFTPositionManagerAddress)
}

func UniswapV3NFTDescriptorAddress() common.Address {
	return common.HexToAddress(uniswapv3NFTDesciptorAddress)
}

func UniswapV3FacotryAddress() common.Address {
	return common.HexToAddress(uniswapv3FactoryAddress)
}

func UniswapV3SwapRouterAddress() common.Address {
	return common.HexToAddress(uniswapv3SwapRouterAddress)
}

func UniswapV3SwapRouter02Address() common.Address {
	return common.HexToAddress(uniswapv3SwapRouter02Address)
}

func Multicall3Address() common.Address {
	return common.HexToAddress(multicall3Address)
}
