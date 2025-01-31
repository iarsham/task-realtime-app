package ws

import (
	"github.com/iarsham/task-realtime-app/chat-service/domain"
	"go.uber.org/zap"
)

type Pool struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	unregister chan *Client
	logger     *zap.Logger
}

func NewPool(logger *zap.Logger) *Pool {
	return &Pool{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		unregister: make(chan *Client),
		logger:     logger,
	}
}

func (p *Pool) Run(brokerUsecase domain.BrokerUsecase) {
	defer func() {
		close(p.unregister)
		close(p.broadcast)
	}()
	for {
		select {
		case client := <-p.unregister:
			if _, ok := p.clients[client]; ok {
				delete(p.clients, client)
				close(client.send)
			}
		case msg := <-p.broadcast:
			brokerUsecase.PublishQueue("notification", msg)
			for client := range p.clients {
				select {
				case client.send <- msg:
				default:
					close(client.send)
					delete(p.clients, client)
				}
			}
		}
	}
}
