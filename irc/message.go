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

import (
	"bytes"
	"errors"
	"strings"
)

// Message represents a single message from the server.
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
	// ErrEmptyMessage is returned if the message is empty
	ErrEmptyMessage = errors.New("empty message")
	// ErrMalformedMessage is returned if any error occurs when
	// parsing the message
	ErrMalformedMessage = errors.New("malformed message")
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

func (m unhandledMessage) Sender() string {
	return m.sender
}

func (m unhandledMessage) Command() string {
	return m.command
}

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
