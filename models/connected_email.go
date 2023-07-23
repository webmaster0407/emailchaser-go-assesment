package models

import (
	"fmt"

	"github.com/knadh/go-pop3"
)

type ConnectedEmail struct {
	ID           int
	Email        string
	Pop3Host     string
	Pop3Port     int
	Pop3Username string
	Pop3Password string
}

func (c *ConnectedEmail) connectAndAuth() (*pop3.Conn, error) {
	// Create a new POP3 client using the host and port specified in the ConnectedEmail struct.
	client := pop3.New(pop3.Opt{
		Host:       c.Pop3Host,
		Port:       c.Pop3Port,
		TLSEnabled: true,
	})

	// Try to establish a connection to the POP3 server.
	connectPop3, err := client.NewConn()
	if err != nil {
		return nil, err
	}

	// Ensure the Quit method is called before the function exits.
	defer connectPop3.Quit()

	// Authenticate to the POP3 server using the username and password specified in the ConnectedEmail struct.
	if err := connectPop3.Auth(
		c.Pop3Username,
		c.Pop3Password,
	); err != nil {
		return nil, err
	}

	return connectPop3, nil
}

func (c *ConnectedEmail) AuthPop3() error {
	_, err := c.connectAndAuth()
	return err
}

func (c *ConnectedEmail) ReadEmailsPop3() error {
	// Establish a connection and authenticate with the POP3 server.
	connectPop3, err := c.connectAndAuth()
	if err != nil {
		return err
	}

	// Get count and size of all messages.
	count, _, err := connectPop3.Stat()
	if err != nil {
		return err
	}

	// Pull all messages on the server. Message IDs go from 1 to N.
	for id := 1; id <= count; id++ {
		m, _ := connectPop3.Retr(id)
		fmt.Println(id, "=", m.Header.Get("subject"))
	}

	return nil
}
