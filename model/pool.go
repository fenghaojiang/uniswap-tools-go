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

type Slot0 struct {
	SqrtPriceX96               *big.Int `json:"sqrtPriceX96"`
	Tick                       *big.Int `json:"tick"`
	ObservationIndex           uint16   `json:"observationIndex"`
	ObservationCardinality     uint16   `json:"observationCardinality"`
	ObservationCardinalityNext uint16   `json:"observationCardinalityNext"`
	FeeProtocol                uint8    `json:"feeProtocol"`
	Unlocked                   bool     `json:"unlocked"`
}

type Tick struct {
	LiquidityGross                 *big.Int `json:"liquidityGross"`
	LiquidityNet                   *big.Int `json:"liquidityNet"`
	FeeGrowthOutside0X128          *big.Int `json:"feeGrowthOutside0X128"`
	FeeGrowthOutside1X128          *big.Int `json:"feeGrowthOutside1X128"`
	TickCumulativeOutside          *big.Int `json:"tickCumulativeOutside"`
	SecondsPerLiquidityOutsideX128 *big.Int `json:"secondsPerLiquidityOutsideX128"`
	SecondsOutside                 uint32   `json:"secondsOutside"`
	Initialized                    bool     `json:"initialized"`
}

type PoolAggregated struct {
	FeeGrowthGlobal0X128 *big.Int
	FeeGrowthGlobal1X128 *big.Int
	Liquidity            *big.Int
	Slot0                *Slot0
	TickLowerTicks       *Tick
	TickUpperTicks       *Tick
}
