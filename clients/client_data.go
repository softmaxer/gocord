package clients

import "net"

type Client struct {
	ID   string
	Conn net.Conn
}

type Message struct {
	SenderID string
	Content  string
}
