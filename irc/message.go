package irc

import (
	"bytes"
	"errors"
	"strings"
)

type Message interface {
	Sender() string
	Command() string
	Arguments() []string

	asBytes() []byte
}

type unhandledMessage struct {
	sender    string
	command   string
	arguments []string
}

var (
	ErrEmptyMessage     = errors.New("Empty message")
	ErrMalformedMessage = errors.New("Malformed message")
)

func parseMessage(line string) (msg unhandledMessage, err error) {
	if len(line) == 0 {
		err = ErrEmptyMessage
		return
	}
	if line[0] == ':' {
		pos := strings.IndexRune(line, ' ')
		if pos == -1 {
			err = ErrMalformedMessage
			return
		}
		msg.sender = line[1:pos]
		line = line[pos+1:]
	}

	pos := strings.IndexRune(line, ' ')
	if pos == -1 {
		msg.command = line
		return
	}

	msg.command = line[:pos]
	line = line[pos+1:]

	msg.arguments = []string{}

	for len(line) > 0 {
		// Allow for spaces if the argument starts with ':'
		if line[0] == ':' {
			msg.arguments = append(msg.arguments, line[1:])
			break
		}

		var arg string
		if pos := strings.IndexRune(line, ' '); pos == -1 {
			arg = line
			line = line[len(arg):]
		} else {
			arg = line[:pos]
			line = line[pos+1:]
		}

		msg.arguments = append(msg.arguments, arg)

	}
	return
}

// The client/server that sent the message
func (m unhandledMessage) Sender() string {
	return m.sender
}

// The message's command
func (m unhandledMessage) Command() string {
	return m.command
}

// The arguments (if any) for this message
func (m unhandledMessage) Arguments() []string {
	return m.arguments
}

func (m unhandledMessage) asBytes() []byte {
	var buf bytes.Buffer

	if m.sender != "" {
		buf.WriteRune(':')
		buf.WriteString(m.sender)
		buf.WriteRune(' ')
	}

	buf.WriteString(m.command)
	buf.WriteRune(' ')

	for _, a := range m.arguments {
		if strings.ContainsRune(a, ' ') {
			buf.WriteRune(':')
		}
		buf.WriteString(a)
		buf.WriteRune(' ')
	}

	return buf.Bytes()[:buf.Len()-1]
}
