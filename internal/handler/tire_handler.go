package handler

import (
	"github.com/asylzhan/go-asylzhan-project/internal/db"
	"github.com/asylzhan/go-asylzhan-project/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllTires(c *gin.Context) {
	var tires []models.Tire
	if err := db.DB.Find(&tires).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tires"})
		return
	}
	c.JSON(http.StatusOK, tires)
}

func AddTire(c *gin.Context) {
	var tire models.Tire
	if err := c.ShouldBindJSON(&tire); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := db.DB.Create(&tire).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add tire"})
		return
	}

	c.JSON(http.StatusCreated, tire)
}

func UpdateTire(c *gin.Context) {
	id := c.Param("id")
	var tire models.Tire
	if err := db.DB.First(&tire, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tire not found"})
		return
	}

	if err := c.ShouldBindJSON(&tire); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := db.DB.Save(&tire).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update tire"})
		return
	}

	c.JSON(http.StatusOK, tire)
}

func DeleteTire(c *gin.Context) {
	id := c.Param("id")
	var tire models.Tire
	if err := db.DB.First(&tire, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tire not found"})
		return
	}

	if err := db.DB.Delete(&tire).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tire"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tire deleted"})
}
