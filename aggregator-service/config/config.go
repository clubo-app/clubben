package config

import "github.com/spf13/viper"

type Config struct {
	PORT                           string `mapstructure:"PORT"`
	PROFILE_SERVICE_ADDRESS        string `mapstructure:"PROFILE_SERVICE_ADDRESS"`
	AUTH_SERVICE_ADDRESS           string `mapstructure:"AUTH_SERVICE_ADDRESS"`
	PARTY_SERVICE_ADDRESS          string `mapstructure:"PARTY_SERVICE_ADDRESS"`
	STORY_SERVICE_ADDRESS          string `mapstructure:"STORY_SERVICE_ADDRESS"`
	RELATION_SERVICE_ADDRESS       string `mapstructure:"RELATION_SERVICE_ADDRESS"`
	COMMENT_SERVICE_ADDRESS        string `mapstructure:"COMMENT_SERVICE_ADDRESS"`
	PARTICIPATION_SERVICE_ADDRESS  string `mapstructure:"PARTICIPATION_SERVICE_ADDRESS"`
	SEARCH_SERVICE_ADDRESS         string `mapstructure:"SEARCH_SERVICE_ADDRESS"`
	GOOGLE_APPLICATION_CREDENTIALS string `mapstructure:"GOOGLE_APPLICATION_CREDENTIALS"`
}

func LoadConfig() (config Config, err error) {
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
