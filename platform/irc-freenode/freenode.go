// Package irc -
package irc

import (
	"bufio"
	"errors"
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

var ErrClientCreation = errors.New("cannot create a new client")

func NewFreenodeClient(username, password string) (*client, error) {
	if username == "" {
		return nil, fmt.Errorf("%w, missing username", ErrClientCreation)
	}

	if password == "" {
		return nil, fmt.Errorf("%w, missing password", ErrClientCreation)
	}

	c := &client{
		Username: username,
		password: password,
	}

	return c, nil
}

// chat.freenode.net on ports 6666-6667 and 8000-8002 for plain-text
// connections, or ports 6697, 7000 and 7070 for TLS-encrypted connections.
const freenode = "irc.libera.chat:6665"
const channel = "#software-development"

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
		if _, err := reader.ReadLine(); err != nil {
			slog.Warn("waiting for IRC server to send lines", "value", err)
		}
	}

	// Start sending login information
	if err := c.writer.PrintfLine("USER %s 8 * :%s", c.Username, c.Username); err != nil {
		slog.Warn("trouble writing username to IRC", "value", err)
	}

	if err := c.writer.PrintfLine("NICK %s", c.Username); err != nil {
		slog.Warn("trouble writing nick to IRC", "value", err)
	}

	message := fmt.Sprintf("PRIVMSG NickServ :identify %s %s", c.Username, c.password)
	if err := c.writer.PrintfLine(message); err != nil {
		slog.Warn("trouble identifying to IRC server", "value", err)
	}

	// Join channel
	message = fmt.Sprintf("JOIN %s", channel)
	if err := c.writer.PrintfLine(message); err != nil {
		slog.Warn("trouble joining channel", "value", err)
	}

	return nil
}

const maxRetries = 5

var ErrIRCConnection = errors.New("cannot connect to IRC")

func (c *client) PublishContent(content map[string]string) error {
	var retries int
RETRY:
	retries++

	//nolint:gomnd
	if retries > 3 {
		slog.Warn("disconnecting")
		c.disconnect()
		time.Sleep(5 * time.Second)

		c.conn = nil
	}

	if retries >= maxRetries {
		slog.Warn("exceeded retries")

		return ErrIRCConnection
	}

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
	message := fmt.Sprintf("PRIVMSG %s :Blog Announcement", channel)
	message += fmt.Sprintf(" %s is now available.", title)

	if err := c.writer.PrintfLine(message); err != nil {
		slog.Warn("trouble writing blog announcement", "value", err)

		goto RETRY
	}

	//nolint:gomnd
	time.Sleep(time.Millisecond * 500)

	message = fmt.Sprintf("PRIVMSG %s :Further information can be found at %s", channel, content["link"])
	if err := c.writer.PrintfLine(message); err != nil {
		slog.Warn("trouble writing further information", "value", err)

		goto RETRY
	}
	// Wait for the server - this is not good, but not getting a clear view on
	// why the server won't print out what's sent if I disconnect before this
	// time amount.
	//nolint:gomnd
	time.Sleep(time.Second * 7)

	return nil
}

func (c *client) disconnect() {
	if err := c.writer.PrintfLine("QUIT"); err != nil {
		slog.Warn("trouble quitting from IRC", "value", err)
	}

	if err := (*c.conn).Close(); err != nil {
		slog.Warn("trouble closing IRC connection", "value", err)
	}
}
