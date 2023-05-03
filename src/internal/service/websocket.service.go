package service

import (
	"github.com/2110336-2565-2/cu-freelance-chat/src/constant"
	"github.com/2110336-2565-2/cu-freelance-chat/src/pkg/domain/entity"
	"github.com/2110336-2565-2/cu-freelance-chat/src/pkg/strategy"
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
	"github.com/gofiber/websocket/v2"
)

type WebsocketChatServiceImpl struct {
	Logger     gosdk.Logger
	ClientList map[string]entity.WebsocketClient
	Strategy   strategy.WebsocketStrategy
}

func (s *WebsocketChatServiceImpl) RegisterClient(conn *websocket.Conn, userId string) error {
	serverId := constant.ServerID.String()

	if err := s.Strategy.RegisterMappingService(userId, serverId); err != nil {
		s.Logger.
			Error(err).
			Keyword("user_id", userId).
			Keyword("server_id", serverId).
			Msg("error while registering chat client")
		return err
	}

	s.ClientList[userId] = entity.NewWebsocketClient(conn)

	s.Logger.
		Info().
		Keyword("user_id", userId).
		Keyword("server_id", serverId).
		Msg("chat client registered")

	return nil
}

func (s *WebsocketChatServiceImpl) UnregisterClient(userId string) error {
	serverId := constant.ServerID.String()

	if err := s.Strategy.UnregisterMappingService(userId); err != nil {
		s.Logger.
			Error(err).
			Keyword("user_id", userId).
			Keyword("server_id", serverId).
			Msg("error while registering chat client")
		return err
	}

	delete(s.ClientList, userId)

	s.Logger.
		Info().
		Keyword("user_id", userId).
		Keyword("server_id", serverId).
		Msg("chat client registered")

	return nil
}

func (s *WebsocketChatServiceImpl) Broadcast(message string, senderId string, targetList []string) error {
	serverId := constant.ServerID.String()

	if err := s.Strategy.Broadcast(message, senderId, targetList, s.ClientList); err != nil {
		s.Logger.
			Error(err).
			Keyword("user_id", senderId).
			Keyword("server_id", serverId).
			Msg("error while broadcasting message")
		return err
	}

	return nil
}
