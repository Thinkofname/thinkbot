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
	"time"
)

func main() {
	for {
		log.Println("Connecting...")
		bot, err := thinkbot.NewBot("irc.spi.gt", 6667, "ThinkTest")
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
				bot.JoinChannel("#thinkbot")
			case thinkbot.Stop:
				if bot.Error() != nil {
					log.Println(bot.Error())
					time.Sleep(5 * time.Second)
					continue
				}
			default:
				log.Printf("Unhandled event: %#v\n", event)
			}
		}
	}
}
