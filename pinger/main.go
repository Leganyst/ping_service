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

			// Определяем финальный статус контейнера
			finalStatus := "OK"
			if container.Status != "running" || container.Health == "unhealthy" {
				finalStatus = "FAIL"
			}

			// Формируем результат для отправки, добавляем текущее время
			result := PingResult{
				IPAddress:   container.IPAddress,
				Status:      finalStatus,
				LastChecked: time.Now(), // 💡 Добавляем метку времени
			}

			log.Printf("[INFO] Отправляем результат пинга: %+v", result)
			err := SendPingResult(result, config.BackendURL)
			if err != nil {
				log.Printf("[ERROR] Ошибка отправки результата для %s: %v", container.IPAddress, err)
			} else {
				log.Printf("[INFO] Результат успешно отправлен для %s", container.IPAddress)
			}
		}

		log.Printf("[INFO] Ждём %v перед следующим пингом...", config.PingInterval)
		time.Sleep(config.PingInterval)
	}
}
