package main

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/streadway/amqp"
)

const queueName = "containers_queue"

var (
	conn    *amqp.Connection
	channel *amqp.Channel
	once    sync.Once
)

func initRabbitMQ() {
	config := LoadConfig()

	var err error
	conn, err = amqp.Dial(config.RabbitMQURL)
	if err != nil {
		log.Fatalf("Ошибка подключения к RabbitMQ: %s", err)
	}

	channel, err = conn.Channel()
	if err != nil {
		log.Fatalf("Ошибка при создании канала: %s", err)
	}

	_, err = channel.QueueDeclare(
		queueName,
		true,  // Долговечная очередь
		false, // Не удалять при отсутствии подписчиков
		false, // Не эксклюзивная
		false, // Без ожидания
		nil,
	)
	if err != nil {
		log.Fatalf("Ошибка при объявлении очереди: %s", err)
	}
}

func SendMessage(container PingResult) {
	once.Do(initRabbitMQ)

	message, err := json.Marshal(container)
	if err != nil {
		log.Printf("Ошибка кодирования JSON: %s", err)
		return
	}

	err = channel.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         message,
			DeliveryMode: amqp.Persistent,
		},
	)

	if err != nil {
		log.Printf("Ошибка отправки сообщения: %s", err)
	} else {
		log.Printf(" [x] Сообщение отправлено: %s", message)
	}
}
