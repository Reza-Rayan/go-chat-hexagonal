package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Jwt      JWTConfig
}

type ServerConfig struct {
	Port       string `mapstructure:"port"`
	Addr       string `mapstructure:"addr"`
	UploadPath string `mapstructure:"upload_path"`
}

type DatabaseConfig struct {
	Path string `mapstructure:"path"`
}

type JWTConfig struct {
	Secret string `mapstructure:"secret"`
}

func LoadConfig() *Config {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading .env file: %v", err)
	}

	v.AutomaticEnv()

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return &cfg
}
