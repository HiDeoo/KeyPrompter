package net

type Pool struct {
	Broadcast  chan []byte
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
}

func newPool() *Pool {
	return &Pool{
		Broadcast:  make(chan []byte),
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (pool *Pool) run() {
	for {
		select {
		case client := <-pool.register:
			pool.clients[client] = true
		case client := <-pool.unregister:
			if _, ok := pool.clients[client]; ok {
				delete(pool.clients, client)
				close(client.send)
			}
		case message := <-pool.Broadcast:
			for client := range pool.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(pool.clients, client)
				}
			}
		}
	}
}
