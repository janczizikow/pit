package bot

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Token string `json:"token"`
	App   string `json:"clientId"`
	Guild string `json:"guildId"`
}

// ReadConfig reads the config.json file and unmarshals it into the Config struct
func ReadConfig() (config *Config, err error) {
	// TODO: get rid of fmt + get everything from env
	fmt.Println("Reading config.json...")
	file, err := os.ReadFile("./config.json")

	if err != nil {
		return nil, err
	}

	fmt.Println("Unmarshalling config.json...")

	// unmarshall file into config struct
	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println("Error unmarshalling config.json")
		return nil, err
	}

	return config, nil

}
