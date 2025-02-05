package repository

import (
	"errors"
	"log"
	"time"
	"vktest/database"
	"vktest/models"

	"gorm.io/gorm"
)

func GetAllContainers() ([]models.ContainerStatus, error) {
	var containers []models.ContainerStatus
	result := database.DB.Find(&containers)
	return containers, result.Error
}

func GetContainerByIp(ip string) (*models.ContainerStatus, error) {
	var container models.ContainerStatus
	result := database.DB.Where("ip_address = ?", ip).First(&container)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &container, nil
}

func CreateContainer(container *models.ContainerStatus) error {
	log.Printf("[DEBUG] Данные перед сохранением в БД: %+v", container)

	existingContainer, err := GetContainerByIp(container.IPAddress)
	if err != nil {
		return err
	}

	if existingContainer != nil {
		log.Printf("[INFO] Контейнер %s найден, обновляем данные", container.IPAddress)
		existingContainer.Status = container.Status
		existingContainer.LastChecked = time.Now()

		result := database.DB.Save(existingContainer)
		return result.Error
	}

	// Устанавливаем время создания
	container.LastChecked = time.Now()

	log.Printf("[INFO] Создаём новый контейнер: %+v", container)
	result := database.DB.Create(container)
	return result.Error
}

func UpdateContainer(container *models.ContainerStatus) error {
	result := database.DB.Save(container)
	return result.Error
}

func DeleteContainer(ip string) error {
	result := database.DB.Where("ip_address = ?", ip).Delete(&models.ContainerStatus{})
	return result.Error
}
