package repository

import (
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/domain/entity"
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/repository"
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
	"github.com/gocql/gocql"
)

func NewChatPrivateRepository(db *gocql.Session) ChatPrivateRepository {
	return &repository.ChatPrivateRepositoryImpl{Db: db}
}

type ChatPrivateRepository interface {
	// FindAllRecentMessage is the method use to find all recent message between the users with the pagination
	FindAllRecentMessage(metadata *gosdk.PaginationMetadata, user1Id string, user2Id string, chat *[]*entity.ChatPrivate) error
	// Create the chat message entity in database
	Create(chat *entity.ChatPrivate) error
	// Update the chat message entity in database
	Update(id string, chat *entity.ChatPrivate) error
	// Read mark the `read_at` timestamp to be the time that user read the message
	Read(user1Id, user2Id, messageId string) error
	// Delete the chat message in the database
	Delete(id string) error
}
