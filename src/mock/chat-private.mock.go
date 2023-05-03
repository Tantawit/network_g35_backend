package mocks

import (
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/domain/entity"
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
	"github.com/stretchr/testify/mock"
)

type ChatPrivateRepositoryMock struct {
	mock.Mock
}

func (r *ChatPrivateRepositoryMock) FindAllRecentMessage(metadata *gosdk.PaginationMetadata, user1Id string, user2Id string, chat *[]*entity.ChatPrivate) error {
	//TODO implement me
	panic("implement me")
}

func (r *ChatPrivateRepositoryMock) Read(user1Id, user2Id, messageId string) error {
	args := r.Called(user1Id, user2Id, messageId)

	return args.Error(0)
}

func (r *ChatPrivateRepositoryMock) Create(chat *entity.ChatPrivate) error {
	args := r.Called(chat)

	if args.Get(0) != nil {
		*chat = *args.Get(0).(*entity.ChatPrivate)
	}

	return args.Error(1)
}

func (r *ChatPrivateRepositoryMock) Update(id string, chat *entity.ChatPrivate) error {
	//TODO implement me
	panic("implement me")
}

func (r *ChatPrivateRepositoryMock) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
