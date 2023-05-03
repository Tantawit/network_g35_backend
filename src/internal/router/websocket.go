package router

import (
	"github.com/gofiber/websocket/v2"
)

type WebSocketContext struct {
	*websocket.Conn
}

func (c *WebSocketContext) GetConn() *websocket.Conn {
	return c.Conn
}

func NewWebSocketContext(c *websocket.Conn) *WebSocketContext {
	return &WebSocketContext{
		c,
	}
}

func (c *WebSocketContext) ID() string {
	return c.Params("id", "")
}

func NewWebsocketCtx(c *websocket.Conn) *WebSocketContext {
	return &WebSocketContext{Conn: c}
}
