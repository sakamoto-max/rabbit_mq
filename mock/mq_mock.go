package mock

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

type MqMock struct {
	Open bool
}

func (m *MqMock) Consume() (<-chan amqp.Delivery, error) {
	if !m.Open {
		return nil, fmt.Errorf("mq is closed")
	}

	msgs := make(chan amqp.Delivery)

	return msgs, nil
}

func (m *MqMock) Publish(ctx context.Context, data *[]byte) error {
	if !m.Open {
		return fmt.Errorf("mq is closed")
	}

	return nil
}

