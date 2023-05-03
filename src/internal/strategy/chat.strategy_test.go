package strategy

import (
	"github.com/2110336-2565-2/cu-freelance-chat/src/mock"
	"github.com/2110336-2565-2/cu-freelance-chat/src/pkg/domain/entity"
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ChatStrategyTest struct {
	suite.Suite
	Logger          gosdk.Logger
	ChatPub         *mocks.ChatPublisherMock
	ChatSessionRepo *mocks.ChatSessionRepositoryMock
	//ChatRepo        *mocks.ChatRepositoryMock
}

func TestChatStrategy(t *testing.T) {
	suite.Run(t, new(ChatStrategyTest))
}

func (t *ChatStrategyTest) SetupTest() {
	t.Logger = gosdk.NewLogger("strategy-test")
	t.ChatPub = &mocks.ChatPublisherMock{}
	//t.ChatRepo = &mocks.ChatRepositoryMock{}
	t.ChatSessionRepo = &mocks.ChatSessionRepositoryMock{}
}

func (t *ChatStrategyTest) TestRegisterSuccess() {
	userId := uuid.NewString()
	serverId := uuid.NewString()

	t.ChatSessionRepo.On("RegisterClient", userId, serverId).
		Return(nil).
		Once()

	stg := WebsocketChatStrategy{
		Logger:          t.Logger,
		ChatSessionRepo: t.ChatSessionRepo,
	}

	err := stg.RegisterMappingService(userId, serverId)

	assert.NoError(t.T(), err)
}

func (t *ChatStrategyTest) TestRegisterInternalError() {
	userId := uuid.NewString()
	serverId := uuid.NewString()

	t.ChatSessionRepo.On("RegisterClient", userId, serverId).
		Return(errors.New("internal error")).
		Once()

	stg := WebsocketChatStrategy{
		Logger:          t.Logger,
		ChatSessionRepo: t.ChatSessionRepo,
	}

	err := stg.RegisterMappingService(userId, serverId)

	assert.Error(t.T(), err)
}

func (t *ChatStrategyTest) TestUnregisterSuccess() {
	userId := uuid.NewString()

	t.ChatSessionRepo.On("UnregisterClient", userId).
		Return(nil).
		Once()

	stg := WebsocketChatStrategy{
		Logger:          t.Logger,
		ChatSessionRepo: t.ChatSessionRepo,
	}

	err := stg.UnregisterMappingService(userId)

	assert.NoError(t.T(), err)
}

func (t *ChatStrategyTest) TestUnregisterNotfound() {
	userId := uuid.NewString()

	t.ChatSessionRepo.On("UnregisterClient", userId).
		Return(redis.Nil).
		Once()

	stg := WebsocketChatStrategy{
		Logger:          t.Logger,
		ChatSessionRepo: t.ChatSessionRepo,
	}

	err := stg.UnregisterMappingService(userId)

	assert.NoError(t.T(), err)
}

func (t *ChatStrategyTest) TestUnregisterInternalError() {
	userId := uuid.NewString()

	t.ChatSessionRepo.On("UnregisterClient", userId).
		Return(errors.New("internal error")).
		Once()

	stg := WebsocketChatStrategy{
		Logger:          t.Logger,
		ChatSessionRepo: t.ChatSessionRepo,
	}

	err := stg.UnregisterMappingService(userId)

	assert.Error(t.T(), err)
}

func (t *ChatStrategyTest) TestBroadcastSingleClientInSameServerSuccess() {
	strategy := &WebsocketChatStrategy{
		Logger:  t.Logger,
		ChatPub: t.ChatPub,
		//ChatRepo:        t.ChatRepo,
		ChatSessionRepo: t.ChatSessionRepo,
	}

	senderId := uuid.NewString()
	user := uuid.NewString()

	client := &mocks.WebsocketClientMock{}
	client.On("WriteMessage", websocket.TextMessage, []byte("test message")).Return(nil).Once()
	client.On("Lock").Return().Once()
	client.On("Unlock").Return().Once()

	//t.ChatRepo.On("Create", entity.NewChatEntity(senderId, user, "test message")).
	//	Return(nil).
	//	Once()

	clientList := map[string]entity.WebsocketClient{
		user: client,
	}

	err := strategy.Broadcast("test message", senderId, []string{user}, clientList)
	assert.NoError(t.T(), err)

	client.AssertCalled(t.T(), "WriteMessage", websocket.TextMessage, []byte("test message"))
	t.ChatSessionRepo.AssertNotCalled(t.T(), "FindServer")
	t.ChatPub.AssertNotCalled(t.T(), "Broadcast")
}

func (t *ChatStrategyTest) TestBroadcastToMultipleClientInSameServerSuccess() {
	strategy := &WebsocketChatStrategy{
		Logger:          t.Logger,
		ChatPub:         t.ChatPub,
		ChatSessionRepo: t.ChatSessionRepo,
	}

	senderId := uuid.NewString()
	user1 := uuid.NewString()
	user2 := uuid.NewString()
	user3 := uuid.NewString()

	client1 := &mocks.WebsocketClientMock{}
	client1.On("WriteMessage", websocket.TextMessage, []byte("test message")).Return(nil).Once()
	client2 := &mocks.WebsocketClientMock{}
	client2.On("WriteMessage", websocket.TextMessage, []byte("test message")).Return(nil).Once()
	client3 := &mocks.WebsocketClientMock{}
	client3.On("WriteMessage", websocket.TextMessage, []byte("test message")).Return(nil).Once()
	client1.On("Lock").Return().Once()
	client1.On("Unlock").Return().Once()
	client2.On("Lock").Return().Once()
	client2.On("Unlock").Return().Once()
	client3.On("Lock").Return().Once()
	client3.On("Unlock").Return().Once()

	clientList := map[string]entity.WebsocketClient{
		user1: client1,
		user2: client2,
		user3: client3,
	}

	//chat1 := entity.NewChatEntity(senderId, user1, "test message")
	//chat2 := entity.NewChatEntity(senderId, user2, "test message")
	//chat3 := entity.NewChatEntity(senderId, user3, "test message")
	//
	//t.ChatRepo.On("Create", chat1).
	//	Return(nil).
	//	Once()
	//t.ChatRepo.On("Create", chat2).
	//	Return(nil).
	//	Once()
	//t.ChatRepo.On("Create", chat3).
	//	Return(nil).
	//	Once()

	err := strategy.Broadcast("test message", senderId, []string{user1, user2, user3}, clientList)
	assert.NoError(t.T(), err)

	t.ChatSessionRepo.AssertNotCalled(t.T(), "FindServer")
	t.ChatPub.AssertNotCalled(t.T(), "Broadcast")
	//t.ChatRepo.AssertExpectations(t.T())
	client1.AssertExpectations(t.T())
	client2.AssertExpectations(t.T())
	client3.AssertExpectations(t.T())
}

func (t *ChatStrategyTest) TestBroadcastToMultipleClientToOtherServerSuccess() {
	strategy := &WebsocketChatStrategy{
		Logger:          t.Logger,
		ChatPub:         t.ChatPub,
		ChatSessionRepo: t.ChatSessionRepo,
	}

	senderId := uuid.NewString()
	otherServerUser1 := uuid.NewString()
	otherServerUser2 := uuid.NewString()

	client1 := &mocks.WebsocketClientMock{}
	client2 := &mocks.WebsocketClientMock{}
	client1.On("WriteMessage", websocket.TextMessage, []byte("test message")).Return(nil).Once()
	client2.On("WriteMessage", websocket.TextMessage, []byte("test message")).Return(nil).Once()

	clientList := map[string]entity.WebsocketClient{}

	serverID1 := "server-2"
	serverID2 := "server-3"
	t.ChatSessionRepo.
		On("FindServer", otherServerUser1).
		Return(serverID1, nil).
		Once()
	t.ChatSessionRepo.
		On("FindServer", otherServerUser2).
		Return(serverID2, nil).
		Once()

	t.ChatPub.
		On("Broadcast", "test message", otherServerUser1, serverID1).
		Return(nil).
		Once()
	t.ChatPub.
		On("Broadcast", "test message", otherServerUser2, serverID2).
		Return(nil).
		Once()

	err := strategy.Broadcast("test message", senderId, []string{otherServerUser1, otherServerUser2}, clientList)
	assert.NoError(t.T(), err)

	client1.AssertNotCalled(t.T(), "WriteMessage")
	client2.AssertNotCalled(t.T(), "WriteMessage")
	t.ChatSessionRepo.AssertCalled(t.T(), "FindServer", otherServerUser1)
	t.ChatSessionRepo.AssertCalled(t.T(), "FindServer", otherServerUser2)
	t.ChatPub.AssertCalled(t.T(), "Broadcast", "test message", otherServerUser1, serverID1)
	t.ChatPub.AssertCalled(t.T(), "Broadcast", "test message", otherServerUser2, serverID2)

}

func (t *ChatStrategyTest) TestBroadcastSingleClientToOtherServerSuccess() {
	strategy := &WebsocketChatStrategy{
		Logger:          t.Logger,
		ChatPub:         t.ChatPub,
		ChatSessionRepo: t.ChatSessionRepo,
	}

	senderId := uuid.NewString()
	user := uuid.NewString()
	otherServerUser := uuid.NewString()

	client := &mocks.WebsocketClientMock{}
	client.On("WriteMessage", websocket.TextMessage, []byte("test message")).Return(nil).Once()

	clientList := map[string]entity.WebsocketClient{
		user: client,
	}

	serverID := "server-2"
	t.ChatSessionRepo.
		On("FindServer", otherServerUser).
		Return(serverID, nil).
		Once()

	t.ChatPub.
		On("Broadcast", "test message", otherServerUser, serverID).
		Return(nil).
		Once()

	err := strategy.Broadcast("test message", senderId, []string{otherServerUser}, clientList)
	assert.NoError(t.T(), err)

	client.AssertNotCalled(t.T(), "WriteMessage")
	t.ChatSessionRepo.AssertCalled(t.T(), "FindServer", otherServerUser)
	t.ChatPub.AssertCalled(t.T(), "Broadcast", "test message", otherServerUser, serverID)
}
