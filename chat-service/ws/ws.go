package ws

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ServeWs(pool *Pool, wr http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(wr, r, nil)
	if err != nil {
		return
	}
	client := &Client{
		pool: pool,
		conn: conn,
		send: make(chan []byte, 256),
	}
	client.pool.clients[client] = true
	go client.read()
	go client.write()
}
