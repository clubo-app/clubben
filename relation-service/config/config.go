package config

import "github.com/spf13/viper"

type Config struct {
	PORT         string `mapstructure:"PORT"`
	CQL_KEYSPACE string `mapstructure:"CQL_KEYSPACE"`
	CQL_HOSTS    string `mapstructure:"CQL_HOSTS"`
	NATS_CLUSTER string `mapstructure:"NATS_CLUSTER"`
}

func LoadConfig() (config Config) {
	viper.AddConfigPath("./config/envs")
	viper.AddConfigPath("../../config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.ReadInConfig()

	viper.SetConfigName("prod")
	viper.MergeInConfig()

	viper.AutomaticEnv()

	viper.Unmarshal(&config)
	return
}
