package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/brendisurfs/go-chatapp/pkg/websocket"
)

// WEBSOCKET ENDPOINT
func ServeWS(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)
	// NOW we upgrade the connection from http to WS
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Fatal(err)
	}
	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}
	// send the client to our pool handler
	pool.Register <- client
	client.Read()
}

/*



 */
// ROUTES
func setupRoutes() {
	pool := websocket.NewPool()
	// create a new goroutine for hitting our route.
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWS(pool, w, r)
	})
}

func main() {
	fmt.Println("chat app v0.0.1")
	setupRoutes()
	fmt.Println("Houston, we have liftoff.")
	http.ListenAndServe(":8080", nil)
}
