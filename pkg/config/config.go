package config

import (
	"fmt"

	"github.com/spf13/viper"
)

//Initconfig Inicializa las configuracion por archivo
func Initconfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("../")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./../../")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
}
