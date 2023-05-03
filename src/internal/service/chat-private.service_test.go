package service

import (
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/domain/entity"
	mocks "github.com/2110336-2565-2/cu-freelance-chat/src/mock"
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type ChatPrivateServiceTest struct {
	suite.Suite
	Logger gosdk.Logger
	Repo   *mocks.ChatPrivateRepositoryMock
	Chat   *entity.ChatPrivate
}

func TestChatPrivateService(t *testing.T) {
	suite.Run(t, new(ChatPrivateServiceTest))
}

func (t *ChatPrivateServiceTest) SetupTest() {
	t.Logger = gosdk.NewLogger("chat-private-service-test")
	t.Repo = &mocks.ChatPrivateRepositoryMock{}
	t.Chat = &entity.ChatPrivate{
		User1ID:   gosdk.UUIDAdr(uuid.New()),
		User2ID:   gosdk.UUIDAdr(uuid.New()),
		MessageID: gosdk.UUIDAdr(uuid.New()),
		SentAt:    time.Now(),
		ReadAt:    time.Now(),
		Message:   faker.Word(),
	}
}

func (t *ChatPrivateServiceTest) TestSendMessageSuccess() {
	t.Repo.On("Create", mock.AnythingOfType("*entity.ChatPrivate")).
		Return(t.Chat, nil)

	svc := ChatPrivateServiceImpl{
		Logger:                t.Logger,
		ChatPrivateRepository: t.Repo,
	}

	actual, err := svc.SendMessage(&entity.ChatPrivate{
		User1ID:   t.Chat.User1ID,
		User2ID:   t.Chat.User2ID,
		MessageID: t.Chat.MessageID,
		Message:   t.Chat.Message,
	})

	t.NoError(err)
	t.True(actual)
}

func (t *ChatPrivateServiceTest) TestSendMessageInternalError() {
	t.Repo.On("Create", mock.AnythingOfType("*entity.ChatPrivate")).
		Return(nil, errors.New("internal error"))

	svc := ChatPrivateServiceImpl{
		Logger:                t.Logger,
		ChatPrivateRepository: t.Repo,
	}

	actual, err := svc.SendMessage(&entity.ChatPrivate{
		User1ID:   t.Chat.User1ID,
		User2ID:   t.Chat.User2ID,
		MessageID: t.Chat.MessageID,
		Message:   t.Chat.Message,
	})

	t.Error(err)
	t.False(actual)
}

func (t *ChatPrivateServiceTest) TestReadMessageSuccess() {
	t.Repo.On("Read", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(nil)

	svc := ChatPrivateServiceImpl{
		Logger:                t.Logger,
		ChatPrivateRepository: t.Repo,
	}

	actual, err := svc.ReadMessage(uuid.NewString(), uuid.NewString(), uuid.NewString())

	t.NoError(err)
	t.True(actual)
}

func (t *ChatPrivateServiceTest) TestReadMessageInternalError() {
	t.Repo.On("Read", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(errors.New("internal error"))

	svc := ChatPrivateServiceImpl{
		Logger:                t.Logger,
		ChatPrivateRepository: t.Repo,
	}

	actual, err := svc.ReadMessage(uuid.NewString(), uuid.NewString(), uuid.NewString())

	t.Error(err)
	t.False(actual)
}

func (t *ChatPrivateServiceTest) TestFindAllRecentChatSuccess() {
	//t.Repo.On("FindAllRecentChat", mock.AnythingOfType("string"), mock.AnythingOfType("[]*entity.ChatPrivate")).
	//	Return(t.Chat, nil)
	//
	//svc := ChatPrivateServiceImpl{
	//	Logger:                t.Logger,
	//	ChatPrivateRepository: t.Repo,
	//}
	//
	//actual, err := svc.FindAllRecentChat(uuid.NewString())
	//
	//t.NoError(err)
	//t.Equal(actual)
}

func (t *ChatPrivateServiceTest) TestFindAllRecentChatInternal() {}
