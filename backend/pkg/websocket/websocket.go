package websocket

import (
	"fmt"
	"io"
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
		fmt.Printf("%s\n", p)

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

// UPDATE CHAT to WEBSOCKET
func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return ws, err
	}
	return ws, nil
}

// WRTIE TO CHAT
func Writer(conn *websocket.Conn) {
	for {
		fmt.Println("sending")
		msgType, r, err := conn.NextReader()
		if err != nil {
			log.Fatal(err)
			return
		}
		w, err := conn.NextWriter(msgType)
		if err != nil {
			fmt.Println(err)
			return
		}
		if _, err := io.Copy(w, r); err != nil {
			fmt.Println(err)
			return
		}
		if err := w.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}
}
