package client

import (
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Clients struct {
	ethClients   []*ethclient.Client
	contractAbis *ContractABIs
}

func NewClientsWithEndpoints(endpoints []string) (*Clients, error) {
	ethClients := make([]*ethclient.Client, 0)
	for _, endpoint := range endpoints {
		ethClient, err := ethclient.Dial(endpoint)
		if err != nil {
			return nil, err
		}
		ethClients = append(ethClients, ethClient)
	}
	return &Clients{
		ethClients:   ethClients,
		contractAbis: NewContractAbis(),
	}, nil
}

func (c *Clients) Client() *ethclient.Client {
	if len(c.ethClients) == 0 {
		return nil
	}
	rand.NewSource(time.Now().Unix())
	return c.ethClients[rand.Intn(len(c.ethClients))]
}
