package strategy

import (
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/strategy"
	"github.com/2110336-2565-2/cu-freelance-chat/src/pkg/domain/entity"
	"github.com/2110336-2565-2/cu-freelance-chat/src/pkg/publisher"
	"github.com/2110336-2565-2/cu-freelance-chat/src/pkg/repository"
)

func NewWebsocketChatStrategy(chatPub publisher.ChatPublisher, chatSessionRepo repository.ChatSessionRepository) WebsocketStrategy {
	return &strategy.WebsocketChatStrategy{
		ChatPub:         chatPub,
		ChatSessionRepo: chatSessionRepo,
	}
}

type WebsocketStrategy interface {
	RegisterMappingService(userId string, serverId string) error
	UnregisterMappingService(userId string) error
	Broadcast(message string, senderId string, targetId []string, clientList map[string]entity.WebsocketClient) error
}
