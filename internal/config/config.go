// this package is responsible for reading and writing the JSON file
package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
	// get the .gatorconfig.json location
	fullPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	// open the file
	file, err := os.Open(fullPath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	// decode the JSON
	decoder := json.NewDecoder(file)
	data := Config{}
	err = decoder.Decode(&data)
	if err != nil {
		return Config{}, err
	}

	return data, nil
}

func (cfg *Config) SetUser(name string) error {
	// update the name (write it to the Config struct)
	cfg.CurrentUserName = name

	// update the json file (.gatorconfig.json) with updated struct
	err := write(*cfg)
	return err
}

func getConfigFilePath() (string, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	fullPath := filepath.Join(path, configFileName)

	return fullPath, nil
}

func write(cfg Config) error {
	// get the full path
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	// create or overwrite the file
	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// create a new encoder and encode the data into the file
	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
