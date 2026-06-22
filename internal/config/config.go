package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	Repositories []string `json:"repositories"`
}

func GetConfigPath() (string, error) {
	home, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	return filepath.Join(
		home,
		".config",
		"forge",
		"config.json",
	), nil
}

func EnsureConfigDir() error {
	home, err := os.UserHomeDir()

	if err != nil {
		return err
	}

	return os.MkdirAll(
		filepath.Join(home, ".config", "forge"),
		0755,
	)
}

func CreateDefaultConfig() error {
	cfg := Config{
		Repositories: []string{},
	}

	path, err := GetConfigPath()

	if err != nil {
		return err
	}

	file, err := os.Create(path)

	if err != nil {
		return err
	}

	defer file.Close()

	return json.NewEncoder(file).Encode(cfg)
}

func LoadConfig() (*Config, error) {
	path, err := GetConfigPath()

	if err != nil {
		return nil, err
	}

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var cfg Config

	err = json.NewDecoder(file).Decode(&cfg)

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func SaveConfig(cfg *Config) error {
	path, err := GetConfigPath()

	if err != nil {
		return err
	}

	file, err := os.Create(path)

	if err != nil {
		return err
	}

	defer file.Close()

	return json.NewEncoder(file).Encode(cfg)
}

func AddRepository(path string) error {
	cfg, err := LoadConfig()

	if err != nil {
		return err
	}

	for _, repo := range cfg.Repositories {
		if repo == path {
			return nil
		}
	}

	cfg.Repositories = append(
		cfg.Repositories,
		path,
	)

	return SaveConfig(cfg)
}