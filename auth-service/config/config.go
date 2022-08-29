package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	PORT                      string `mapstructure:"PORT"`
	NATS_CLUSTER              string `mapstructure:"NATS_CLUSTER"`
	SPACES_ENDPOINT           string `mapstructure:"SPACES_ENDPOINT"`
	SPACES_TOKEN              string `mapstructure:"SPACES_TOKEN"`
	TOKEN_SECRET              string `mapstructure:"TOKEN_SECRET"`
	GOOGLE_CLIENTID           string `mapstructure:"GOOGLE_CLIENTID"`
	POSTGRES_URL_AUTH_SERVICE string `mapstructure:"POSTGRES_URL_AUTH_SERVICE"`
}

func LoadConfig() (config Config) {
	viper.AddConfigPath("./config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.ReadInConfig()

	viper.SetConfigName("prod")
	viper.MergeInConfig()

	viper.AutomaticEnv()

	viper.Unmarshal(&config)
	return
}
