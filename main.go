package main

import (
	"log"

	"github.com/asylzhan/go-asylzhan-project/internal/models"
	"github.com/asylzhan/go-asylzhan-project/internal/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "postgres://postgres:Asilzhan7@localhost:5432/postgres?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	if err := db.AutoMigrate(&models.Tire{}); err != nil {
		log.Fatal("Migration failed:", err)
	}

	r := gin.Default()
	routes.SetupRoutes(r, db)
	r.Run(":8080")
}
