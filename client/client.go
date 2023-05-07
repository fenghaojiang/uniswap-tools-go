package client

import (
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/fenghaojiang/uniswap-tools-go/constants"
)

type Clients struct {
	ethClients []*wrapClient
	// limit concurrency of rpc calls
	limitChan chan struct{}
	network   constants.Network

	contractAbis *ContractABIs
}

type wrapClient struct {
	client    *rpc.Client
	ethclient *ethclient.Client
}

func (w *wrapClient) RPCClient() *rpc.Client {
	return w.client
}

func (w *wrapClient) ETHClient() *ethclient.Client {
	return w.ethclient
}

func NewClientsWithEndpoints(endpoints []string) (*Clients, error) {
	ethClients := make([]*wrapClient, 0)
	for _, endpoint := range endpoints {
		callClient, err := rpc.Dial(endpoint)
		if err != nil {
			return nil, err
		}

		ethClient, err := ethclient.Dial(endpoint)
		if err != nil {
			return nil, err
		}
		ethClients = append(ethClients, &wrapClient{
			client:    callClient,
			ethclient: ethClient,
		})
	}
	return &Clients{
		ethClients:   ethClients,
		limitChan:    make(chan struct{}, 10),
		contractAbis: NewContractAbis(),
	}, nil
}

func (c *Clients) WithLimitRPC(limit int) {
	c.limitChan = make(chan struct{}, limit)
}

func (c *Clients) WithNetwork(network constants.Network) *Clients {
	c.network = network
	return c
}

func (c *Clients) Client() *wrapClient {
	if len(c.ethClients) == 0 {
		return nil
	}
	rand.NewSource(time.Now().Unix())
	return c.ethClients[rand.Intn(len(c.ethClients))]
}
