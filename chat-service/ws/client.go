package ws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/iarsham/task-realtime-app/chat-service/domain"
	"github.com/iarsham/task-realtime-app/chat-service/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"time"
)

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
	maxMsgSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type Client struct {
	pool *Pool
	conn *websocket.Conn
	send chan []byte
}

func (c *Client) read(msgUsecase domain.MessageUsecase, userID primitive.ObjectID) {
	data := new(entities.MessageRequest)
	data.SenderID = userID
	defer func() {
		c.pool.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMsgSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		if err := c.conn.ReadJSON(&data); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.pool.logger.Error("unexpected close error while reading message", zap.Error(err))
			}
			break
		}
		createdMsg, err := msgUsecase.CreateMessage(data)
		if err != nil {
			c.pool.logger.Error("error while creating message", zap.Error(err))
		}
		msg, err := json.Marshal(createdMsg)
		if err != nil {
			c.pool.logger.Error("error while marshaling message", zap.Error(err))
		}
		c.pool.broadcast <- msg
	}
}
func (c *Client) write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case msg, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.pool.logger.Error("send channel closed while writing message")
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				c.pool.logger.Error("next writer error while writing message", zap.Error(err))
				return
			}
			w.Write(msg)
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}
			if err := w.Close(); err != nil {
				c.pool.logger.Error("writer close error while writing message", zap.Error(err))
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				c.pool.logger.Error("ping error while writing message", zap.Error(err))
				return
			}
		}
	}
}
