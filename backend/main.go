package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Need to upgrade to a websocket rather than an http connection. 101 protocol.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// |
	// |
	// v
	// check the origin, make sure we are good.
	CheckOrigin: func(r *http.Request) bool { return true },
}

// define a ws reader to listen for messages.
// for Brendans notes: remember we use a pointer to the struct, because we dont want to make copies of it, just read its data.
func Reader(conn *websocket.Conn) {
	for {
		// read message:
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// NOTE: QUICK PRINT CHECK HERE
		fmt.Printf("%s", p)

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

/*



 */
// WEBSOCKET ENDPOINT
func ServeWS(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)
	// NOW we upgrade the connection from http to WS
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// now we will listen for our messages, and send them to the reader func.
	Reader(ws)
}

/*



 */
// ROUTES
func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})
	// mapping our "/ws" end point to the ServeWS func.
	http.HandleFunc("/ws", ServeWS)
}

func main() {
	fmt.Println("chat app v0.0.1")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
