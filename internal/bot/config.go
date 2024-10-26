package bot

import (
	"os"
)

type Config struct {
	Token string `json:"token"`
	App   string `json:"clientId"`
	Guild string `json:"guildId"`
}

// ReadConfig reads the config.json file and unmarshals it into the Config struct
func ReadConfig() (config *Config, err error) {
	cfg := &Config{}
	token := os.Getenv("DISCORD_BOT_TOKEN")
	clientId := os.Getenv("DISCORD_BOT_CLIENT_ID")
	guildId := os.Getenv("DISCORD_BOT_GUILD_ID")

	cfg.Token = token
	cfg.App = clientId
	cfg.Guild = guildId

	return cfg, nil
}
