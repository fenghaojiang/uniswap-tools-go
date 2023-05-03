package constants

type Network string

const (
	EthereumNetwork          Network = "ethereum"
	PolygonNetwork           Network = "polygon"
	ArbitrumNetwork          Network = "arbitrum"
	OptimismNetwork          Network = "optimism"
	BinanceSmartChainNetwork Network = "binanace_smart_chain"
)

func (n Network) String() string {
	return string(n)
}
