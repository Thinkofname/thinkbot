package spigot

import (
	"github.com/thinkofdeath/thinkbot"
)

// Init sets up the bot with spigot specific features
func Init(b *thinkbot.Bot) {
	initCommands(&b.Commands)
	initChat(b)
}
