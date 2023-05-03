package service

import (
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/domain/entity"
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/service"
	"github.com/2110336-2565-2/cu-freelance-chat/src/pkg/repository"
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
)

func NewChatPrivateService(
	chatRepository repository.ChatPrivateRepository,
) ChatPrivateService {
	logger := gosdk.NewLogger("chat-private-service")
	return &service.ChatPrivateServiceImpl{
		Logger:                logger,
		ChatPrivateRepository: chatRepository,
	}
}

type ChatPrivateService interface {
	// SendMessage save chat message to database
	SendMessage(chat *entity.ChatPrivate) (bool, error)

	// ReadMessage mark message that it was read
	ReadMessage(user1Id string, user2Id string, messageId string) (bool, error)

	// FindAllRecentChat that was sent to user
	FindAllRecentChat(userId string) ([]*entity.ChatPrivate, error)

	// FindMessage that in chat with pagination
	FindMessage(metadata *gosdk.PaginationMetadata, user1ID string, user2ID string) ([]*entity.ChatPrivate, error)

	// Sync is to find the message since the recent one
	//
	// syncToken is the encryption of `sent_at` timestamp
	Sync(metadata *gosdk.PaginationMetadata, user1ID string, user2ID string, syncToken string) ([]*entity.ChatPrivate, error)
}
