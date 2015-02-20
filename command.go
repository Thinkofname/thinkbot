package thinkbot

func (b *Bot) init() {
	b.Commands.Register("join %", join)
}

func join(b *Bot, sender User, target, channel string) {
	if len(channel) < 0 || channel[0] != '#' {
		panic("Invalid channel")
	}
	b.JoinChannel(channel)
}
