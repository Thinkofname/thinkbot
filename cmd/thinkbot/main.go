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
