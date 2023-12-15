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
	"sync"
	"time"
)

type client struct {
	Username string
	password string
	conn     *net.Conn
	writer   *textproto.Writer
	reader   *textproto.Reader
	mu       sync.RWMutex
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

var channels = []string{"#software-development", "##software-development"}

var ErrBadConnect = errors.New("unable to connect")

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
	c.reader = textproto.NewReader(r)
	// Wait for the server to send 4 lines
	for i := 0; i < 4; i++ {
		if _, err := c.reader.ReadLine(); err != nil {
			slog.Warn("waiting for IRC server to send lines", "value", err)
		}
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	// Start sending login information
	if err := c.writer.PrintfLine("USER %s 8 * :%s", c.Username, c.Username); err != nil {
		slog.Warn("trouble writing username to IRC", "value", err)

		return fmt.Errorf("%w sending username", ErrBadConnect)
	}

	if err := c.writer.PrintfLine("NICK %s", c.Username); err != nil {
		slog.Warn("trouble writing nick to IRC", "value", err)

		return fmt.Errorf("%w setting nick", ErrBadConnect)
	}

	message := fmt.Sprintf("PRIVMSG NickServ :identify %s %s", c.Username, c.password)
	if err := c.writer.PrintfLine(message); err != nil {
		slog.Warn("trouble identifying to IRC server", "value", err)

		return fmt.Errorf("%w sending identify", ErrBadConnect)
	}

	// Join channel
	for _, channel := range channels {
		message = fmt.Sprintf("JOIN %s", channel)
		if err := c.writer.PrintfLine(message); err != nil {
			slog.Warn("trouble joining channel", "value", err)

			return fmt.Errorf("%w joining channel", ErrBadConnect)
		}
	}

	return nil
}

func (c *client) watch() {
	for {
		data, err := c.reader.ReadLine()
		if err != nil {
			slog.Warn("reading from server", "value", err)
		}

		switch {
		case strings.HasPrefix(data, "PING"):
			go c.keepalive(data)
		default:
			continue
		}
	}
}

func (c *client) keepalive(ping string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if err := c.writer.PrintfLine(strings.Replace(ping, "PING", "PONG", 1)); err != nil {
		slog.Warn("trouble sending PONG to IRC server", "value", err)
	}
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

		go c.watch()
	}

	for _, channel := range channels {
		// Send release information
		message := fmt.Sprintf("PRIVMSG %s :Blog Announcement", channel)

		// Take write lock -
		// do NOT defer this lock, because of the loop that will cause a deadlock.
		c.mu.Lock()

		if err := c.writer.PrintfLine(message); err != nil {
			slog.Warn("trouble writing blog announcement", "value", err)

			c.mu.Unlock()

			goto RETRY
		}

		//nolint:gomnd
		time.Sleep(time.Millisecond * 500)

		message = fmt.Sprintf("PRIVMSG %s :Link to blog: %s", channel, content["link"])
		if err := c.writer.PrintfLine(message); err != nil {
			slog.Warn("trouble writing further information", "value", err)

			c.mu.Unlock()

			goto RETRY
		}
		// Wait for the server - this is not good, but not getting a clear view on
		// why the server won't print out what's sent if I disconnect before this
		// time amount.
		//nolint:gomnd
		time.Sleep(time.Second * 7)

		// Unlock the mutex
		c.mu.Unlock()
	}

	return nil
}

func (c *client) disconnect() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if err := c.writer.PrintfLine("QUIT"); err != nil {
		slog.Warn("trouble quitting from IRC", "value", err)
	}

	if err := (*c.conn).Close(); err != nil {
		slog.Warn("trouble closing IRC connection", "value", err)
	}
}
