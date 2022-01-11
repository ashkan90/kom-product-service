package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

// Config ...
type Config interface {
	GetString(key string) string
	GetStringSlice(key string) []string
	GetInt(key string) int
	GetBool(key string) bool
}

type viperConfig struct {}

// New ...
func New() Config {
	v := &viperConfig{}
	v.init()
	return v
}

func (v *viperConfig) init() {
	viper.SetEnvPrefix("go-clean")
	viper.AutomaticEnv()

	viper.SetConfigType("json")
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
}

func (v *viperConfig) GetString(key string) string {
	return viper.GetString(key)
}

func (v *viperConfig) GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}

func (v *viperConfig) GetInt(key string) int {
	return viper.GetInt(key)
}

func (v *viperConfig) GetBool(key string) bool {
	return viper.GetBool(key)
}
