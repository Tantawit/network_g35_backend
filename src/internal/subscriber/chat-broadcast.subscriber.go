package subscriber

import (
	"encoding/json"
	"fmt"
	"github.com/2110336-2565-2/cu-freelance-chat/src/constant"
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/domain/dto"
	"github.com/2110336-2565-2/cu-freelance-chat/src/pkg/service"
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
)

type ChatBroadcastClientEventSubscriberImpl struct {
	Logger       gosdk.Logger
	Rabbitmq     gosdk.RabbitMQ
	WebsocketSvc service.WebsocketService
}

func (s *ChatBroadcastClientEventSubscriberImpl) Listen() {
	//TODO implement subscriber service
	raw, err := s.Rabbitmq.Listen()
	if err != nil {
		s.Logger.Error(err).
			Keyword("exchange", constant.ChatExchangeName).
			Keyword("topic", fmt.Sprintf("%v.%v", constant.ChatBroadcastTopicNameBase, constant.ServerID)).
			Msg("error while receive message from rabbitmq")
	}

	s.Logger.Info().
		Keyword("exchange", constant.ChatExchangeName).
		Keyword("topic", fmt.Sprintf("%v.%v", constant.ChatBroadcastTopicNameBase, constant.ServerID)).
		Msg("successfully receive raw from rabbitmq")

	event := dto.BroadcastEventDto{}
	if err := json.Unmarshal(raw, &event); err != nil {
		s.Logger.Error(err).
			Keyword("exchange", constant.ChatExchangeName).
			Keyword("topic", fmt.Sprintf("%v.%v", constant.ChatBroadcastTopicNameBase, constant.ServerID)).
			Keyword("raw", string(raw)).
			Msg("error while decoding message from rabbitmq")
		return
	}

	// business logic: broadcast to target
	if err := s.WebsocketSvc.Broadcast(event.Message, event.UserId, event.Targets); err != nil {
		s.Logger.Error(err).
			Keyword("exchange", constant.ChatExchangeName).
			Keyword("topic", fmt.Sprintf("%v.%v", constant.ChatBroadcastTopicNameBase, constant.ServerID)).
			Keyword("raw", string(raw)).
			Keyword("user_id", event.UserId).
			Keyword("targets", event.Targets).
			Msg("error while decoding message from rabbitmq")
	}

	s.Logger.Info().
		Keyword("exchange", constant.ChatExchangeName).
		Keyword("topic", fmt.Sprintf("%v.%v", constant.ChatBroadcastTopicNameBase, constant.ServerID)).
		Keyword("user_id", event.UserId).
		Keyword("targets", event.Targets).
		Msg("successfully receive raw from rabbitmq")
}

func (s *ChatBroadcastClientEventSubscriberImpl) Close() {
	s.Rabbitmq.Close()
}
