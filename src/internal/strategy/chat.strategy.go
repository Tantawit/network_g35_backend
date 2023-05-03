package strategy

import (
	"encoding/json"
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/domain/dto"
	"github.com/2110336-2565-2/cu-freelance-chat/src/pkg/domain/entity"
	"github.com/2110336-2565-2/cu-freelance-chat/src/pkg/publisher"
	"github.com/2110336-2565-2/cu-freelance-chat/src/pkg/repository"
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/websocket/v2"
	"github.com/pkg/errors"
	"sync"
)

type WebsocketChatStrategy struct {
	Logger          gosdk.Logger
	ChatPub         publisher.ChatPublisher
	ChatSessionRepo repository.ChatSessionRepository
}

func (s *WebsocketChatStrategy) RegisterMappingService(userId string, serverId string) error {
	if userId == "" || serverId == "" {
		s.Logger.
			Warn().
			Keyword("user_id", userId).
			Keyword("server_id", serverId).
			Msg("error while registering chat client")
		return errors.New("invalid user id or server id")
	}

	if err := s.ChatSessionRepo.RegisterClient(userId, serverId); err != nil {
		s.Logger.
			Error(err).
			Keyword("user_id", userId).
			Keyword("server_id", serverId).
			Msg("error while registering chat client")
		return err
	}

	return nil
}

func (s *WebsocketChatStrategy) UnregisterMappingService(userId string) error {
	if err := s.ChatSessionRepo.UnregisterClient(userId); err != nil && !errors.Is(redis.Nil, err) {
		s.Logger.
			Error(err).
			Keyword("user_id", userId).
			Msg("error while unregistering chat client")
		return err
	}

	return nil
}

func (s *WebsocketChatStrategy) Broadcast(message string, senderId string, targetIdList []string, clientList map[string]entity.WebsocketClient) error {
	var wg sync.WaitGroup

	for _, targetId := range targetIdList {
		wg.Add(1)
		client, ok := clientList[targetId]
		if ok {
			go func(client entity.WebsocketClient) {
				defer wg.Done()
				client.Lock()
				defer client.Unlock()

				if err := client.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
					s.Logger.
						Error(err).
						Keyword("sender_id", senderId).
						Keyword("target_id", targetId).
						Msg("error while sending message to target")
					s.errorHandler(clientList[senderId])
					return
				}
			}(client)

			continue
		}

		go func(targetId string) {
			defer wg.Done()
			serverId, err := s.ChatSessionRepo.FindServer(targetId)
			if err != nil {
				s.Logger.
					Error(err).
					Keyword("target_id", targetId).
					Msg("error while finding server for target")
				s.errorHandler(clientList[senderId])
				return
			}

			if err := s.ChatPub.Broadcast(message, targetId, serverId); err != nil {
				s.Logger.
					Error(err).
					Keyword("target_id", targetId).
					Keyword("server_id", serverId).
					Msg("error while broadcasting to server")
				s.errorHandler(clientList[senderId])
				return
			}
		}(targetId)
	}

	wg.Wait()

	return nil
}

func (s *WebsocketChatStrategy) errorHandler(sender entity.WebsocketClient) {
	errMsg, _ := json.Marshal(dto.WebSocketErrorMessage{
		Type:    dto.ChatWebSocketErrorMessage,
		Message: "failed to send message",
	})

	if err := sender.WriteMessage(websocket.TextMessage, errMsg); err != nil {
		s.Logger.
			Error(err).
			Msg("error while sending error message to sender")
		return
	}
}
