package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

// SetUser sets the current_user_name field and writes the config back to the JSON file.
func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	return cfg.write()
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
	defer data.Close()

	// Decode the JSON string into a new Config struct
	decoder := json.NewDecoder(data)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

// locate the config file.
func getConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, configFileName), nil
}

// encode the struct to JSON and saves it
func write(cfg Config) error {
	fullPath, err := getConfigPath()
	if err != nil {
		return err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
