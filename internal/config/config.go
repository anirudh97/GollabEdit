package config

import (
	"log"

	"github.com/spf13/viper"
)

type Conifiguration struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Username string
	Password string
	DBName   string `mapstructure:"dbname"`
	Host     string
	Port     string
}

func Load() (*Conifiguration, error) {
	viper.SetConfigName("default")
	viper.AddConfigPath("./internal/config")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
		return nil, err
	}

	var configurations Conifiguration
	err := viper.Unmarshal(&configurations)
	if err != nil {
		log.Printf("Failed to Configurations %v", err)
		return nil, err
	}

	return &configurations, nil

}
