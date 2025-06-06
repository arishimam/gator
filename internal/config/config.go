package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(userName string) error {
	config, err := Read()
	config.CurrentUserName = userName

	write(config)

	if err != nil {
		return err
	}
	return nil
}

func write(cfg Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		fmt.Println("error getting config file path")
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal("error creating file")
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(cfg)
	if err != nil {
		log.Fatal("error encoding config file")
		return err
	}
	return nil
}

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		fmt.Println("error getting config file path")
		return Config{}, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return Config{}, err
	}
	defer file.Close()

	config := Config{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("error getting home directory")
		return "", err
	}
	configPath := path.Join(homeDir, configFileName)

	return configPath, nil

}
