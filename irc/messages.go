package irc

func NewMessage(command string, args ...string) Message {
	return unhandledMessage{
		command:   command,
		arguments: args,
	}
}

// An irc notice message
type Notice struct {
	unhandledMessage
}

// Creates a new Notice with the given target and text
func NewNotice(target, text string) Notice {
	return Notice{
		unhandledMessage{
			command:   "NOTICE",
			arguments: []string{target, text},
		},
	}
}

// Returns the target of this notice
func (n Notice) Target() string {
	return n.arguments[0]
}

// Returns the text of this notice
func (n Notice) Text() string {
	return n.arguments[1]
}

// An irc nick message
type Nick struct {
	unhandledMessage
}

// Creates a new Nick with the given nickname
func NewNick(nickname string) Nick {
	return Nick{
		unhandledMessage{
			command:   "NICK",
			arguments: []string{nickname},
		},
	}
}

// The nickname of this nick message
func (n Nick) Nickname() string {
	return n.arguments[0]
}

// An irc user message
type User struct {
	unhandledMessage
}

// Creates a new User with the given username and realname
func NewUser(username, realname string) User {
	return User{
		unhandledMessage{
			command:   "USER",
			arguments: []string{username, "0", "*", realname},
		},
	}
}

// The username of the User message
func (u User) Username() string {
	return u.arguments[0]
}

// The realname of the User message
func (u User) Realname() string {
	return u.arguments[1]
}

// An irc ping message
type Ping struct {
	unhandledMessage
}

// Creates a new Ping with the given code
func NewPing(code string) Ping {
	return Ping{
		unhandledMessage{
			command:   "PING",
			arguments: []string{code},
		},
	}
}

// The code of the ping
func (p Ping) Code() string {
	return p.arguments[0]
}

// An irc pong message
type Pong struct {
	unhandledMessage
}

// Creates a new Ppng with the given code
func NewPong(code string) Pong {
	return Pong{
		unhandledMessage{
			command:   "PONG",
			arguments: []string{code},
		},
	}
}

// The code of the ppng
func (p Pong) Code() string {
	return p.arguments[0]
}

// An irc privmsg message
type PrivateMessage struct {
	unhandledMessage
}

func NewPrivateMessage(target, msg string) PrivateMessage {
	return PrivateMessage{
		unhandledMessage{
			command:   "PRIVMSG",
			arguments: []string{target, msg},
		},
	}
}

// Returns the target of the message
func (p PrivateMessage) Target() string {
	return p.arguments[0]
}

// Returns the message
func (p PrivateMessage) Message() string {
	return p.arguments[1]
}

type Join struct {
	unhandledMessage
}

func NewJoin(channel string) Join {
	return Join{
		unhandledMessage{
			command:   "JOIN",
			arguments: []string{channel},
		},
	}
}

func (j Join) Channel() string {
	return j.arguments[0]
}

type Mode struct {
	unhandledMessage
}

func NewMode(target, mode string) Mode {
	return Mode{
		unhandledMessage{
			command:   "MODE",
			arguments: []string{target, mode},
		},
	}
}

func (m Mode) Target() string {
	return m.arguments[0]
}

func (m Mode) Mode() string {
	return m.arguments[1]
}

func (m Mode) User() string {
	return m.arguments[2]
}

func (m Mode) HasUser() bool {
	return len(m.arguments) >= 3
}
