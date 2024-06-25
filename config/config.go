package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	HOST            string
	PORT            string
	DBUserName      string        `mapstructure:"POSTGRES_USER"`
	DBUserPassword  string        `mapstructure:"POSTGRES_PASSWORD"`
	DBName          string        `mapstructure:"POSTGRES_DB"`
	DBPort          string        `mapstructure:"POSTGRES_PORT"`
	MaxIdleConns    int           // MaxIdleConns for database connection pool
	MaxOpenConns    int           // MaxOpenConns for database connection pool
	ConnMaxLifetime time.Duration // ConnMaxLifetime for database connection pool
	SchemaName      string        `mapstructure:"SCHEMA_NAME"`
}

func NewConfig() *Config {
	config := &Config{}
	viper.SetConfigName("config-debug")
	viper.SetConfigType("env")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %v", err))
	} else {
		viper.Unmarshal(config)
	}
	return config
}
