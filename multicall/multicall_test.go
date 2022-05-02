package multicall_test

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jon4hz/geth-multicall/multicall"
)

func TestExampleViewCall(t *testing.T) {
	eth, err := getETH("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}
	vc := multicall.NewViewCall(
		"key.1",
		"0x52c9F319990395a214bf45E73D6ee86B85D69fde",
		"totalSupply()(uint256)",
		[]interface{}{},
	)
	vcs := multicall.ViewCalls{vc}
	mc, _ := multicall.New(eth, multicall.WithContractAddress(multicall.RopstenAddress))
	block := "latest"
	res, err := mc.Call(vcs, block)
	if err != nil {
		panic(err)
	}

	resJson, _ := json.Marshal(res)
	fmt.Println(string(resJson))
	fmt.Println(res)
	fmt.Println(res.Calls["key.1"].Decoded[0].(*big.Int))
	fmt.Println(err)
}

func getETH(url string) (*ethclient.Client, error) {
	return ethclient.Dial(url)
}

func TestUnmarshaltoUint8(t *testing.T) {
	eth, err := getETH("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}
	vc := multicall.NewViewCall(
		"key.1",
		"0xc715abcd34c8ed9ebbf95990e0c43401fbbc122d",
		"decimals()(uint8)",
		[]interface{}{},
	)
	vcs := multicall.ViewCalls{vc}
	mc, _ := multicall.New(eth, multicall.WithContractAddress(multicall.RopstenAddress))
	block := "latest"
	res, err := mc.Call(vcs, block)
	if err != nil {
		panic(err)
	}

	resJson, _ := json.Marshal(res)
	fmt.Println(string(resJson))
	fmt.Println(res)
	fmt.Println(res.Calls["key.1"].Decoded[0].(uint8))
	fmt.Println(err)
}

func TestUnmarshaltoUint8OoverWS(t *testing.T) {
	eth, err := getETH("ws://127.0.0.1:8546")
	if err != nil {
		panic(err)
	}
	vc := multicall.NewViewCall(
		"key.1",
		"0xc715abcd34c8ed9ebbf95990e0c43401fbbc122d",
		"decimals()(uint8)",
		[]interface{}{},
	)
	vcs := multicall.ViewCalls{vc}
	mc, _ := multicall.New(eth, multicall.WithContractAddress(multicall.RopstenAddress))
	block := "latest"
	res, err := mc.Call(vcs, block)
	if err != nil {
		panic(err)
	}

	resJson, _ := json.Marshal(res)
	fmt.Println(string(resJson))
	fmt.Println(res)
	fmt.Println(res.Calls["key.1"].Decoded[0].(uint8))
	fmt.Println(err)
}
