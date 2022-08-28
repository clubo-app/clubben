package config

import "github.com/spf13/viper"

type Config struct {
	PORT                         string `mapstructure:"PORT"`
	NATS_CLUSTER                 string `mapstructure:"NATS_CLUSTER"`
	SPACES_ENDPOINT              string `mapstructure:"SPACES_ENDPOINT"`
	SPACES_TOKEN                 string `mapstructure:"SPACES_TOKEN"`
	POSTGRES_URL_PROFILE_SERVICE string `mapstructure:"POSTGRES_URL_PROFILE_SERVICE"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
