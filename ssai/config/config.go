package config

import (
	"time"

	"github.com/spf13/viper"
)

// Config represents the application configuration
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	LogLevel string
	Version  string
	Appname  string
	DevMode  bool
}

// ServerConfig represents the HTTP server configuration
type ServerConfig struct {
	Address      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

// DatabaseConfig represents the database configuration
type DatabaseConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	Name            string
	SSLMode         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

// Load loads the configuration from environment variables and config files
func Load() (*Config, error) {
	viper.SetDefault("DEV_MODE", true)
	viper.SetDefault("SERVER_ADDRESS", ":9002")
	viper.SetDefault("SERVER_READ_TIMEOUT", 5*time.Second)
	viper.SetDefault("SERVER_WRITE_TIMEOUT", 10*time.Second)
	viper.SetDefault("SERVER_IDLE_TIMEOUT", 120*time.Second)
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", 5432)
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_PASSWORD", "postgres")
	viper.SetDefault("DB_NAME", "service_db")
	viper.SetDefault("DB_SSL_MODE", "disable")
	viper.SetDefault("DB_MAX_OPEN_CONNS", 25)
	viper.SetDefault("DB_MAX_IDLE_CONNS", 25)
	viper.SetDefault("DB_CONN_MAX_LIFETIME", 5*time.Minute)
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("VERSION", "dev")

	// Environment variables
	viper.AutomaticEnv()

	// Config file (optional)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	_ = viper.ReadInConfig() // Ignore error if config file is not found

	return &Config{
		Server: ServerConfig{
			Address:      viper.GetString("SERVER_ADDRESS"),
			ReadTimeout:  viper.GetDuration("SERVER_READ_TIMEOUT"),
			WriteTimeout: viper.GetDuration("SERVER_WRITE_TIMEOUT"),
			IdleTimeout:  viper.GetDuration("SERVER_IDLE_TIMEOUT"),
		},
		Database: DatabaseConfig{
			Host:            viper.GetString("DB_HOST"),
			Port:            viper.GetInt("DB_PORT"),
			User:            viper.GetString("DB_USER"),
			Password:        viper.GetString("DB_PASSWORD"),
			Name:            viper.GetString("DB_NAME"),
			SSLMode:         viper.GetString("DB_SSL_MODE"),
			MaxOpenConns:    viper.GetInt("DB_MAX_OPEN_CONNS"),
			MaxIdleConns:    viper.GetInt("DB_MAX_IDLE_CONNS"),
			ConnMaxLifetime: viper.GetDuration("DB_CONN_MAX_LIFETIME"),
		},
		LogLevel: viper.GetString("LOG_LEVEL"),
		Version:  viper.GetString("VERSION"),
		DevMode:  viper.GetBool("DEV_MODE"),
	}, nil
}
