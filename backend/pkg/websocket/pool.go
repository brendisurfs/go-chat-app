package websocket

import "fmt"

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

// NewPool func creates a new pool of channels.
func NewPool() *Pool {
	return &Pool{
		Register: make(chan *Client),
		// Unregister a user and notify the pool when a client disconnects.
		Unregister: make(chan *Client),
		// a map of clients to a boolean value. bool = active inactive, but not disconnected.
		Clients: make(map[*Client]bool),
		// a channel, when a message is passed to it, will loop through all clients in the pool and send the message through the socket.
		Broadcast: make(chan Message),
	}
}

// fancy pool controllers and what to do with each message.
func (pool *Pool) Start() {
	for {
		select {
		// CASE 1
		// |
		// v
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("size of the connection pool: ", len(pool.Clients))
			// NOTE: what is this doing exactly??
			for client := range pool.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined!"})
			}

			//CASE 2
			//	|
			//	v
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("size of connection pool: ", len(pool.Clients))
			for client := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "User disconnected"})
			}
		//CASE 3
		//	|
		//	v
		case message := <-pool.Broadcast:
			fmt.Println("sending message to all clients in pool.")
			for client := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
