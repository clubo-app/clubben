package config

import "github.com/spf13/viper"

type Config struct {
	PORT                       string `mapstructure:"PORT"`
	POSTGRES_URL_PARTY_SERVICE string `mapstructure:"POSTGRES_URL_PARTY_SERVICE"`
	NATS_CLUSTER               string `mapstructure:"NATS_CLUSTER"`
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
