package ws

import (
	"github.com/Reza-Rayan/internal/applications"
	"github.com/Reza-Rayan/internal/domain/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type WSHandler struct {
	hub            *Hub
	messageUsecase *applications.MessageUsecase
}

func NewWSHandler(hub *Hub, messageUC *applications.MessageUsecase) *WSHandler {
	return &WSHandler{
		hub:            hub,
		messageUsecase: messageUC,
	}
}

func (h *WSHandler) ServeWs() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userIDParam := ctx.Param("user_id")
		userID64, _ := strconv.ParseUint(userIDParam, 10, 64)
		userID := uint(userID64)

		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Println("WebSocket upgrade error:", err)
			return
		}
		client := &Client{
			UserID: userID,
			Conn:   conn,
			Send:   make(chan *models.Message, 256),
		}

		h.hub.Register <- client

		go h.readPump(client)
		go h.writePump(client)
	}
}

func (h *WSHandler) readPump(c *Client) {
	defer func() {
		h.hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		var msg models.Message
		err := c.Conn.ReadJSON(&msg)
		if err != nil {
			log.Println("read error:", err)
			break
		}

		//		Save Message in database
		savedMsg, err := h.messageUsecase.SendMessage(msg.SenderID, msg.ReceiverID, msg.Content)
		if err != nil {
			log.Println("save message error:", err)
			continue
		}
		// Send message to Receiver
		h.hub.Broadcast <- savedMsg

	}
}

func (h *WSHandler) writePump(c *Client) {
	defer c.Conn.Close()

	for msg := range c.Send {
		if err := c.Conn.WriteJSON(msg); err != nil {
			log.Println("write error:", err)
			break
		}
	}
}
