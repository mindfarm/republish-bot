// Package irc -
package irc

import (
	"bufio"
	"fmt"
	"log/slog"
	"net"
	"net/textproto"
	"strings"
	"time"
)

type client struct {
	Username string
	password string
	conn     *net.Conn
	writer   *textproto.Writer
}

// NewFreenodeClient -
//
//nolint:golint,revive
func NewFreenodeClient(username, password string) (*client, error) {
	if username == "" {
		return nil, fmt.Errorf("cannot create a new client, missing username")
	}
	if password == "" {
		return nil, fmt.Errorf("cannot create a new client, missing password")
	}

	c := &client{
		Username: username,
		password: password,
	}
	return c, nil
}

// chat.freenode.net on ports 6665-6667 and 8000-8002 for plain-text connections, or ports 6697, 7000 and 7070 for TLS-encrypted connections.
const freenode = "chat.freenode.net:6665"
const channel = "#go-nuts"

func (c *client) connect() error {
	conn, err := net.Dial("tcp", freenode)
	if err != nil {
		slog.Error("Unable to connect to IRC server with error", err)
		return fmt.Errorf("unable to connect to %s with error %w", freenode, err)
	}
	c.conn = &conn
	w := bufio.NewWriter(*c.conn)
	r := bufio.NewReader(*c.conn)
	c.writer = textproto.NewWriter(w)
	reader := textproto.NewReader(r)
	// Wait for the server to send 4 lines
	for i := 0; i < 4; i++ {
		reader.ReadLine()
	}
	// Start sending login information
	c.writer.PrintfLine("USER %s 8 * :%s", c.Username, c.Username)
	c.writer.PrintfLine("NICK %s", c.Username)
	s := fmt.Sprintf("PRIVMSG NickServ :identify %s %s", c.Username, c.password)
	c.writer.PrintfLine(s)

	// Join channel
	s = fmt.Sprintf("JOIN %s", channel)
	c.writer.PrintfLine(s)

	return nil
}

func (c *client) PublishContent(content map[string]string) error {
	if c.conn == nil {
		err := c.connect()
		if err != nil {
			return fmt.Errorf("cannot chat with error %w", err)
		}
	}
	title := content["title"]
	// Go project has a weird title structure
	// [release-branch.go1.15] go1.15.2
	tmp := strings.Split(title, "]")
	if len(tmp) > 1 {
		title = strings.TrimSpace(tmp[1])
	}
	// Send release information
	s := fmt.Sprintf("PRIVMSG %s :Release Announcement", channel)
	s += fmt.Sprintf(" %s is now available.", title)
	c.writer.PrintfLine(s)
	time.Sleep(time.Millisecond * 500)
	s = fmt.Sprintf("PRIVMSG %s :Further information can be found at %s", channel, content["link"])
	c.writer.PrintfLine(s)
	// Wait for the server - this is not good, but not getting a clear view on
	// why the server won't print out what's sent if I disconnect before this
	// time amount.
	time.Sleep(time.Second * 7)
	return c.disconnect()
}

func (c *client) disconnect() error {
	c.writer.PrintfLine("QUIT")
	(*c.conn).Close()
	return nil
}
