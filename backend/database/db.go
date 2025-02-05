package database

import (
	"fmt"
	"log"
	"vktest/configs"
	"vktest/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	cfg := configs.LoadConfigDB()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Address, cfg.User, cfg.Password, cfg.DBName, cfg.Port)

	fmt.Printf("🛠 DSN строка: %s\n", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к БД: ", err)
	}
	err = db.AutoMigrate(&models.ContainerStatus{})
	if err != nil {
		log.Fatal("Ошибка миграции: ", err)
	}
	DB = db
	fmt.Println("Чики-бамбони. БД подключена")
}
