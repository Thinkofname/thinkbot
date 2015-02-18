package main

import (
	"encoding/json"
	"os"
)

type botConfig struct {
	Server   string                 `json:"server"`
	Port     uint16                 `json:"port"`
	Username string                 `json:"username"`
	Channels []string               `json:"channels"`
	Users    map[string]*userConfig `json:"users"`
}

type userConfig struct {
	Permissions map[string]bool `json:"permissions"`
}

func loadConfig() *botConfig {
	var config botConfig
	initDefaults(&config)
	f, err := os.Open("config.json")
	if err == nil {
		defer f.Close()
		d := json.NewDecoder(f)
		err = d.Decode(&config)
		if err != nil {
			panic(err)
		}
	}
	return &config
}

func saveConfig(c *botConfig) {
	f, err := os.Create("config.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		panic(err)
	}
	f.Write(b)
}

func initDefaults(c *botConfig) {
	c.Server = "irc.example.com"
	c.Port = 6667
	c.Username = "BotName"
	c.Channels = []string{"#banana"}
	c.Users = map[string]*userConfig{
		"oops.i.broke.thinkofdeath.co.uk": {
			Permissions: map[string]bool{
				"bot.admin": true,
			},
		},
	}
}
