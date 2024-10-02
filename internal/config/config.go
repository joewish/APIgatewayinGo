package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port      string `mapstructure:"port"`
		Timeout   string `mapstructure:"timeout"`
		EnableTLS bool   `mapstructure:"enable_tls"`
		TLSCert   string `mapstructure:"tls_cert_file"`
		TLSKey    string `mapstructure:"tls_key_file"`
	}
	Logging struct {
		Level   string `mapstructure:"level"`
		LogFile string `mapstructure:"log_file"`
	}
	// Add other configurations like auth, rate_limit, etc.
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("../..")  // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil

}
