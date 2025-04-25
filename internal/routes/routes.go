package routes

import (
	"github.com/asylzhan/go-asylzhan-project/internal/auth"
	"github.com/asylzhan/go-asylzhan-project/internal/handler" // handler пакетін тек маршрута үшін қолданамыз
	"github.com/asylzhan/go-asylzhan-project/internal/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", auth.Register)
		authRoutes.POST("/login", auth.Login)
		authRoutes.GET("/me", middleware.AuthRequired(), auth.Me)
	}

	tireRoutes := r.Group("/api/tires")
	tireRoutes.Use(middleware.AuthRequired()) // барлық пайдаланушылар үшін
	{
		tireRoutes.GET("/", handler.GetAllTires) // Барлық пайдаланушыларға тауарларды көру
	}

	adminRoutes := r.Group("/admin")
	adminRoutes.Use(middleware.AuthRequired(), middleware.RoleRequired("admin"))
	{
		adminRoutes.POST("/tires", handler.AddTire)          // Тауар қосу
		adminRoutes.PUT("/tires/:id", handler.UpdateTire)    // Тауарды өзгерту
		adminRoutes.DELETE("/tires/:id", handler.DeleteTire) // Тауарды өшіру
	}
}
