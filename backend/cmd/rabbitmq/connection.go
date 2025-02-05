package rabbitmq

import (
	"log"
	"os"
	"sync"

	"github.com/streadway/amqp"
)

const queueName = "containers_queue"

var (
	conn    *amqp.Connection
	channel *amqp.Channel
	once    sync.Once
)

func InitRabbitMQ() {
	once.Do(func() {
		rabbitMQURL := os.Getenv("RABBITMQ_URL")
		if rabbitMQURL == "" {
			rabbitMQURL = "amqp://guest:guest@rabbitmq:5672/"
		}

		var err error
		conn, err = amqp.Dial(rabbitMQURL)
		if err != nil {
			log.Fatalf("Ошибка подключения к RabbitMQ: %s", err)
		}

		channel, err = conn.Channel()
		if err != nil {
			log.Fatalf("Ошибка создания канала: %s", err)
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

		log.Println("[INFO] Соединение с RabbitMQ установлено")
	})
}

func GetChannel() *amqp.Channel {
	return channel
}
