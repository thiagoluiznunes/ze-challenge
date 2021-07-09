package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	AppName    string `mapstructure:"app-name" json:"app-name,omitempty"`
	Debug      bool   `mapstructure:"debug" json:"debug,omitempty"`
	HTTPPort   uint   `mapstructure:"http-port" json:"http-port,omitempty"`
	HTTPPrefix string `mapstructure:"http-prefix" json:"http-prefix,omitempty"`
	DBHost     string `mapstructure:"db-host" json:"db-host,omitempty"`
	DBPort     uint   `mapstructure:"db-port" json:"db-port,omitempty"`
	DBName     string `mapstructure:"db-name" json:"db-name,omitempty"`
	DBUser     string `mapstructure:"db-user" json:"db-user,omitempty"`
	DBPassword string `mapstructure:"db-password" json:"db-password,omitempty"`
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
