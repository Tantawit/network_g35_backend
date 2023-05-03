package service

import (
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/service"
	"github.com/2110336-2565-2/cu-freelance-chat/src/pkg/domain/entity"
	"github.com/2110336-2565-2/cu-freelance-chat/src/pkg/strategy"
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
	"github.com/gofiber/websocket/v2"
)

func NewWebsocketChatService(stg strategy.WebsocketStrategy) WebsocketService {
	// TODO register server in the mapping service

	return &service.WebsocketChatServiceImpl{
		Logger:     gosdk.NewLogger("websocket-service"),
		ClientList: map[string]entity.WebsocketClient{},
		Strategy:   stg,
	}
}

type WebsocketService interface {
	RegisterClient(conn *websocket.Conn, userId string) error
	UnregisterClient(userId string) error
	Broadcast(message string, senderId string, targetList []string) error
}
