package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	return write(*c)
}

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	dat, err := os.ReadFile(filePath)
	if err != nil {
		err = fmt.Errorf("error parsing file path")
		return Config{}, err
	}

	cfg := Config{}

	err = json.Unmarshal(dat, &cfg)
	if err != nil {
		err = fmt.Errorf("error parsing json")
		return Config{}, err
	}
	return cfg, nil
}

func getConfigFilePath() (string, error) {
	userHome, err := (os.UserHomeDir())
	if err != nil {
		err = fmt.Errorf("error parsing home dir")
		return "", err
	}
	return filepath.Join(userHome, configFileName), nil
}

func write(cfg Config) error {
	dat, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, dat, 0644)
	if err != nil {
		return err
	}
	return nil
}
