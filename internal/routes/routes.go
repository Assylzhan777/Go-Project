package routes

import (
	"github.com/asylzhan/go-asylzhan-project/internal/handler"
	"github.com/asylzhan/go-asylzhan-project/internal/repository"
	"github.com/asylzhan/go-asylzhan-project/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	repo := repository.NewTireRepository(db)
	svc := service.NewTireService(repo)
	h := handler.NewTireHandler(svc)

	api := r.Group("/api/tires")
	{
		api.POST("/", h.Create)
		api.GET("/", h.GetAll)
		api.GET("/:id", h.GetByID)
		api.PUT("/:id", h.Update)
		api.DELETE("/:id", h.Delete)
	}
}
