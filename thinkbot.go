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

package thinkbot

import (
	"fmt"
	"github.com/thinkofdeath/thinkbot/command"
	"github.com/thinkofdeath/thinkbot/irc"
	"log"
	"regexp"
	"strings"
)

// Bot contains all the information of the currently
// running bot.
//
// Event should be continuously read from and the
// returned events handled.
type Bot struct {
	client   *irc.Client
	err      error
	username string
	password string

	commands command.Registry

	chatHandlers []chatHandler

	Events    chan Event
	writeChan chan irc.Message
	funcChan  chan func()

	channels      []string
	commandPrefix []string
	modes         map[rune]struct{}

	permissionContainer PermissionContainer
}

type chatHandlerFunc func(b *Bot, sender User, target, message string) error

type chatHandler struct {
	r *regexp.Regexp
	f chatHandlerFunc
}

// NewBot creates an instance of a bot connecting to the target
// server with the specified username
//
// The init function is called before any messages/commands
// are handled to allow for setup
//
// Returns an error if it fails to connect
func NewBot(server string, port uint16, username string, init func(*BotConfig)) (*Bot, error) {
	c, err := irc.NewClient(server, port)
	if err != nil {
		return nil, err
	}
	b := &Bot{
		client:    c,
		Events:    make(chan Event, 100),
		writeChan: make(chan irc.Message, 100),
		funcChan:  make(chan func(), 100),
		username:  username,
		channels:  []string{},
		modes:     map[rune]struct{}{},
		commands: command.Registry{
			// User and target parameters
			ExtraParameters: 2,
		},
	}
	config := &BotConfig{
		bot:      b,
		Prefixes: []string{"+"},
	}
	init(config)

	b.password = config.Password
	b.commandPrefix = config.Prefixes

	go b.run()
	return b, nil
}

func (b *Bot) run() {
	c := b.client
	rc := make(chan irc.Message)

	go func() {
		for {
			m, err := c.Read()
			if err != nil {
				b.err = err
				b.kill()
				close(rc)
				return
			}
			rc <- m
		}
	}()

	c.Write(irc.NewNick(b.username))
	c.Write(irc.NewUser(b.username, b.username))
	if b.password != "" {
		c.Write(irc.NewPass(b.password))
	}

	for {
		select {
		case f := <-b.funcChan:
			f()
		case m := <-b.writeChan:
			c.Write(m)
		case m, ok := <-rc:
			if !ok {
				return
			}
			b.handleIRCMessage(m)
		}
	}
}

func (b *Bot) handleIRCMessage(m irc.Message) {
	c := b.client
	switch m := m.(type) {
	case irc.Reply:
		b.handleReply(m)
	case irc.Ping:
		c.Write(irc.NewPong(m.Code()))
	case irc.Join:
		for _, c := range b.channels {
			if c == m.Channel() {
				return
			}
		}
		b.channels = append(b.channels, m.Channel())
		b.Events <- JoinChannel{Channel: m.Channel()}
	case irc.Part:
		for i, c := range b.channels {
			if c == m.Channel() {
				b.channels = append(b.channels[:i], b.channels[i+1:]...)
				b.Events <- PartChannel{Channel: m.Channel()}
				return
			}
		}
	case irc.PrivateMessage:
		msg := m.Message()
		ctcp := len(msg) > 2 && msg[0] == '\x01'
		if ctcp {
			msg = msg[1 : len(msg)-1]
		}
		for _, prefix := range b.commandPrefix {
			if strings.HasPrefix(msg, prefix) {
				go b.handleCommand(
					parseUser(m.Sender()),
					m.Target(),
					m.Message()[len(prefix):],
				)
				return
			}
		}

		b.handleMessage(parseUser(m.Sender()), m.Target(), m.Message())
	case irc.Notice:
	// Ignored
	case irc.Mode:
		// TODO Track others + channels
		if m.Target() == b.username {
			state := '#'
			for _, r := range m.Mode() {
				switch r {
				case '-':
					state = '-'
				case '+':
					state = '+'
				default:
					if state == '#' {
						panic("Invalid mode!")
					}
					if state == '+' {
						b.modes[r] = struct{}{}
					} else {
						delete(b.modes, r)
					}
				}
			}
		}
	default:
		log.Printf("Unhandled: %#v\n", m)
	}
}

// AddCommandPrefix adds a prefix that the bot will
// recognise as a command
func (b *Bot) AddCommandPrefix(pre string) {
	b.funcChan <- func() {
		for _, p := range b.commandPrefix {
			if p == pre {
				return
			}
		}
		b.commandPrefix = append(b.commandPrefix, pre)
	}
}

// RemoveCommandPrefix removes a prefix that the bot will
// recognise as a command
func (b *Bot) RemoveCommandPrefix(pre string) {
	b.funcChan <- func() {
		for i, p := range b.commandPrefix {
			if p == pre {
				b.commandPrefix = append(b.commandPrefix[:i], b.commandPrefix[i+1:]...)
				return
			}
		}
	}
}

// AddChatHandler adds a handler which is called
// whenever the passed regexp matches a message
func (b *Bot) AddChatHandler(r *regexp.Regexp, f chatHandlerFunc) {
	b.funcChan <- func() {
		b.chatHandlers = append(b.chatHandlers, chatHandler{r, f})
	}
}

// RemoveChatHandler removes a handler which is called
// whenever the passed regexp matches a message
func (b *Bot) RemoveChatHandler(r *regexp.Regexp) {
	b.funcChan <- func() {
		re := r.String()
		for i, c := range b.chatHandlers {
			if c.r.String() == re {
				b.chatHandlers = append(b.chatHandlers[:i], b.chatHandlers[i+1:]...)
				return
			}
		}
	}
}

func (b *Bot) handleMessage(sender User, target, message string) {
	for _, h := range b.chatHandlers {
		r, f := h.r, h.f
		if r.MatchString(message) {
			go func() {
				if err := f(b, sender, target, message); err != nil {
					b.Reply(sender, target, err.Error())
				}
			}()
			break
		}
	}
}

func (b *Bot) handleCommand(user User, target, msg string) {
	err := b.commands.Execute(b, msg, user, target)
	if err != nil && err != command.ErrCommandNotFound {
		b.Reply(user, target, err.Error())
	}
}

func (b *Bot) handleReply(r irc.Reply) {
	switch r.Code() {
	case irc.ReplyWelcome:
		b.Events <- Connected{}
	}
}

func (b *Bot) kill() {
	b.Events <- Stop{}
	close(b.Events)
}

// Error returns the error the bot stopped with or nil
// if no error has occurred
func (b *Bot) Error() error {
	return b.err
}

// JoinChannel attempts to join the target channel
//
// Channel names are generally prefixed with #
func (b *Bot) JoinChannel(channel string) {
	b.writeChan <- irc.NewJoin(channel)
}

// PartChannel attempts to part the target channel
//
// Channel names are generally prefixed with #
func (b *Bot) PartChannel(channel string) {
	b.writeChan <- irc.NewPart(channel)
}

// Channels returns the list of channels this bot is currently
// connected to
func (b *Bot) Channels() []string {
	ret := make(chan []string)
	b.funcChan <- func() {
		ret <- b.channels
	}
	return <-ret
}

// SendMessage sends a message to the target channel or user
func (b *Bot) SendMessage(target, message string) {
	b.writeChan <- irc.NewPrivateMessage(target, message)
}

// SendCTCP sends a CTCP message to the target channel or user
func (b *Bot) SendCTCP(target, message string) {
	b.writeChan <- irc.NewPrivateMessage(target, "\x01"+message+"\x01")
}

// AddMode sets the mode(s) on the bot
func (b *Bot) AddMode(modes ...rune) {
	b.writeChan <- irc.NewMode(b.username, "+"+string(modes))
}

// RemoveMode removes the mode(s) from the bot
func (b *Bot) RemoveMode(modes ...rune) {
	b.writeChan <- irc.NewMode(b.username, "-"+string(modes))
}

// Reply sends a message to the user in the same way
// the message was sent.
//
// If the message was sent in a channel the message
// will be sent back to that channel with the sender's
// nickname prefixed.
//
// If the message was sent in a private message then
// this will just reply normally
func (b *Bot) Reply(sender User, target, message string) {
	if target[0] == '#' {
		b.SendMessage(target, fmt.Sprintf("%s: %s", sender.Nickname, message))
	} else {
		b.SendMessage(sender.Nickname, message)
	}
}
