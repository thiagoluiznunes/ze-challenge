package config

import (
	"os"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/spf13/viper"
)

type Config struct {
	AppName                  string `mapstructure:"app-name" json:"app-name,omitempty"`
	Debug                    bool   `mapstructure:"debug" json:"debug,omitempty"`
	HTTPPort                 uint   `mapstructure:"http-port" json:"http-port,omitempty"`
	HTTPPrefix               string `mapstructure:"http-prefix" json:"http-prefix,omitempty"`
	DBHost                   string `mapstructure:"db-host" json:"db-host,omitempty"`
	DBPort                   uint   `mapstructure:"db-port" json:"db-port,omitempty"`
	DBName                   string `mapstructure:"db-name" json:"db-name,omitempty"`
	DBUser                   string `mapstructure:"db-user" json:"db-user,omitempty"`
	DBPassword               string `mapstructure:"db-password" json:"db-password,omitempty"`
	NewRelicApplicationName  string `mapstructure:"new-relic-application-name"`
	NewRelicLicenseKey       string `mapstructure:"new-relic-license-key"`
	NewRelicEnabled          bool   `mapstructure:"new-relic-enabled"`
	NewRelicErrorCollecting  bool   `mapstructure:"new-relic-error-collecting"`
	DistributedTracerEnabled bool   `mapstructure:"new-relic-distributed-tracer"`
}

func Read() (*Config, error) {

	var config Config
	var paramatersPath = os.Getenv("APP_PARAMETERS_PATH")

	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../..")
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.SetEnvPrefix("app")
	viper.SetTypeByDefaultValue(true)

	if paramatersPath != "" {
		err := SetSecrets("sa-east-1", paramatersPath, GetKeys(config))
		if err != nil {
			return nil, err
		}
	} else {
		viper.ReadInConfig()
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func GetKeys(config Config) (keyNames []string) {

	val := reflect.ValueOf(config)
	for i := 0; i < val.Type().NumField(); i++ {
		keyNames = append(keyNames, val.Type().Field(i).Tag.Get("mapstructure"))
	}

	return keyNames
}

func SetSecrets(region string, path string, keynames []string) (err error) {

	for i := range keynames {
		keynames[i] = path + keynames[i]
	}

	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String(region)},
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return err
	}
	ssmsvc := ssm.New(sess, aws.NewConfig().WithRegion(region))

	params, err := ssmsvc.GetParameters(&ssm.GetParametersInput{
		Names:          aws.StringSlice(keynames),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		return err
	}

	for _, param := range params.Parameters {
		name := strings.ReplaceAll(*param.Name, path, "")
		viper.Set(name, *param.Value)
	}

	return nil
}

// func initConfig() {

// 	err := setViperDefaults(Config)
// 	if err != nil {
// 		cobra.CheckErr(err)
// 	}

// 	viper.SetTypeByDefaultValue(true)
// 	viper.AddConfigPath(".")
// 	viper.SetConfigFile(".env")

// 	viper.ReadInConfig()
// 	viper.AutomaticEnv()

// 	err = viper.Unmarshal(&Config)
// 	cobra.CheckErr(err)
// }

// func setViperDefaults(cfg infrastructure.Config) (err error) {

// 	cfgMap := make(map[string]interface{})
// 	err = mapstructure.Decode(cfg, &cfgMap)
// 	if err != nil {
// 		return err
// 	}

// 	err = viper.MergeConfigMap(cfgMap)

// 	return err
// }
