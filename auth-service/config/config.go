package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	PORT                           string `mapstructure:"PORT"`
	NATS_CLUSTER                   string `mapstructure:"NATS_CLUSTER"`
	GOOGLE_APPLICATION_CREDENTIALS string `mapstructure:"GOOGLE_APPLICATION_CREDENTIALS"`
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
