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
	"github.com/thinkofdeath/thinkbot/spigot"
	"log"
	"time"
)

func main() {
	config := loadConfig()
	saveConfig(config)

	for {
		log.Println("Connecting...")
		bot, err := thinkbot.NewBot(
			config.Server,
			config.Port,
			config.Username,
			func(b *thinkbot.Bot) {
				spigot.Init(b)
			},
		)
		if err != nil {
			log.Println(err)
			time.Sleep(5 * time.Second)
			continue
		}

		for event := range bot.Events {
			switch event := event.(type) {
			case thinkbot.Connected:
				log.Println("Connected")
				bot.AddMode('B')
				for _, c := range config.Channels {
					bot.JoinChannel(c)
				}
			case thinkbot.Stop:
				if bot.Error() != nil {
					log.Println(bot.Error())
					time.Sleep(5 * time.Second)
					continue
				}
				return
			default:
				log.Printf("Unhandled event: %#v\n", event)
			}
		}
	}
}
