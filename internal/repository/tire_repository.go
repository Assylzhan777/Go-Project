package repository

import (
	"github.com/asylzhan/go-asylzhan-project/internal/models"
	"gorm.io/gorm"
)

type TireRepositoryInterface interface {
	Create(tire *models.Tire) error
	GetAll() ([]models.Tire, error)
	GetByID(id uint) (*models.Tire, error)
	Update(tire *models.Tire) error
	Delete(id uint) error
}

type TireRepository struct {
	db *gorm.DB
}

func NewTireRepository(db *gorm.DB) *TireRepository {
	return &TireRepository{db}
}

func (r *TireRepository) Create(tire *models.Tire) error {
	if err := r.db.Create(tire).Error; err != nil {
		return err
	}
	return nil
}

func (r *TireRepository) GetAll() ([]models.Tire, error) {
	var tires []models.Tire
	if err := r.db.Find(&tires).Error; err != nil {
		return nil, err
	}
	return tires, nil
}

func (r *TireRepository) GetByID(id uint) (*models.Tire, error) {
	var tire models.Tire
	if err := r.db.First(&tire, id).Error; err != nil {
		return nil, err
	}
	return &tire, nil
}

func (r *TireRepository) Update(tire *models.Tire) error {
	if err := r.db.Save(tire).Error; err != nil {
		return err
	}
	return nil
}

func (r *TireRepository) Delete(id uint) error {
	if err := r.db.Delete(&models.Tire{}, id).Error; err != nil {
		return err
	}
	return nil
}
