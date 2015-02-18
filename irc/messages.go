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

package irc

// NewMessage creates a raw irc message from the passed
// parameters. No checking is done to make sure the message
// is valid.
func NewMessage(command string, args ...string) Message {
	return unhandledMessage{
		command:   command,
		arguments: args,
	}
}

// Notice is a message which should not be replied to
// via automatic responses
type Notice struct {
	unhandledMessage
}

// NewNotice creates a new Notice with the given target
// and text
func NewNotice(target, text string) Notice {
	return Notice{
		unhandledMessage{
			command:   "NOTICE",
			arguments: []string{target, text},
		},
	}
}

// Target returns the target of this notice
func (n Notice) Target() string {
	return n.arguments[0]
}

// Message returns the message of this notice
func (n Notice) Message() string {
	return n.arguments[1]
}

// Nick is a message used to set or change change
// your current nickname
type Nick struct {
	unhandledMessage
}

// NewNick creates a new Nick with the given nickname
func NewNick(nickname string) Nick {
	return Nick{
		unhandledMessage{
			command:   "NICK",
			arguments: []string{nickname},
		},
	}
}

// Nickname returns new/target nickname of this
// nick message
func (n Nick) Nickname() string {
	return n.arguments[0]
}

// User is a message used to register yourself to
// the server
type User struct {
	unhandledMessage
}

// NewUser creates a new User with the given username and realname
func NewUser(username, realname string) User {
	return User{
		unhandledMessage{
			command:   "USER",
			arguments: []string{username, "0", "*", realname},
		},
	}
}

// Username returns the username of the User message
func (u User) Username() string {
	return u.arguments[0]
}

// Realname returns the realname of the User message
func (u User) Realname() string {
	return u.arguments[1]
}

// Ping is sent by the server to check if the client is still
// there. It should be replied to by a Pong with the same code
type Ping struct {
	unhandledMessage
}

// NewPing creates a new Ping with the given code
func NewPing(code string) Ping {
	return Ping{
		unhandledMessage{
			command:   "PING",
			arguments: []string{code},
		},
	}
}

// Code returns the code of the ping
func (p Ping) Code() string {
	return p.arguments[0]
}

// Pong is sent as a reply to a Ping message
type Pong struct {
	unhandledMessage
}

// NewPong creates a new Ppng with the given code
func NewPong(code string) Pong {
	return Pong{
		unhandledMessage{
			command:   "PONG",
			arguments: []string{code},
		},
	}
}

// Code retruns the code of the pong
func (p Pong) Code() string {
	return p.arguments[0]
}

// PrivateMessage is a message sent by the server or
// by the client. This can target a user or a channel
type PrivateMessage struct {
	unhandledMessage
}

// NewPrivateMessage returns a private message with the
// given target and message
func NewPrivateMessage(target, msg string) PrivateMessage {
	return PrivateMessage{
		unhandledMessage{
			command:   "PRIVMSG",
			arguments: []string{target, msg},
		},
	}
}

// Target returns the target of the message
func (p PrivateMessage) Target() string {
	return p.arguments[0]
}

// Message returns the message
func (p PrivateMessage) Message() string {
	return p.arguments[1]
}

// Join is a message sent by the server to tell the client
// they joined a channel. A client can send this message
// to join a channel
type Join struct {
	unhandledMessage
}

// NewJoin creates a Join message for the passed channel
func NewJoin(channel string) Join {
	return Join{
		unhandledMessage{
			command:   "JOIN",
			arguments: []string{channel},
		},
	}
}

// Channel returns the channel this join is for
func (j Join) Channel() string {
	return j.arguments[0]
}

// Mode is a message sent by the server when a user's mode is
// changed, a channel's mode is changed or a flag on the channel
// changed. The client can send this to change a mode on a user
// or a channel
type Mode struct {
	unhandledMessage
}

// NewMode creates a Mode message for the passed target
// and mode string
func NewMode(target, mode string) Mode {
	return Mode{
		unhandledMessage{
			command:   "MODE",
			arguments: []string{target, mode},
		},
	}
}

// Target returns the target of this Mode message
func (m Mode) Target() string {
	return m.arguments[0]
}

// Mode returns the mode string of this mode Message
func (m Mode) Mode() string {
	return m.arguments[1]
}

// Extra returns the extra argument for this mode.
// (optional)
func (m Mode) Extra() string {
	return m.arguments[2]
}

// HasExtra returns whether this mode message has
// extra data
func (m Mode) HasExtra() bool {
	return len(m.arguments) >= 3
}
