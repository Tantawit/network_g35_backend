package entity

import (
	"github.com/gofiber/websocket/v2"
	"sync"
)

type WebsocketClient struct {
	*websocket.Conn
	IsClosing *bool
	Mu        *sync.Mutex
}

func (e WebsocketClient) Lock() {
	e.Mu.Lock()
}

func (e WebsocketClient) Unlock() {
	e.Mu.Unlock()
}

func (e WebsocketClient) IsClose() bool {
	return *e.IsClosing
}
