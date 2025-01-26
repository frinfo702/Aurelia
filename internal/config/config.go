package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost string
	DBPort int
	DBUser string
	DBPass string
	DBName string

	AppPort int
}

// InitConfig はアプリ起動時に一度呼び出す。envファイルを読み込み、構造体へ格納
func InitConfig() *Config {
	// .envを読み込む(なければスキップ)
	_ = godotenv.Load()

	c := &Config{}

	c.DBHost = getEnv("DB_HOST", "localhost")
	c.DBPort = getEnvAsInt("DB_PORT", 5432)
	c.DBUser = getEnv("DB_USERNAME", "postgres")
	c.DBPass = getEnv("DB_PASSWORD", "postgres")
	c.DBName = getEnv("DB_NAME", "aurelia_db")

	c.AppPort = getEnvAsInt("APP_PORT", 8080)

	return c
}

func (c *Config) GetDBConnString() string {
	// PostgreSQL例
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
		c.DBHost, c.DBPort, c.DBUser, c.DBPass, c.DBName,
	)
}

// 環境変数取得用のユーティリティ
func getEnv(key, defaultVal string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return defaultVal
}

func getEnvAsInt(key string, defaultVal int) int {
	if valStr, exists := os.LookupEnv(key); exists {
		val, err := strconv.Atoi(valStr)
		if err != nil {
			log.Printf("Invalid value for %s: %v", key, err)
			return defaultVal
		}
		return val
	}
	return defaultVal
}
