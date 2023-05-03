package handler

import (
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/handler"
	"github.com/2110336-2565-2/cu-freelance-chat/src/pkg/service"
	"github.com/2110336-2565-2/cu-freelance-chat/src/pkg/validator"
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
	"github.com/gofiber/websocket/v2"
)

func NewWebSocketChatHandler(authSvc service.AuthService, websocketSvc service.WebsocketService, keepAliveInterval int) WebsocketHandler {
	v, _ := validator.NewValidator()
	return &handler.ChatHandlerImpl{
		Logger:            gosdk.NewLogger("websocket"),
		Unregister:        make(chan *websocket.Conn),
		Broadcast:         make(chan []byte),
		AuthSvc:           authSvc,
		WebsocketSvc:      websocketSvc,
		KeepAliveInterval: keepAliveInterval,
		Validator:         v,
	}
}

type WebsocketHandler interface {
	Listen(c handler.IChatContext)
}
