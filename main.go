package main

import (
	"github.com/asylzhan/go-asylzhan-project/internal/db"
	"github.com/asylzhan/go-asylzhan-project/internal/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// DB инициализациясы
	db.InitDB()

	// Gin сервері
	r := gin.Default()

	// Роуттар орнату
	routes.SetupRoutes(r, db.DB)

	// Серверді іске қосу
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
