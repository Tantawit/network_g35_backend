package mocks

import (
	"github.com/stretchr/testify/mock"
)

type ChatPublisherMock struct {
	mock.Mock
}

func (p *ChatPublisherMock) Broadcast(message string, userId string, serverId string) error {
	args := p.Called(message, userId, serverId)

	return args.Error(0)
}

type ChatSessionRepositoryMock struct {
	mock.Mock
}

func (r *ChatSessionRepositoryMock) FindServer(userId string) (string, error) {
	args := r.Called(userId)

	return args.String(0), args.Error(1)
}

func (r *ChatSessionRepositoryMock) RegisterClient(userId string, serverId string) error {
	args := r.Called(userId, serverId)

	return args.Error(0)
}

func (r *ChatSessionRepositoryMock) UnregisterClient(userId string) error {
	args := r.Called(userId)

	return args.Error(0)
}
