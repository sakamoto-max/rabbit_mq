package queue

import (
	"context"
	"fmt"
	"log"
	"os"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	ApplicationJsonType string = "application/json"
)

func NewConn() *amqp.Connection {
	conn, err := amqp.Dial(os.Getenv("MQ_URL"))
	if err != nil {
		log.Fatalf("error opening a connection to rabbit mq : %v", err)
	}

	return conn
}

type MessageQueue struct {
	Ch    *amqp.Channel
	queue *amqp.Queue
}

func NewMessageQueue(conn *amqp.Connection, QueueName string) *MessageQueue {
	channel := createChannel(conn)
	queue := createQueue(channel, QueueName)

	return &MessageQueue{Ch: channel, queue: &queue}
}

type ConsumerChan chan<- amqp.Delivery

func (m *MessageQueue) Publish(ctx context.Context, data *[]byte) error {
	fmt.Println("data in bytes", data)
	msg := amqp.Publishing{
		ContentType:   ApplicationJsonType,
		Body:          *data,
	}
	fmt.Println("publishing", msg)
	fmt.Println("queue name", m.queue.Name)
	err := m.Ch.PublishWithContext(ctx, "", m.queue.Name, false, false, msg)
	if err != nil {
		return fmt.Errorf("error in publishing : %w", err)
	}
	return nil
}

func (m *MessageQueue) Consume() (<-chan amqp.Delivery, error) {
	consumerChan, err := m.Ch.Consume(m.queue.Name, "", true, false, false, false, nil)
	if err != nil {
		return consumerChan, fmt.Errorf("error occured while consuming from queue %v : %w", m.queue.Name, err)
	}

	return consumerChan, nil
}

func createChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("error in creating a channel : %v", err)
	}

	return ch
}
func createQueue(ch *amqp.Channel, queueName string) amqp.Queue {
	queue, err := ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		log.Fatalf("error creating %v : %v", queue.Name, err)
	}

	return queue
}