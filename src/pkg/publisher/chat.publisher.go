package publisher

import (
	"github.com/2110336-2565-2/cu-freelance-chat/src/constant"
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/publisher"
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
	"github.com/rabbitmq/amqp091-go"
)

type ChatPublisher interface {
	Broadcast(message string, userId string, serverId string) error
}

func NewChatPublisher(rabbitMQConn *amqp091.Connection) (ChatPublisher, error) {
	rabbitmq, err := gosdk.NewRabbitMQ(rabbitMQConn)
	if err != nil {
		return nil, err
	}

	if err := rabbitmq.CreateExchange(
		constant.ChatExchangeName,
		constant.ChatExchangeKind,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return nil, err
	}

	return &publisher.ChatPublisherImpl{
		Rabbitmq: rabbitmq,
	}, nil
}
