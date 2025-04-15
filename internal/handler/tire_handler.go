package handler

import (
	"net/http"
	"strconv"

	"github.com/asylzhan/go-asylzhan-project/internal/models"
	"github.com/asylzhan/go-asylzhan-project/internal/service"
	"github.com/gin-gonic/gin"
)

type TireHandler struct {
	service service.TireService
}

func NewTireHandler(service service.TireService) *TireHandler {
	return &TireHandler{service}
}

func (h *TireHandler) Create(c *gin.Context) {
	var tire models.Tire
	if err := c.ShouldBindJSON(&tire); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.Create(&tire); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tire"})
		return
	}
	c.JSON(http.StatusCreated, tire)
}

func (h *TireHandler) GetAll(c *gin.Context) {
	tires, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tires"})
		return
	}
	c.JSON(http.StatusOK, tires)
}

func (h *TireHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	tire, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tire not found"})
		return
	}
	c.JSON(http.StatusOK, tire)
}

func (h *TireHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var tire models.Tire
	if err := c.ShouldBindJSON(&tire); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tire.ID = uint(id)
	if err := h.service.Update(&tire); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update tire"})
		return
	}
	c.JSON(http.StatusOK, tire)
}

func (h *TireHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tire"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tire deleted"})
}
