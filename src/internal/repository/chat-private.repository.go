package repository

import (
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/domain/entity"
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
)

type ChatPrivateRepositoryImpl struct {
	Db *gocql.Session
}

func (r *ChatPrivateRepositoryImpl) Read(user1Id string, user2Id string, messageId string) error {
	return r.Db.Query(`
		UPDATE chat_private SET read_at = ? WHERE user1_id = ? AND user2_id = ? AND message_id = ? IF EXISTS
	`, user1Id, user2Id, messageId).Exec()
}

func (r *ChatPrivateRepositoryImpl) FindAllRecentMessage(metadata *gosdk.PaginationMetadata, user1Id string, user2Id string, chat *[]*entity.ChatPrivate) error {
	return r.Db.Query("").Scan()
}

func (r *ChatPrivateRepositoryImpl) Create(chat *entity.ChatPrivate) error {
	return r.Db.Query(`
		INSERT INTO chat_private (user1_id, user2_id, message_id, sent_at, message)
		VALUES (?, ?, ?, toUnixTimestamp(now()), ?)
	`, chat.User1ID, chat.User2ID, uuid.NewString(), chat.Message).Exec()
}

func (r *ChatPrivateRepositoryImpl) Update(id string, chat *entity.ChatPrivate) error {
	return r.Db.Query("").Exec()
}

func (r *ChatPrivateRepositoryImpl) Delete(id string) error {
	return r.Db.Query("").Exec()
}
