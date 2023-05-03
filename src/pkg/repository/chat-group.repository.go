package repository

import (
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/domain/entity"
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/repository"
	"github.com/gocql/gocql"
)

func NewChatGroupRepository(db *gocql.Session) ChatGroupRepository {
	return &repository.ChatGroupRepositoryImpl{Db: db}
}

type ChatGroupRepository interface {
	FindAllGroup(userId string, groupList *[]*entity.ChatUserMembership) error
	FindLatestMessage(groupId string, userId string, chat *entity.ChatGroup) error
	Create(chat *entity.ChatGroup) error
	Update(id string, chat *entity.ChatGroup) error
	Delete(id string) error
}
