/*
 * Copyright 2015 Matthew Collins
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"encoding/json"
	"os"
)

type botConfig struct {
	Server        string                 `json:"server"`
	Port          uint16                 `json:"port"`
	Username      string                 `json:"username"`
	Password      string                 `json:"password"`
	Channels      []string               `json:"channels"`
	Users         map[string]*userConfig `json:"users"`
	CommandPrefix []string               `json:"command_prefix"`
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
				"*": true,
			},
		},
	}
	c.CommandPrefix = []string{"+"}
}
