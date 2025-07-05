package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	configPath = "PATH_CONFIG"
)

func MustLoad() *Config {
	cfg, err := Load()
	if err != nil {
		panic(err)
	}
	return cfg
}

func Load() (*Config, error) {
	const op = "config.Load"

	path := os.Getenv(configPath)
	if path == "" {
		return nil, fmt.Errorf("%s: CONFIG_PATH is not set", op)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("%s: %s does not exist", op, path)
	}

	cfg := &Config{}

	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return cfg, nil
}

func MustLoadWithEnv() *Config {
	cfg, err := LoadWithEnv()
	if err != nil {
		panic(err)
	}
	return cfg
}

func LoadWithEnv() (*Config, error) {
	const op = "config.LoadWithEnv"

	path := os.Getenv(configPath)
	if path == "" {
		return nil, fmt.Errorf("%s: CONFIG_PATH is not set", op)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("%s: %s does not exist", op, path)
	}

	cfg := &Config{}

	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return cfg, nil
}

func MustLoadLogger() *LoggerConfig {
	cfg, err := LoadLogger()
	if err != nil {
		panic(err)
	}
	return cfg
}

func LoadLogger() (*LoggerConfig, error) {
	const op = "config.LoadLogger"

	path := os.Getenv(configPath)
	if path == "" {
		return nil, fmt.Errorf("%s: CONFIG_PATH is not set", op)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("%s: %s does not exist", op, path)
	}

	cfg := struct {
		Logger LoggerConfig `yaml:"logger"`
	}{}

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return &cfg.Logger, nil
}
