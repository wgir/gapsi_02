package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	AppEnv     string `mapstructure:"APP_ENV"`
	ServerPort string `mapstructure:"SERVER_PORT"`

	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     int    `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBSSLMode  string `mapstructure:"DB_SSLMODE"`

	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JWTAccessTTL  int    `mapstructure:"JWT_ACCESS_TTL_MINUTES"`
	JWTRefreshTTL int    `mapstructure:"JWT_REFRESH_TTL_DAYS"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(fmt.Sprintf("%s/.env", path))
	viper.AutomaticEnv()

	// Default values
	viper.SetDefault("APP_ENV", "development")
	viper.SetDefault("SERVER_PORT", "8088")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", 5432)
	viper.SetDefault("DB_USER", "root")
	viper.SetDefault("DB_NAME", "orders_db")
	viper.SetDefault("DB_SSLMODE", "disable")
	viper.SetDefault("JWT_ACCESS_TTL_MINUTES", 15)
	viper.SetDefault("JWT_REFRESH_TTL_DAYS", 7)

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// File found but couldn't be read
			fmt.Printf("Warning: couldn't read .env file: %v\n", err)
		}
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}
