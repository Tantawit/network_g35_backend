package service

import (
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/domain/entity"
	"github.com/2110336-2565-2/cu-freelance-chat/src/pkg/repository"
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
)

type ChatGroupServiceImpl struct {
	Logger              gosdk.Logger
	ChatGroupRepository repository.ChatGroupRepository
}

func (s *ChatGroupServiceImpl) SendMessage(chat *entity.ChatGroup) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ChatGroupServiceImpl) GetGroupList(userId string) ([]*entity.ChatGroupMembership, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ChatGroupServiceImpl) JoinGroup(groupId string, userId string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ChatGroupServiceImpl) LeaveGroup(groupId string, userId string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ChatGroupServiceImpl) Sync(syncToken string) ([]*entity.ChatGroup, error) {
	//TODO implement me
	panic("implement me")
}
