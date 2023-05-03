package entity

import (
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/domain/entity"
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
	"github.com/gofiber/websocket/v2"
	"sync"
)

func NewWebsocketClient(conn *websocket.Conn) WebsocketClient {
	return &entity.WebsocketClient{
		Conn:      conn,
		IsClosing: gosdk.BoolAdr(false),
		Mu:        &sync.Mutex{},
	}
}

type WebsocketClient interface {
	WriteMessage(messageType int, data []byte) error
	ReadMessage() (messageType int, p []byte, err error)
	IsClose() bool
	Lock()
	Unlock()
}
