package utils

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"github.com/wealdtech/go-ens/v3"
	"log"
	"math/big"
	"time"
)

type Ether struct {
	StepBlock int //签名过期高度
	Client    *ethclient.Client
}

func NewEther(uri string, stepBlock int) Ether {
	if uri == "" {
		log.Fatalf("uri is nil")
	}
	if stepBlock < 1 {
		stepBlock = 20
	}
	client, err := ethclient.Dial(uri)
	if err != nil {
		panic(err)
	}
	return Ether{
		Client:    client,
		StepBlock: stepBlock,
	}
}
func (s Ether) GetBlockByNum() (interface{}, error) {
	log.Println("[Error]", "GetBlockByNum is deprecated, use GetBlockByNumber instead")
	return nil, errors.New("GetBlockByNum is deprecated, use  Client.GetBlockByNumber instead")
}

func (s Ether) GetEnsName(address string) (string, error) {
	return ens.ReverseResolve(s.Client, common.HexToAddress(address))
}

//BlockNum 获取以太坊当前高度
func (s Ether) BlockNum() int {
	var ctx = context.Background()
	block, err := s.Client.BlockNumber(ctx)
	if err != nil {
		fmt.Printf("getBlockNumError: %v \n", err)
		time.Sleep(time.Millisecond * 200) //200 毫秒
		return s.BlockNum()
	}
	return int(block)
}

//ExpiredBlock 获取以太坊当前高度 + 过期高度
func (s Ether) ExpiredBlock() (block int) {
	block = s.BlockNum()
	block = block + s.StepBlock
	return
}

func (s Ether) Balance(address string) (balance string, err error) {
	var (
		amount *big.Int
		ctx    = context.Background()
	)
	//DECIMAL
	amount, err = s.Client.BalanceAt(ctx, common.HexToAddress(address), nil)
	decimal.NewFromBigInt(amount, 18)
	balance = amount.String()
	return
}
