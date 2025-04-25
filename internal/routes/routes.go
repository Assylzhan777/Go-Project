package routes

import (
	"github.com/asylzhan/go-asylzhan-project/internal/auth"
	"github.com/asylzhan/go-asylzhan-project/internal/handler"
	"github.com/asylzhan/go-asylzhan-project/internal/middleware"
	"github.com/asylzhan/go-asylzhan-project/internal/repository"
	"github.com/asylzhan/go-asylzhan-project/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	tireRepo := repository.NewTireRepository(db)
	tireService := service.NewTireService(tireRepo)
	tireHandler := handler.NewTireHandler(tireService)

	tireRoutes := r.Group("/api/tires")
	{
		tireRoutes.GET("/", middleware.AuthRequired(), tireHandler.GetAll)
		tireRoutes.GET("/:id", middleware.AuthRequired(), tireHandler.GetByID)

		tireRoutes.POST("/", middleware.AuthRequired(), tireHandler.Create)
		tireRoutes.PUT("/:id", middleware.AuthRequired(), tireHandler.Update)
		tireRoutes.DELETE("/:id", middleware.AuthRequired(), tireHandler.Delete)
	}

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", auth.Register)
		authRoutes.POST("/login", auth.Login)
		authRoutes.GET("/me", middleware.AuthRequired(), auth.Me)
	}
}
