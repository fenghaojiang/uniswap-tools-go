
# uniswap-tools-go  



## Goals  

`uniswap-tools-go` aims to provide you a neat multi-chain solution for checking your holding positions and profit in `uniswap-v3`. You can request for anyone's uniswap-v3's portfolio using this toolkit in `Ethereum / Polygon / Arbitrum / Optimism / BSC`.  

## Non-goals 

This project does not make money for you. It does not contain any logic that predict or something relevant to trading. 


## Requirements
- Go Version: 1.19+

## Get Start
```shell
go get github.com/fenghaojiang/uniswap-tools-go
```



## Example  

```go
package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/fenghaojiang/uniswap-tools-go/client"
	"github.com/fenghaojiang/uniswap-tools-go/constants"
)

func main() {
	polygonClis, err := client.NewClientsWithEndpoints([]string{
		"https://rpc.ankr.com/polygon",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
    
	results, err := polygonClis.WithLimitRPC(100).WithNetwork(constants.PolygonNetwork).AggregatedPosition(context.Background(), []*big.Int{
		new(big.Int).SetInt64(869899),
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	for i := range results {
		fmt.Printf("%+v\n", results[i])
	}
}

```


## Reference  

- [Uniswap V3 Whitepaper](https://uniswap.org/whitepaper-v3.pdf)
- [Uniswap V3 Liquidity Math](https://atiselsts.github.io/pdfs/uniswap-v3-liquidity-math.pdf)  


## Contribution
Any suggestions, comments (including criticisms) and contributions are welcome.


