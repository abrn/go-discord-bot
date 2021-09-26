// config.go: defines the Config structure and implements reading and saving a config file
// 	globally across the app
// - also can interact with the Redis cache store and sync values
package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/yaml.v3"
)

// DiscordServer - aka a 'Guild' in the API
// 	GuildObject: https://discord.com/developers/docs/resources/guild#guild-object
type DiscordServer struct {
	Guild *discordgo.Guild  // the Guild API object
	ID int  // the guild ID
	Name string  // the guild name
	Icon string  // the URL of the guild icon image
	Banner string  // the URL of the guild banner imagw
	OwnerID string  // the user ID of the server owner
	Members []*discordgo.UserGuild
}

// The global configuration for the whole Discord bot
type AppConfig struct {
	Token string  // the Discord developer Bot token
	Debug bool  // if to run in debug mode
	Servers []*DiscordServer  // the servers the bot is configured for
	OwnerID string  // the user id of the bot owner account
}

// The marshalled AppConfig parsed from config,yaml
type AppConfigYML struct {
	Discord yaml.Node `yaml:"discord"`
	Token string `yaml:"discord.Token"`
	DebugMode bool `yaml:"discord.DebugMode`
	ServerID string `yaml:"discord.ServerID"`
	OwnerID string `yaml:"discord.OwnerID"`
}

var config *AppConfig = nil

// GetConfig - return the saved config file or load it
func GetConfig() *AppConfig {
	if config == nil {
		config = loadConfig()
	}
	return config
}

// loadConfig - try read and unmarshal the config and parse it
func loadConfig() *AppConfig {
	conf, err := unmarshalFile()
	if err != nil {
		log.Fatalln(err)
	}
	return &AppConfig{
		Token: "test",
		Debug: false,
	}
}

// unmarshalFile - finds the executable path and loads/unmarshalls the config.yml file
func unmarshalFile() (*AppConfig, error) {
	configFile := filepath.Join(BasePath(), "config-files", "config.yaml")
	_, err := os.Stat(configFile); if err != nil {
		// alert how to edit the sample file and then exit
		sendSampleAlert()
		return &AppConfig{}, err
	}
	// attempt to read the config file
	bytes, err := os.ReadFile(configFile); if err != nil {
		log.Fatalf("error opening the 'config.yaml' file: %s\n", err)
		return &AppConfig{}, err
	} else {
		conf := &AppConfigYML{}
		// attempt to parse the config file
		err = yaml.Unmarshal(bytes, conf)
		if err != nil {
			log.Fatalf("error parsing the 'config.yaml' file: %s\n", err)
		}
		// return the unmarshalled config data
		return &AppConfig{
			Token: conf.Token,
			Debug: conf.DebugMode,
			Servers: []*DiscordServer{},
			OwnerID: conf.OwnerID,
		}, nil
	}
}

// sendSampleAlert - send a warning there's a sample template they can use then exit
func sendSampleAlert() {
	log.Printf("could not find the 'config.yaml' file in the 'config-files' folder\n")
	log.Fatalf("edit, rename 'config.sample.yaml' and try again - exiting")
}