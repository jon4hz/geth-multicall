package multicall

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"unsafe"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/jon4hz/web3-go/ethrpc"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Multicall interface {
	CallRaw(calls ViewCalls, block string) (*Result, error)
	Call(calls ViewCalls, block string) (*Result, error)
	Contract() string
}

type multicall struct {
	rpc    *rpc.Client
	config *Config
}

func New(eth *ethclient.Client, opts ...Option) (Multicall, error) {
	config := &Config{
		MulticallAddress: MainnetAddress,
		Gas:              "0x400000000",
	}

	clientValue := reflect.ValueOf(eth).Elem()
	fieldStruct := clientValue.FieldByName("c")
	clientPointer := reflect.NewAt(fieldStruct.Type(), unsafe.Pointer(fieldStruct.UnsafeAddr())).Elem()
	rpcClient, ok := clientPointer.Interface().(*rpc.Client)
	if !ok {
		return nil, fmt.Errorf("failed to get rpc client")
	}

	for _, opt := range opts {
		opt(config)
	}

	return &multicall{
		config: config,
		rpc:    rpcClient,
	}, nil
}

type CallResult struct {
	Success bool
	Raw     []byte
	Decoded []interface{}
}

type Result struct {
	BlockNumber uint64
	Calls       map[string]CallResult
}

const AggregateMethod = "0x17352e13"

func (mc multicall) CallRaw(calls ViewCalls, block string) (*Result, error) {
	resultRaw, err := mc.makeRequest(calls, block)
	if err != nil {
		return nil, err
	}
	return calls.decodeRaw(resultRaw)
}

func (mc multicall) Call(calls ViewCalls, block string) (*Result, error) {
	resultRaw, err := mc.makeRequest(calls, block)
	if err != nil {
		return nil, err
	}
	return calls.decode(resultRaw)
}

func (mc multicall) makeRequest(calls ViewCalls, block string) (string, error) {
	payloadArgs, err := calls.callData()
	if err != nil {
		return "", err
	}
	payload := make(map[string]string)
	payload["to"] = mc.config.MulticallAddress
	payload["data"] = AggregateMethod + hex.EncodeToString(payloadArgs)
	payload["gas"] = mc.config.Gas
	var resultRaw string
	err = mc.rpc.Call(&resultRaw, ethrpc.ETHCall, payload, block)
	return resultRaw, err
}

func (mc multicall) Contract() string {
	return mc.config.MulticallAddress
}
