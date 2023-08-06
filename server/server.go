package server

import (
	"chatserver/clients"
	"chatserver/greeting"
	"log"
	"net"

	"github.com/google/uuid"
)

func InitServer(url string) {
	// Initialize listener
	ln, err := net.Listen("tcp", url)
	if err != nil {
		log.Fatalln("Something went wrong: ", err)
	}
	defer ln.Close()

	// Declare a new text channel
	textChannel := make(chan clients.Message)

	// Declare a new connections table
	connectionsTable := make(map[string]*clients.Client)

	// Start the loop
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln("Connection couldn't be established: ", err)
		}
		clientID := uuid.New().String()
		client := &clients.Client{ID: clientID, Conn: conn}
		connectionsTable[clientID] = client
		greeting.SayHello(client)
		conn.Write([]byte("ME: "))
		go ReadMessage(client, textChannel, connectionsTable)
		go BroadcastMessage(textChannel, connectionsTable)
		defer conn.Close()
	}
}
