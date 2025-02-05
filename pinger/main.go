package main

import (
	"log"
	"time"
)

func main() {
	log.Println("[INFO] Загружаем конфигурацию...")
	config := LoadConfig()
	log.Println("[INFO] Конфигурация загружена успешно.")

	for {
		log.Println("[INFO] Запрашиваем информацию о контейнерах...")
		statuses, err := GetContainerStatuses()
		if err != nil {
			log.Printf("[ERROR] Ошибка получения данных о контейнерах: %v", err)
			time.Sleep(config.PingInterval)
			continue
		}

		log.Printf("[INFO] Найдено %d контейнеров", len(statuses))

		for _, container := range statuses {
			log.Printf("[INFO] Проверяем контейнер: ID=%s, IP=%s, Статус=%s, Health=%s",
				container.ID, container.IPAddress, container.Status, container.Health)

			finalStatus := "OK"
			if container.Status != "running" || container.Health == "unhealthy" {
				finalStatus = "FAIL"
			}

			result := PingResult{
				IPAddress:   container.IPAddress,
				Status:      finalStatus,
				LastChecked: time.Now(),
			}

			log.Printf("[INFO] Отправляем сообщение в RabbitMQ: %+v", result)
			SendMessage(result)
		}

		log.Printf("[INFO] Ждём %v перед следующим пингом...", config.PingInterval)
		time.Sleep(config.PingInterval)
	}
}
