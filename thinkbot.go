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
	"github.com/thinkofdeath/thinkbot/irc"
	"log"
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

	Events    chan Event
	writeChan chan irc.Message
	funcChan  chan func()

	channels      []string
	commandPrefix string
	modes         map[rune]struct{}
}

// NewBot creates an instance of a bot connecting to the target
// server with the specified username
//
// Returns an error if it fails to connect
func NewBot(server string, port uint16, username string) (*Bot, error) {
	c, err := irc.NewClient(server, port)
	if err != nil {
		return nil, err
	}
	b := &Bot{
		client:        c,
		Events:        make(chan Event, 100),
		writeChan:     make(chan irc.Message, 100),
		funcChan:      make(chan func(), 100),
		username:      username,
		channels:      []string{},
		commandPrefix: "+",
		modes:         map[rune]struct{}{},
	}
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
			switch m := m.(type) {
			case irc.Reply:
				b.handleReply(m)
			case irc.Ping:
				c.Write(irc.NewPong(m.Code()))
			case irc.Join:
				b.channels = append(b.channels, m.Channel())
			case irc.PrivateMessage:
				msg := m.Message()
				ctcp := len(msg) > 2 && msg[0] == '\x01'
				if ctcp {
					msg = msg[1 : len(msg)-1]
				}
				isCommand := strings.HasPrefix(msg, b.commandPrefix)

				if !isCommand {
					b.handleMessage(parseUser(m.Sender()), m.Target(), m.Message())
				} else {
					go b.handleCommand(
						parseUser(m.Sender()),
						m.Target(),
						m.Message()[len(b.commandPrefix):],
					)
				}
			case irc.Notice:
			// Notice
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
