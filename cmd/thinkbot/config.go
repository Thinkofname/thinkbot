package main

import (
	"encoding/json"
	"os"
)

type config struct {
	Server   string   `json:"server"`
	Port     uint16   `json:"port"`
	Username string   `json:"username"`
	Channels []string `json:"channels"`
}

func loadConfig() *config {
	var config config
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

func saveConfig(c *config) {
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

func initDefaults(c *config) {
	c.Server = "irc.example.com"
	c.Port = 6667
	c.Username = "BotName"
	c.Channels = []string{"#banana"}
}
