package multicall_test

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/jon4hz/web3-go/ethrpc"
	"github.com/jon4hz/web3-go/ethrpc/provider/httprpc"
	"github.com/jon4hz/web3-multicall-go/multicall"
)

const rpcURL = "https://ropsten.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"

func TestExampleViewCall(t *testing.T) {
	eth, err := getETH(rpcURL)
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
	mc, _ := multicall.New(eth, multicall.ContractAddress(multicall.RopstenAddress))
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

func getETH(url string) (ethrpc.ETHInterface, error) {
	provider, err := httprpc.New(url)
	if err != nil {
		return nil, err
	}
	provider.SetHTTPTimeout(5 * time.Second)
	return ethrpc.New(provider)
}

func TestUnmarshaltoUint8(t *testing.T) {
	eth, err := getETH(rpcURL)
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
	mc, _ := multicall.New(eth, multicall.ContractAddress(multicall.RopstenAddress))
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
