package service

import (
	"github.com/asylzhan/go-asylzhan-project/internal/models"
	"github.com/asylzhan/go-asylzhan-project/internal/repository"
)

type TireService interface {
	Create(tire *models.Tire) error
	GetAll() ([]models.Tire, error)
	GetByID(id uint) (*models.Tire, error)
	Update(tire *models.Tire) error
	Delete(id uint) error
}

type tireService struct {
	repo repository.TireRepository
}

func NewTireService(repo repository.TireRepository) TireService {
	return &tireService{repo}
}

func (s *tireService) Create(tire *models.Tire) error {
	return s.repo.Create(tire)
}

func (s *tireService) GetAll() ([]models.Tire, error) {
	return s.repo.GetAll()
}

func (s *tireService) GetByID(id uint) (*models.Tire, error) {
	return s.repo.GetByID(id)
}

func (s *tireService) Update(tire *models.Tire) error {
	return s.repo.Update(tire)
}

func (s *tireService) Delete(id uint) error {
	return s.repo.Delete(id)
}
