package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	CurrentUserName string `json:"current_user_name"`
}

// Read reads the config file from the home directory and returns a Config struct.
func Read() (Config, error) {
	fullPath, err := getConfigPath()
	if err != nil {
		return Config{}, err
	}

	// Read the file from the HOME directory
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return Config{}, err
	}

	// Decode the JSON string into a new Config struct
	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

// SetUser sets the current_user_name field and writes the config back to the JSON file.
func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	return cfg.write()
}

// encode the struct to JSON and saves it
func (cfg *Config) write() error {
	fullPath, err := getConfigPath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(fullPath, data, 0644)
}

// locate the config file.
func getConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, configFileName), nil
}
