// +build spigot

package main

import (
	"github.com/thinkofdeath/thinkbot"
	"github.com/thinkofdeath/thinkbot/spigot"
)

func initSpigotFeatures(b *thinkbot.BotConfig) {
	spigot.Init(b)
}
