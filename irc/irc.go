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
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

// Client represents a connection to an irc server
// as a client
type Client struct {
	conn    *net.TCPConn
	scanner *bufio.Scanner
}

// NewClient creates a new irc client connecting to
// the server at the passed address and port. This
// does not use ssl.
func NewClient(address string, port uint16) (cli *Client, err error) {
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		return
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		return
	}

	cli = &Client{
		conn:    conn,
		scanner: bufio.NewScanner(conn),
	}
	return
}

// Read reads a single message from the server
func (c *Client) Read() (msg Message, err error) {
	c.conn.SetReadDeadline(time.Now().Add(time.Second * 240))

	if !c.scanner.Scan() {
		err = c.scanner.Err()
		return
	}

	m, err := parseMessage(c.scanner.Text())

	switch strings.ToLower(m.command) {
	case "notice":
		msg = Notice{m}
	case "nick":
		msg = Nick{m}
	case "user":
		msg = User{m}
	case "pass":
		msg = Pass{m}
	case "ping":
		msg = Ping{m}
	case "pong":
		msg = Pong{m}
	case "join":
		msg = Join{m}
	case "part":
		msg = Part{m}
	case "privmsg":
		msg = PrivateMessage{m}
	case "mode":
		msg = Mode{m}
	default:
		if m.command[0] >= '0' && m.command[0] <= '9' {
			msg = Reply{m}
		} else {
			msg = m
		}
	}
	return
}

var newLine = []byte("\n")

// Write writes a single message to the server as
// the client
func (c *Client) Write(msg Message) (err error) {
	c.conn.SetWriteDeadline(time.Now().Add(time.Second * 240))

	b := msg.asBytes()

	_, err = c.conn.Write(b)
	if err != nil {
		return
	}
	_, err = c.conn.Write(newLine)
	if err != nil {
		return
	}
	return
}
