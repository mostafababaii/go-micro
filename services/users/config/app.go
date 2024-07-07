package config

import (
	"fmt"

	"github.com/mostafababaii/go-micro/services/users/db"
	"github.com/spf13/viper"
)

var (
	JsonResponseContentType = "application/json"
	DatabaseConfig          *db.Config
	ServerHost              string
	ServerPort              int32
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	DatabaseConfig = &db.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		User: viper.GetString("db.user"),
		Pass: viper.GetString("db.pass"),
		Name: viper.GetString("db.name"),
	}

	ServerHost = viper.GetString("server.host")
	ServerPort = viper.GetInt32("server.port")
}
