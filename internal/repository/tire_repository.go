package repository

import (
	"github.com/asylzhan/go-asylzhan-project/internal/models"
	"gorm.io/gorm"
)

type TireRepository interface {
	Create(tire *models.Tire) error
	GetAll() ([]models.Tire, error)
	GetByID(id uint) (*models.Tire, error)
	Update(tire *models.Tire) error
	Delete(id uint) error
}

type tireRepository struct {
	db *gorm.DB
}

func NewTireRepository(db *gorm.DB) TireRepository {
	return &tireRepository{db}
}

func (r *tireRepository) Create(tire *models.Tire) error {
	return r.db.Create(tire).Error
}

func (r *tireRepository) GetAll() ([]models.Tire, error) {
	var tires []models.Tire
	err := r.db.Find(&tires).Error
	return tires, err
}

func (r *tireRepository) GetByID(id uint) (*models.Tire, error) {
	var tire models.Tire
	err := r.db.First(&tire, id).Error
	return &tire, err
}

func (r *tireRepository) Update(tire *models.Tire) error {
	return r.db.Save(tire).Error
}

func (r *tireRepository) Delete(id uint) error {
	return r.db.Delete(&models.Tire{}, id).Error
}
