package service

import (
	"github.com/2110336-2565-2/cu-freelance-chat/src/constant"
	mocks "github.com/2110336-2565-2/cu-freelance-chat/src/mock"
	"github.com/2110336-2565-2/cu-freelance-chat/src/pkg/domain/entity"
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type WebsocketServiceTest struct {
	suite.Suite
	Logger     gosdk.Logger
	ClientList map[string]entity.WebsocketClient
	Strategy   *mocks.WebsocketStrategyMock
}

func TestWebsocketService(t *testing.T) {
	suite.Run(t, new(WebsocketServiceTest))
}

func (t *WebsocketServiceTest) SetupTest() {
	t.Logger = gosdk.NewLogger("server-test")
	constant.ServerID = gosdk.UUIDAdr(uuid.New())
	t.ClientList = map[string]entity.WebsocketClient{}
	t.Strategy = &mocks.WebsocketStrategyMock{}
}

func (t *WebsocketServiceTest) TestRegisterClientSuccess() {
	conn := websocket.Conn{}

	userId := uuid.NewString()

	t.Strategy.On("RegisterMappingService", userId, mock.AnythingOfType("string")).
		Return(nil).
		Once()

	srv := WebsocketChatServiceImpl{
		Logger:     t.Logger,
		ClientList: t.ClientList,
		Strategy:   t.Strategy,
	}

	err := srv.RegisterClient(&conn, userId)

	assert.NoError(t.T(), err)
	assert.Len(t.T(), srv.ClientList, 1)
	assert.Contains(t.T(), srv.ClientList, userId)
	assert.NotNil(t.T(), srv.ClientList[userId])
}

func (t *WebsocketServiceTest) TestRegisterClientInternalError() {
	conn := websocket.Conn{}

	userId := uuid.NewString()

	expectedErr := errors.New("some error")

	t.Strategy.On("RegisterMappingService", userId, mock.AnythingOfType("string")).
		Return(expectedErr).
		Once()

	srv := WebsocketChatServiceImpl{
		Logger:     t.Logger,
		ClientList: t.ClientList,
		Strategy:   t.Strategy,
	}

	err := srv.RegisterClient(&conn, userId)

	assert.Error(t.T(), err)
	assert.Equal(t.T(), expectedErr, err)
}

func (t *WebsocketServiceTest) TestUnregisterClientSuccess() {
	userId := uuid.NewString()

	t.Strategy.On("UnregisterMappingService", userId).
		Return(nil).
		Once()

	t.ClientList[userId] = &mocks.WebsocketClientMock{}

	srv := WebsocketChatServiceImpl{
		Logger:     t.Logger,
		ClientList: t.ClientList,
		Strategy:   t.Strategy,
	}

	err := srv.UnregisterClient(userId)

	assert.NoError(t.T(), err)
	assert.NotContains(t.T(), t.ClientList, userId)
}

func (t *WebsocketServiceTest) TestUnregisterClientInternalError() {

	userId := uuid.NewString()

	t.Strategy.On("UnregisterMappingService", userId).
		Return(errors.New("error while unregistering client")).
		Once()

	srv := WebsocketChatServiceImpl{
		Logger:     t.Logger,
		ClientList: t.ClientList,
		Strategy:   t.Strategy,
	}

	err := srv.UnregisterClient(userId)

	assert.Error(t.T(), err)
}

func (t *WebsocketServiceTest) TestBroadcastSuccess() {
	senderId := uuid.NewString()
	user1 := uuid.NewString()
	user2 := uuid.NewString()
	user3 := uuid.NewString()

	client1 := &mocks.WebsocketClientMock{}
	client2 := &mocks.WebsocketClientMock{}
	client3 := &mocks.WebsocketClientMock{}

	clientList := map[string]entity.WebsocketClient{
		user1: client1,
		user2: client2,
		user3: client3,
	}

	t.Strategy.On("Broadcast", "test message", senderId, []string{user1, user2, user3}, clientList).
		Return(nil).
		Once()

	srv := WebsocketChatServiceImpl{
		Logger:     t.Logger,
		ClientList: clientList,
		Strategy:   t.Strategy,
	}

	err := srv.Broadcast("test message", senderId, []string{user1, user2, user3})

	assert.NoError(t.T(), err)
}

func (t *WebsocketServiceTest) TestBroadcastInternalError() {
	senderId := uuid.NewString()
	targetList := []string{uuid.NewString(), uuid.NewString()}
	message := "test message"

	t.Strategy.On("Broadcast", message, senderId, targetList, t.ClientList).
		Return(errors.New("internal error")).
		Once()

	srv := WebsocketChatServiceImpl{
		Logger:     t.Logger,
		ClientList: t.ClientList,
		Strategy:   t.Strategy,
	}

	err := srv.Broadcast(message, senderId, targetList)

	assert.Error(t.T(), err)
}
