package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost       string `mapstructure:"DB_HOST"`
	DBPort       string `mapstructure:"DB_PORT"`
	DBUser       string `mapstructure:"DB_USER"`
	DBPassword   string `mapstructure:"DB_PASSWORD"`
	DBName       string `mapstructure:"DB_NAME"`
	ServerPort   string `mapstructure:"SERVER_PORT"`
	DATABASE_URL string `mapstructure:"DATABASE_URL"`
	Port         string `mapstructure:"PORT"`
}

func LoadConfig() (config Config, err error) {

	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development" // デフォルトは開発環境
	}

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	// .envファイルの読み込みを試みるが、失敗してもエラーとしない
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Printf("Config file not found: %v", err)
		}
	}

	// 構造体にUnmarshal
	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	// 環境変数からPORTを取得
	if port := os.Getenv("PORT"); port != "" {
		config.Port = port
	}

	if dbURL := os.Getenv("DATABASE_URL"); dbURL != "" {
		config.DATABASE_URL = dbURL
	}

	return config, nil
}
