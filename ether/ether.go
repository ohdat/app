package ether

import (
	"github.com/ohdat/app/utils"
	"github.com/spf13/viper"
	"sync"
)

func initWalletEther(rpcUri string) *utils.Ether {
	var ether = utils.NewEther(rpcUri, 60)
	if utils.IsDev() {
		ether = utils.NewEther(rpcUri, 10)
	}
	return &ether
}

var ether *utils.Ether
var etherOnce sync.Once

func GetEther() *utils.Ether {
	etherOnce.Do(func() {
		ether = initWalletEther(viper.GetString("ethereum.rpc_uri"))
	})
	return ether
}

var polygonEther *utils.Ether
var polygonEtherOnce sync.Once

func GetPolygonEther() *utils.Ether {
	polygonEtherOnce.Do(func() {
		polygonEther = initWalletEther(viper.GetString("polygon_info.rpc_uri"))
	})
	return polygonEther
}
