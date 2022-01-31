package worker

import (
	"dataplane/logging"
	"dataplane/messageq"
	"log"
	"os"

	"github.com/gofiber/websocket/v2"
	"github.com/nats-io/nats.go"
)

var Broadcast = make(chan []byte)
var messagereceive = make(chan []byte)

// https://github.com/gorilla/websocket/blob/master/examples/chat/client.go

// https://github.com/marcelo-tm/testws/blob/master/main.go
func WorkerStatsWs(conn *websocket.Conn, subject string) {

	// Subscribe to a specific worker group when the connection is open
	sub, _ := messageq.NATSencoded.Subscribe(subject, func(m *nats.Msg) {

		broadcast <- m.Data

	})

	// When the function returns, unregister the client and close the connection
	defer func() {
		unregister <- conn
		conn.Close()
		sub.Unsubscribe()
	}()

	// Register the client
	register <- conn

	for {

		mt, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("read error:", err)
			}
			return
		}

		if os.Getenv("messagedebug") == "true" {
			logging.PrintSecretsRedact("message received from client:", mt, string(message))
		}

	}

}