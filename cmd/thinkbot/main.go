package main

import (
	"fmt"
	"github.com/thinkofdeath/thinkbot"
	"log"
	"strings"
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
				bot.JoinChannel("#thinkbot")
			case thinkbot.Stop:
				if bot.Error() != nil {
					log.Println(bot.Error())
					time.Sleep(5 * time.Second)
					continue
				}
			case thinkbot.ChannelMessage:
				if !event.CTCP {
					bot.SendMessage(event.Channel, fmt.Sprintf("\\o/ %s!", event.Sender.Nickname))
				}
			case thinkbot.PrivateMessage:
				if event.CTCP {
					switch {
					case strings.HasPrefix(event.Message, "VERSION"):
						bot.SendCTCP(event.Sender.Nickname, "VERSION Thinkbot v0.banana")
					}
				}
			default:
				log.Printf("%#v\n", event)
			}
		}
	}
}
