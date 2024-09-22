package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	// DBHost     string `mapstructure:"DB_HOST"`
	// DBPort     string `mapstructure:"DB_PORT"`
	// DBUser     string `mapstructure:"DB_USER"`
	// DBPassword string `mapstructure:"DB_PASSWORD"`
	// DBName     string `mapstructure:"DB_NAME"`
	// ServerPort string `mapstructure:"SERVER_PORT"`
	DATABASE_URL string `mapstructure:"DATABASE_URL"`
	Port         string `mapstructure:"PORT"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	// .envファイルの読み込みを試みるが、失敗してもエラーとしない
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// ファイルが見つからないエラー以外の場合はログに記録するか、エラーを返す
			// log.Printf("Config file not found: %v", err)
		}
	}

	// 環境変数からPORTを取得
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // デフォルトポート
	}

	// 構造体にUnmarshal
	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	return config, nil
}
