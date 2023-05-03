package mocks

import (
	"github.com/2110336-2565-2/cu-freelance-chat/src/pkg/domain/entity"
	"github.com/stretchr/testify/mock"
)

type WebsocketStrategyMock struct {
	mock.Mock
}

func (s *WebsocketStrategyMock) RegisterMappingService(userId string, serverId string) error {
	args := s.Called(userId, serverId)
	return args.Error(0)
}

func (s *WebsocketStrategyMock) UnregisterMappingService(userId string) error {
	args := s.Called(userId)
	return args.Error(0)
}

func (s *WebsocketStrategyMock) Broadcast(message string, senderId string, targetId []string, clientList map[string]entity.WebsocketClient) error {
	args := s.Called(message, senderId, targetId, clientList)
	return args.Error(0)
}

type WebsocketClientMock struct {
	mock.Mock
}

func (c *WebsocketClientMock) IsClose() bool {
	args := c.Called()

	return args.Bool(0)
}

func (c *WebsocketClientMock) Lock() {
	_ = c.Called()

	return
}

func (c *WebsocketClientMock) Unlock() {
	_ = c.Called()

	return
}

func (c *WebsocketClientMock) WriteMessage(messageType int, data []byte) error {
	args := c.Called(messageType, data)

	return args.Error(0)
}

func (c *WebsocketClientMock) ReadMessage() (messageType int, p []byte, err error) {
	args := c.Called()

	if args.Get(1) != nil {
		p = args.Get(1).([]byte)
	}

	return args.Int(0), p, args.Error(2)
}
