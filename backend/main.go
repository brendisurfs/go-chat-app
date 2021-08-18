package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/brendisurfs/go-chatapp/pkg/websocket"
)

// WEBSOCKET ENDPOINT
func ServeWS(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)
	// NOW we upgrade the connection from http to WS
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Fatal(err)
	}
	// now we will listen for our messages, and send them to the reader func.
	go websocket.Writer(ws)
	websocket.Reader(ws)
}

/*



 */
// ROUTES
func setupRoutes() {
	// mapping our "/ws" end point to the ServeWS func.
	http.HandleFunc("/ws", ServeWS)
}

func main() {
	fmt.Println("chat app v0.0.1")
	setupRoutes()
	fmt.Println("Houston, we have liftoff.")
	http.ListenAndServe(":8080", nil)
}
