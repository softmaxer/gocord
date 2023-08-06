package greeting

import (
	"chatserver/clients"
	"log"
)

func SayHello(client *clients.Client) {
	log.Println("New client: ", client.ID)
	welcomeMessage := []byte("Welcome to Go-cord!\n")
	client.Conn.Write(welcomeMessage)
}

func SayGoodbye(clientID string) {
	log.Printf("Client %s left the server\n", clientID)
}
