package publisher

import (
	"fmt"
	"github.com/2110336-2565-2/cu-freelance-chat/src/constant"
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/domain/dto"
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
)

type ChatPublisherImpl struct {
	Rabbitmq gosdk.RabbitMQ
}

func (p *ChatPublisherImpl) Broadcast(message string, userId string, serverId string) error {
	return p.Rabbitmq.Publish(
		constant.ChatExchangeName,
		fmt.Sprintf("%v.%v", constant.ChatBroadcastTopicNameBase, serverId),
		&dto.BroadcastEventDto{
			Message: message,
			UserId:  userId,
		},
	)
}
