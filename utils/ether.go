package utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mytokenio/ethrpc"
	ens "github.com/wealdtech/go-ens/v3"
	"log"
	"time"
)

type Ether struct {
	EthRpc    ethrpc.EthRPC
	StepBlock int //签名过期高度
	Client    *ethclient.Client
}

func NewEther(rpcUri string, stepBlock int) Ether {
	if rpcUri == "" {
		log.Fatalf("rpcUri is nil")
	}
	if stepBlock < 1 {
		stepBlock = 20
	}
	client, err := ethclient.Dial(rpcUri)
	if err != nil {
		panic(err)
	}
	return Ether{
		EthRpc:    ethrpc.NewNodeAPI(rpcUri),
		Client:    client,
		StepBlock: stepBlock,
	}
}
func (s Ether) GetBlockByNum(num int) (*ethrpc.Block, error) {
	return s.EthRpc.EthGetBlockByNumber(num, true)
}

func (s Ether) GetEnsName(address string) (string, error) {
	return ens.ReverseResolve(s.Client, common.HexToAddress(address))
}

//BlockNum 获取以太坊当前高度
func (s Ether) BlockNum() int {
	block, err := s.EthRpc.EthBlockNumber()

	if err != nil {
		fmt.Printf("getBlockNumError: %v \n", err)
		time.Sleep(time.Millisecond * 200) //200 毫秒
		return s.BlockNum()
	}
	return block
}

//ExpiredBlock 获取以太坊当前高度 + 过期高度
func (s Ether) ExpiredBlock() (block int) {
	block = s.BlockNum()
	block = block + s.StepBlock
	return
}
