package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	App App `mapstructure:"app"`
	Db  Db  `mapstructure:"db"`
}

type App struct {
	Port int `mapstructure:"port"`
}

type Db struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

var config *Config

func InitConfig() *Config {

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	// viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("config file %v", err))
	}

	viper.WatchConfig()

	fmt.Printf("address %v \n", &config)
	fmt.Printf("config %v \n", config)

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("after address %v \n", &config)
	fmt.Printf("after config %v \n", config)

	return config
}
