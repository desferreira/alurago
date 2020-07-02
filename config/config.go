package config

import (
	"fmt"

	"github.com/spf13/viper"
)

/**
Método responsável por ler o arquivo de configuração
*/
func ReadConfig() {
	viper.SetConfigName("application")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
