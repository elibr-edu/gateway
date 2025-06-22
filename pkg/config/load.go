package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

func MustLoad() *Config {
	cfg, err := Load()
	if err != nil {
		panic(err)
	}
	return cfg
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

	path := os.Getenv("CONFIG_PATH")
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

func Load() (*Config, error) {
	const op = "config.Load"

	path := os.Getenv("CONFIG_PATH")
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
