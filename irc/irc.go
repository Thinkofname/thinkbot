package irc

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	conn    *net.TCPConn
	scanner *bufio.Scanner
}

// Creates a new irc client connecting to the server
// at the passed address and port. This does not use
// ssl.
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

// Reads a single message from the client
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
	case "ping":
		msg = Ping{m}
	case "pong":
		msg = Pong{m}
	case "join":
		msg = Join{m}
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
