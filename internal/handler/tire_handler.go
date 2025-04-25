package handler

import (
	"fmt"
	"github.com/asylzhan/go-asylzhan-project/internal/models"
	"github.com/asylzhan/go-asylzhan-project/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TireHandler struct {
	service service.TireServiceInterface
}

func NewTireHandler(service service.TireServiceInterface) *TireHandler {
	return &TireHandler{service}
}

func (h *TireHandler) GetAll(c *gin.Context) {
	tires, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tires)
}

func (h *TireHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	tire, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tire)
}

func (h *TireHandler) Create(c *gin.Context) {
	var tire models.Tire
	if err := c.ShouldBindJSON(&tire); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}

	fmt.Println("Creating tire:", tire)

	if err := h.service.Create(&tire); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error", "details": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, tire)
}

func (h *TireHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var tire models.Tire
	if err := c.ShouldBindJSON(&tire); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tire.ID = uint(id)
	if err := h.service.Update(&tire); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tire)
}

func (h *TireHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
