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
	"github.com/thinkofdeath/thinkbot"
	"log"
	"sync"
	"time"
)

var (
	config     *botConfig
	configLock sync.RWMutex
)

func main() {
	config = loadConfig()
	saveConfig(config)

	for {
		log.Println("Connecting...")
		bot, err := thinkbot.NewBot(
			config.Server,
			config.Port,
			config.Username,
			func(b *thinkbot.BotConfig) {
				b.Password = config.Password
				initSpigotFeatures(b)
				b.SetPermissionContainer(configPermissions{})
				initCommands(b.Commands())
			},
		)
		if err != nil {
			log.Println(err)
			time.Sleep(5 * time.Second)
			continue
		}

		for event := range bot.Events {
			if processEvent(bot, event) {
				return
			}
		}
	}
}

func processEvent(bot *thinkbot.Bot, event thinkbot.Event) (end bool) {
	switch event := event.(type) {
	case thinkbot.Connected:
		log.Println("Connected")
		bot.AddMode('B')
		configLock.RLock()
		for _, c := range config.Channels {
			bot.JoinChannel(c)
		}
		configLock.RUnlock()
	case thinkbot.JoinChannel:
		joinChannel(event.Channel)
	case thinkbot.PartChannel:
		partChannel(event.Channel)
	case thinkbot.Stop:
		if bot.Error() != nil {
			log.Println(bot.Error())
			time.Sleep(5 * time.Second)
			return
		}
		return true
	default:
		log.Printf("Unhandled event: %#v\n", event)
	}
	return
}

func joinChannel(channel string) {
	configLock.Lock()
	defer configLock.Unlock()
	for _, c := range config.Channels {
		if c == channel {
			return
		}
	}
	config.Channels = append(config.Channels, channel)
	saveConfig(config)
}

func partChannel(channel string) {
	configLock.Lock()
	defer configLock.Unlock()
	for i, c := range config.Channels {
		if c == channel {
			config.Channels = append(config.Channels[:i], config.Channels[i+1:]...)
			saveConfig(config)
			return
		}
	}
}
