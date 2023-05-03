package service

import (
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/domain/entity"
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/service"
	"github.com/2110336-2565-2/cu-freelance-chat/src/pkg/repository"
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
)

func NewChatGroupService(
	chatRepository repository.ChatGroupRepository,
) ChatGroupService {
	logger := gosdk.NewLogger("chat-group-service")
	return &service.ChatGroupServiceImpl{
		Logger:              logger,
		ChatGroupRepository: chatRepository,
	}
}

type ChatGroupService interface {
	SendMessage(chat *entity.ChatGroup) (bool, error)
	GetGroupList(userId string) ([]*entity.ChatGroupMembership, error)
	JoinGroup(groupId string, userId string) (bool, error)
	LeaveGroup(groupId string, userId string) (bool, error)
	Sync(syncToken string) ([]*entity.ChatGroup, error)
}
