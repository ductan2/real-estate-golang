package queue

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

type ProjectTask struct {
	ProjectID string    `json:"project_id"`
	Action    string    `json:"action"`
	ExecuteAt time.Time `json:"execute_at"`
}

type IQueueService interface {
	PublishProjectTask(task *ProjectTask, exchange string, routingKey string) error
	PublishMessage(msg interface{}, exchange string, routingKey string) error
	ConsumeProjectTasks() (<-chan ProjectTask, error)
	ConsumeMessages(queueName string) (<-chan amqp.Delivery, error)
	CreateQueueAndBind(exchange, queueName, routingKey string) error
	Close() error
}

type QueueService struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewQueueService(amqpURL string) (*QueueService, error) {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &QueueService{
		conn:    conn,
		channel: ch,
	}, nil
}

func (s *QueueService) CreateQueueAndBind(exchange, queueName, routingKey string) error {
	// Declare exchange
	if err := s.channel.ExchangeDeclare(
		exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return fmt.Errorf("failed to declare exchange: %v", err)
	}

	// Declare queue
	_, err := s.channel.QueueDeclare(
		queueName, // queue name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %v", err)
	}

	// Bind queue to exchange
	err = s.channel.QueueBind(
		queueName,  // queue name
		routingKey, // routing key
		exchange,   // exchange
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to bind queue: %v", err)
	}

	return nil
}

func (s *QueueService) PublishProjectTask(task *ProjectTask, exchange string, routingKey string) error {
	body, err := json.Marshal(task)
	if err != nil {
		return err
	}
	return s.channel.Publish(
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
}

func (s *QueueService) PublishMessage(msg interface{}, exchange string, routingKey string) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return s.channel.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
}

func (s *QueueService) ConsumeProjectTasks() (<-chan ProjectTask, error) {
	msgs, err := s.channel.Consume(
		"project_tasks", // queue
		"",              // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)
	if err != nil {
		return nil, err
	}

	tasks := make(chan ProjectTask)
	go func() {
		for msg := range msgs {
			var task ProjectTask
			if err := json.Unmarshal(msg.Body, &task); err != nil {
				continue
			}
			tasks <- task
		}
	}()

	return tasks, nil
}

func (s *QueueService) ConsumeMessages(queueName string) (<-chan amqp.Delivery, error) {
	return s.channel.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
}

func (s *QueueService) Close() error {
	if err := s.channel.Close(); err != nil {
		return err
	}
	return s.conn.Close()
}
