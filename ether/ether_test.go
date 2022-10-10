package ether

import (
	"github.com/spf13/viper"
	"testing"
)

func TestGetEnsName(t *testing.T) {
	viper.SetConfigFile("../config.yaml")
	viper.ReadInConfig()
	//config.InitConfig()
	//// If a config file is found, read it in.
	//if err := viper.ReadInConfig(); err == nil {
	//	fmt.Println("Using config file:", viper.ConfigFileUsed())
	//}
	ensName, err := GetEther().GetEnsName("0xd0c822E8465Da421c34198f2A98122b350A05Fe1")
	t.Log(ensName)
	t.Log(err)
}
