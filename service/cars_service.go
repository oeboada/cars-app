// Package service implements business logic layer.
// Through implementation of the CarsServiceInterface methods,
// it is possible to define business logic it also invokes repository
// layer.
package service

import (
	"github.com/oeboada/cars-app/dto"
	"github.com/oeboada/cars-app/repository"
)

// CarsServiceInterface is the interface of Cars service layer. Contains definition of
// methods to manage the business logic of Cars models.
type CarsServiceInterface interface {
	GetNewId() string
	Create(c dto.Car) dto.Car
	Update(c dto.Car)
	GetById(id string) (bool, dto.Car)
	GetAll() []dto.Car
}

// CarsService represents the Cars service layer.
type CarsService struct {
	carsRepository repository.CarsRepositoryInterface // Cars repository interface
}

// NewCarsService initializes Cars service layer.
func NewCarsService(carsRepository repository.CarsRepositoryInterface) CarsServiceInterface {
	return &CarsService{
		carsRepository,
	}
}

// GetNewId() generate a new id
func (s *CarsService) GetNewId() string {
	return s.carsRepository.GetNewId()
}

func (s *CarsService) GetAll() []dto.Car {
	return s.carsRepository.FindAll()
}

func (s *CarsService) Create(c dto.Car) dto.Car {
	return s.carsRepository.Create(c)
}

func (s *CarsService) Update(c dto.Car) {
	s.carsRepository.Update(c)
}

func (s *CarsService) GetById(id string) (bool, dto.Car) {
	return s.carsRepository.FindById(id)
}
