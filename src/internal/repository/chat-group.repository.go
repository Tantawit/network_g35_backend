package repository

import (
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/domain/entity"
	"github.com/gocql/gocql"
)

type ChatGroupRepositoryImpl struct {
	Db *gocql.Session
}

func (r *ChatGroupRepositoryImpl) FindAllGroup(userId string, groupList *[]*entity.ChatUserMembership) error {
	iter := r.Db.Query("").Iter()
	for iter.Scan() {

	}

	return iter.Close()
}

func (r *ChatGroupRepositoryImpl) FindLatestMessage(groupId string, userId string, chat *entity.ChatGroup) error {
	return r.Db.Query("").Scan()
}

func (r *ChatGroupRepositoryImpl) Create(chat *entity.ChatGroup) error {
	return r.Db.Query("").Exec()
}

func (r *ChatGroupRepositoryImpl) Update(id string, chat *entity.ChatGroup) error {
	return r.Db.Query("").Exec()
}

func (r *ChatGroupRepositoryImpl) Delete(id string) error {
	return r.Db.Query("").Exec()
}
