package ws

import (
	"github.com/gorilla/websocket"
	"github.com/iarsham/task-realtime-app/chat-service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ServeWs(pool *Pool, userID primitive.ObjectID, msgUsecase domain.MessageUsecase, wr http.ResponseWriter, r *http.Request) {
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
	go client.read(msgUsecase, userID)
	go client.write()
}
