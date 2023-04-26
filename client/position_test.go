package client

import (
	"testing"
)

func TestOnAccountHoldings(t *testing.T) {
	_, err := NewClientsWithEndpoints([]string{
		"https://rpc.ankr.com/eth",
	})
	if err != nil {
		t.Fatal(err)
	}

}
