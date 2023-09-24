package config

import (
	"github.com/anazibinurasheed/dmart-inventory-svc/internal/util"
	"github.com/spf13/viper"
)

type Config struct {
	MongoUrl string `mapstructure:"MONGO_URL"`
	Port     string `mapstructure:"PORT"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./internal/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()

	if util.HasError(err) {
		return
	}

	err = viper.Unmarshal(&config)
	if util.HasError(err) {
		return
	}
	return
}
