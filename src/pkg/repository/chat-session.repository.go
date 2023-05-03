package repository

import (
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/repository"
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
	"github.com/go-redis/redis/v8"
)

func NewChatSessionRepository(client *redis.Client) ChatSessionRepository {
	return &repository.ChatSessionRepositoryImpl{
		Repo: gosdk.NewRedisRepository(client),
	}
}

type ChatSessionRepository interface {
	FindServer(userId string) (string, error)
	RegisterClient(userId string, serverId string) error
	UnregisterClient(userId string) error
}
