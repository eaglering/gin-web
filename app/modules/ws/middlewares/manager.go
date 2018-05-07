package middlewares

import (
	"github.com/gorilla/websocket"
	"encoding/json"
	"time"
	"log"
	. "gin-web/app/config"
	"github.com/gin-gonic/gin"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type Manager struct{
	Clients		map[*Client]bool
	Broadcast	chan []byte
	Register	chan *Client
	Unregister	chan *Client
	router		map[string]gin.HandlerFunc
}

type Client struct {
	User	map[string]string
	UUID 	string
	Conn	*websocket.Conn
	Send	chan []byte
}

type Message struct {
	Sender string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	URI  	string `json:"uri,omitempty"`
	Content string `json:"content,omitempty"`
}

func (m *Manager) New() {
	for {
		select {
		case conn := <-m.Register:
			m.Clients[conn] = true
		case conn := <-m.Unregister:
			if _, ok := m.Clients[conn]; ok {
				close(conn.Send)
				delete(m.Clients, conn)
			}
		case message := <-m.Broadcast:
			for conn := range m.Clients {
				select {
				case conn.Send <- message:
				default:
					close(conn.Send)
					delete(m.Clients, conn)
				}
			}
		}
	}
}

func (m *Manager) GET (uri string, handlerFunc gin.HandlerFunc) {
	m.router[uri] = handlerFunc
}

func (m *Manager) Read(c *Client) {
	defer func() {
		m.Unregister <- c
		c.Conn.Close()
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	if Config.WebSocket.PongTimeout > 0 {
		c.Conn.SetReadDeadline(time.Now().Add(time.Duration(Config.WebSocket.PongTimeout)))
		c.Conn.SetPongHandler(func(string) error {
			c.Conn.SetReadDeadline(time.Now().Add(time.Duration(Config.WebSocket.PongTimeout)))
			return nil
		})
	}

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		var msg Message
		err = json.Unmarshal([]byte(message), &msg)
		if err != nil || m.router[msg.URI] == nil{
			continue
		}
		ctx := &gin.Context{
			Params: gin.Params{
				{Key:"recipient", Value:msg.Recipient},
				{Key:"content", Value:msg.Content},
			},
		}
		ctx.Set("manager", m)
		ctx.Set("sender", c)
		m.router[msg.URI](ctx)
	}
}

func (m *Manager) Write(c *Client) {
	//ticker := time.NewTicker(time.Duration(Config.WebSocket.PongTimeout * 9 / 10))
	defer func() {
		//ticker.Stop()
		c.Conn.Close()
	} ()

	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.Conn.WriteMessage(websocket.TextMessage, message)
		//case <-ticker.C:
		//	c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
		//	if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
		//		return
		//	}
		}
	}
}
