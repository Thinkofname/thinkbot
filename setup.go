package thinkbot

import (
	"github.com/thinkofdeath/thinkbot/command"
	"regexp"
)

// BotConfig allows for initial setup of a bot during
// init
type BotConfig struct {
	bot *Bot
}

// SetPermissionContainer sets the container used for checking
// permissions for irc users.
func (b *BotConfig) SetPermissionContainer(pc PermissionContainer) {
	b.bot.setPermissionContainer(pc)
}

// Commands returns the command registry for this bot.
// It is not safe to retain the result of this method after
// init
func (b *BotConfig) Commands() *command.Registry {
	return &b.bot.commands
}

// AddChatHandler adds a handler which is called
// whenever the passed regexp matches a message
func (b *BotConfig) AddChatHandler(r *regexp.Regexp, f chatHandlerFunc) {
	b.bot.chatHandlers = append(b.bot.chatHandlers, chatHandler{r, f})
}
