package config

import (
    "fmt"
    "os"
    "path/filepath"

    "github.com/spf13/viper"
)

type Config struct {
	AppName    string
	ServerPort string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode string
	JWTSecret string
	RedisHost string
	RedisPort string
	RedisUsername string
	RedisPassword string
	RedisDB int
}

func Load() *Config {
    configPath := os.Getenv("APP_CONFIG_FILE")
    if configPath == "" {
        configPath = ".env"
    }

    viper.SetConfigFile(configPath)

    if err := viper.ReadInConfig(); err != nil {
        absPath, _ := filepath.Abs(configPath)
        panic(fmt.Sprintf(
            "failed to read config file %q: %v",
            absPath,
            err,
        ))
    }

    return &Config{
        AppName:       viper.GetString("APP_NAME"),
        ServerPort:    viper.GetString("SERVER_PORT"),
        DBHost:        viper.GetString("DB_HOST"),
        DBPort:        viper.GetString("DB_PORT"),
        DBUser:        viper.GetString("DB_USER"),
        DBPassword:    viper.GetString("DB_PASSWORD"),
        DBName:        viper.GetString("DB_NAME"),
        DBSSLMode:     viper.GetString("DB_SSLMODE"),
        JWTSecret:     viper.GetString("JWT_SECRET"),
        RedisHost:     viper.GetString("REDIS_HOST"),
        RedisPort:     viper.GetString("REDIS_PORT"),
        RedisUsername: viper.GetString("REDIS_USERNAME"),
        RedisPassword: viper.GetString("REDIS_PASSWORD"),
        RedisDB:       viper.GetInt("REDIS_DB"),
    }
}