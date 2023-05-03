package router

import (
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/handler"
	"github.com/gofiber/websocket/v2"
)

func (r *FiberRouter) WebsocketChat(h func(c handler.IChatContext)) {
	r.Chat.Get("/ws", websocket.New(func(conn *websocket.Conn) {
		h(NewWebsocketCtx(conn))
	}))
}
