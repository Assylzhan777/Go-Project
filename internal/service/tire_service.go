package service

import (
	"github.com/asylzhan/go-asylzhan-project/internal/models"
	"github.com/asylzhan/go-asylzhan-project/internal/repository"
)

type TireServiceInterface interface {
	Create(tire *models.Tire) error
	GetAll() ([]models.Tire, error)
	GetByID(id uint) (*models.Tire, error)
	Update(tire *models.Tire) error
	Delete(id uint) error
}

type TireService struct {
	repo repository.TireRepositoryInterface
}

func NewTireService(repo repository.TireRepositoryInterface) *TireService {
	return &TireService{repo}
}

func (s *TireService) Create(tire *models.Tire) error {
	return s.repo.Create(tire)
}

func (s *TireService) GetAll() ([]models.Tire, error) {
	return s.repo.GetAll()
}

func (s *TireService) GetByID(id uint) (*models.Tire, error) {
	return s.repo.GetByID(id)
}

func (s *TireService) Update(tire *models.Tire) error {
	return s.repo.Update(tire)
}

func (s *TireService) Delete(id uint) error {
	return s.repo.Delete(id)
}
