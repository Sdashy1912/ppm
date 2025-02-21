package configs

import (
	"github.com/spf13/viper"
)

// InitConfig initialize configurations
func InitConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	// gopath := os.Getenv("GOPATH")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	return viper.ReadInConfig()
}
