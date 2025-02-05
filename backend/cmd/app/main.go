package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"vktest/api/routes"
	"vktest/cmd/rabbitmq"
	"vktest/configs"
	"vktest/database"
	_ "vktest/docs"

	"github.com/gin-contrib/cors"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Container Monitoring API
// @version 1.0
// @description API для мониторинга контейнеров Docker.
// @host localhost:8080
// @BasePath /
func main() {
	cfg := configs.LoadConfigDB()

	fmt.Println("🔍 Переменные окружения:")
	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	fmt.Println("DB_USER:", os.Getenv("DB_USER"))
	fmt.Println("DB_PORT:", os.Getenv("DB_PORT"))

	fmt.Printf("Config struct: %+v\n", cfg)

	database.ConnectDB()
	go rabbitmq.ConsumeMessages()
	r := routes.SetupRouter()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := ":8080"

	fmt.Println("Backend started at", port)
	log.Fatal(r.Run(port))
}
