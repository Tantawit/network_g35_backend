package subscriber

import (
	"fmt"
	"github.com/2110336-2565-2/cu-freelance-chat/src/constant"
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/subscriber"
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
	"github.com/rabbitmq/amqp091-go"
)

func NewWebsocketBroadcastClientEventSubscriber(rabbitMQConn *amqp091.Connection) (Subscriber, error) {
	logger := gosdk.NewLogger("websocket-broadcast-client-subscriber")

	rabbitmq, err := gosdk.NewRabbitMQ(rabbitMQConn)
	if err != nil {
		logger.
			Fatal(err).
			Keyword("action", "init rabbitmq service").
			Msg("failed to init rabbitmq service")
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

	queue, err := rabbitmq.CreateQueue(
		constant.ChatBroadcastClientQueue,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	if err := rabbitmq.BindQueueWithExchange(
		queue.Name,
		fmt.Sprintf("%v.%v", constant.ChatBroadcastTopicNameBase, constant.ServerID),
		constant.ChatExchangeName,
		false,
		nil,
	); err != nil {
		return nil, err
	}

	if err := rabbitmq.CreateMessageChannel(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return nil, err
	}

	return &subscriber.ChatBroadcastClientEventSubscriberImpl{
		Logger:   logger,
		Rabbitmq: rabbitmq,
	}, nil
}
