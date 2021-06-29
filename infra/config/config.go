package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	AppName    string `mapstructure:"app-name"`
	Debug      bool   `mapstructure:"debug"`
	HTTPPort   uint   `mapstructure:"http-port"`
	HTTPPrefix string `mapstructure:"http-prefix"`
}

func Read() (*Config, error) {

	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../..")
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.SetEnvPrefix("app")
	viper.SetTypeByDefaultValue(true)

	viper.ReadInConfig()

	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
