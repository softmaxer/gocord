package server

import (
	"chatserver/clients"
	"chatserver/greeting"
)

func clearConnection(connectionsTable map[string]*clients.Client, clientID string) {
	delete(connectionsTable, clientID)
	greeting.SayGoodbye(clientID)
}

func ReadMessage(client *clients.Client, channel chan clients.Message, connectionsTable map[string]*clients.Client) {
	conn := client.Conn
	defer clearConnection(connectionsTable, client.ID)
	for {
		buffer := make([]byte, 2048)
		nbytes, err := conn.Read(buffer)
		if err != nil {
			break
		}
		message := string(buffer[:nbytes])
		channel <- clients.Message{SenderID: client.ID, Content: message}
	}
}

func BroadcastMessage(channel chan clients.Message, connectionsTable map[string]*clients.Client) {
	for {
		select {
		case message := <-channel:
			for _, connection := range connectionsTable {
				if connection.ID == message.SenderID {
					continue
				}
				senderPrompt := "\n" + message.SenderID + ": "
				connection.Conn.Write([]byte(senderPrompt))
				connection.Conn.Write([]byte(message.Content))
				connection.Conn.Write([]byte("\nME: "))
			}
		}
	}
}
