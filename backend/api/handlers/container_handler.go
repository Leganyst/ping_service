package handlers

import (
	"net/http"
	"vktest/repository"

	"github.com/gin-gonic/gin"
)

// @Summary Получить контейнеры
// @Description Возвращает список всех контейнеров
// @Tags Containers
// @Produce json
// @Success 200 {array} models.ContainerStatus
// @Failure 500 {object} models.ErrorResponse
// @Router /containers [get]
func GetContainers(c *gin.Context) {
	containers, err := repository.GetAllContainers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, containers)
}

/*
// CreateContainer добавляет новый контейнер
// @Summary Добавить контейнер
// @Description Создаёт новую запись о контейнере в базе данных
// @Tags Containers
// @Accept json
// @Produce json
// @Param container body models.ContainerStatus true "Данные контейнера"
// @Success 201 {object} models.ContainerStatus
// @Failure 400 {object} models.ErrorResponse
// @Router /containers [post]
func CreateContainer(c *gin.Context) {
	var container models.ContainerStatus

	body, _ := io.ReadAll(c.Request.Body)
	log.Printf("[DEBUG] Сырые данные JSON перед привязкой: %s", string(body))

	if err := json.Unmarshal(body, &container); err != nil {
		log.Printf("[ERROR] Ошибка ручной десериализации JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат JSON"})
		return
	}

	log.Printf("[DEBUG] JSON после десериализации: %+v", container)

	err := repository.CreateContainer(&container)
	if err != nil {
		log.Printf("[ERROR] Ошибка сохранения в БД: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("[INFO] Контейнер %s успешно сохранён: %+v", container.IPAddress, container)
	c.JSON(http.StatusCreated, container)
}
*/

// DeleteContainer удаляет контейнер по IP
// @Summary Удалить контейнер
// @Description Удаляет контейнер из базы данных по IP
// @Tags Containers
// @Param ip path string true "IP контейнера"
// @Success 200 {object} models.SuccessResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /containers/{ip} [delete]
func DeleteContainer(c *gin.Context) {
	ip := c.Param("ip")
	err := repository.DeleteContainer(ip)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Контейнер удалён"})
}
