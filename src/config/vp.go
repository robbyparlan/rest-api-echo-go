package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

//InitConfig ...
func InitConfig() {
	viper.SetConfigName("config") // name of config file (without extension)
	argsLen := len(os.Args)
	if argsLen > 1 {
		if os.Args[1] == "pg" {
			viper.SetConfigName("config-pg") // name of config file (without extension)
		}
	}
	viper.AddConfigPath(".") // optionally look for config in the working directory

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Panic(fmt.Errorf("fatal error reading config file: %s", err))
	}
}