package ws

import (
	"github.com/Reza-Rayan/internal/domain/models"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

type Client struct {
	UserID uint
	Conn   *websocket.Conn
	Send   chan *models.Message
}

type Hub struct {
	Clients    map[uint]*Client
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *models.Message
	mu         sync.RWMutex
}

// NewHub -> Create new hub if it does not exists
func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[uint]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *models.Message),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.mu.Lock()
			h.Clients[client.UserID] = client
			h.mu.Unlock()
			log.Printf("User %d connected", client.UserID)

		case client := <-h.Unregister:
			h.mu.Lock()
			if _, ok := h.Clients[client.UserID]; ok {
				delete(h.Clients, client.UserID)
				close(client.Send)
			}
			h.mu.Unlock()
			log.Printf("User %d disconnected", client.UserID)

		case message := <-h.Broadcast:
			h.mu.Lock()
			if receiver, ok := h.Clients[message.ReceiverID]; ok {
				select {
				case receiver.Send <- message:
				default:
					log.Printf("Send channel full for user %d", receiver.UserID)
				}
			}
			h.mu.Unlock()
		}
	}
}
