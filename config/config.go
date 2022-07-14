package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var CfgFile string

// InitConfig reads in config file and ENV variables if set.
func InitConfig() {
	if CfgFile != "" {
		viper.SetConfigFile(CfgFile)
	} else {
		// Search config in home directory with name .config" (without extension).
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml") //设置配置文件类型，可选
	}
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	viper.SetDefault("jwt.private", viper.GetString("JWT_PRIVATE"))
	viper.SetDefault("jwt.public", viper.GetString("JWT_PUBLIC"))
}
