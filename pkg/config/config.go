package config

import (
	"fmt"
	"time"
)

// ServerConfig represents server configuration
type ServerConfig struct {
	Host         string        `yaml:"host" env:"HOST" env-default:"localhost"`
	Port         string        `yaml:"port" env:"PORT" env-default:"8080"`
	ReadTimeout  time.Duration `yaml:"read_timeout" env:"READ_TIMEOUT" env-default:"10s"`
	WriteTimeout time.Duration `yaml:"write_timeout" env:"WRITE_TIMEOUT" env-default:"30s"`
	IdleTimeout  time.Duration `yaml:"idle_timeout" env:"IDLE_TIMEOUT" env-default:"1m"`
}

// DatabaseConfig represents database configuration
type DatabaseConfig struct {
	Host     string `yaml:"host" env:"DB_HOST" env-default:"localhost"`
	Port     string `yaml:"port" env:"DB_PORT" env-default:"5432"`
	User     string `yaml:"user" env:"DB_USER" env-default:"postgres"`
	Password string `yaml:"password" env:"DB_PASSWORD" env-default:"postgres"`
	DBName   string `yaml:"dbname" env:"DB_NAME" env-default:"fileshare"`
	SSLMode  string `yaml:"sslmode" env:"DB_SSLMODE" env-default:"disable"`
}

// JWTConfig represents JWT configuration
type JWTConfig struct {
	SecretKey string        `yaml:"secret_key" env:"JWT_SECRET" env-default:"your-secret-key"`
	Duration  time.Duration `yaml:"duration" env:"JWT_DURATION" env-default:"24h"`
}

// StorageConfig represents storage configuration
type StorageConfig struct {
	Path string `yaml:"path" env:"STORAGE_PATH" env-default:"./storage"`
}

// LoggerConfig represents logger configuration
type LoggerConfig struct {
	Level string `yaml:"level" env:"LOG_LEVEL" env-default:"info"`
}

// Config represents the application configuration
type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	JWT      JWTConfig      `yaml:"jwt"`
	Storage  StorageConfig  `yaml:"storage"`
	Logger   LoggerConfig   `yaml:"logger"`
}

// DSN returns the database connection string
func (db *DatabaseConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		db.Host,
		db.Port,
		db.User,
		db.Password,
		db.DBName,
		db.SSLMode,
	)
}

// ServerAddr returns the server address
func (srv *ServerConfig) ServerAddr() string {
	return fmt.Sprintf("%s:%s", srv.Host, srv.Port)
}
