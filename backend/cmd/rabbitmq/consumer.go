package rabbitmq

import (
	"encoding/json"
	"log"

	"vktest/models"
	"vktest/repository"
)

func ConsumeMessages() {
	InitRabbitMQ()

	msgs, err := GetChannel().Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Ошибка при подписке на очередь: %s", err)
	}

	log.Println("[INFO] Ожидание сообщений из RabbitMQ...")

	for msg := range msgs {
		var container models.ContainerStatus
		if err := json.Unmarshal(msg.Body, &container); err != nil {
			log.Printf("[ERROR] Ошибка десериализации JSON: %s", err)
			msg.Nack(false, false)
			continue
		}

		log.Printf("[INFO] Получено сообщение: %+v", container)

		err := repository.CreateContainer(&container)
		if err != nil {
			log.Printf("[ERROR] Ошибка сохранения в БД: %s", err)
			msg.Nack(false, false)
		} else {
			msg.Ack(false)
		}
	}
}
