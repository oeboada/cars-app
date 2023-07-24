// Package repository implements a facade to in-memory database.
// Through implementation of the CarsRepositoryInterface methods
package repository

import (
	"fmt"

	"github.com/oeboada/cars-app/db"
	"github.com/oeboada/cars-app/dto"

	"github.com/segmentio/ksuid"
)

// CarsRepositoryInterface is the interface of Cars repository layer.
// Contains definition of methods to manage the database representation of
// Cars entity.
type CarsRepositoryInterface interface {
	GetNewId() string
	Create(c dto.Car) dto.Car
	Update(c dto.Car)
	FindById(id string) (bool, dto.Car)
	FindAll() []dto.Car
}

// CarsRepository  represents the database repository layer of
// Cars entity. It's the registry of Cars entity records.
type CarsRepository struct {
	Db *db.Db // available database
}

// NewCarsRepository initializes repository of Car entities.
func NewCarsRepository(db *db.Db) CarsRepositoryInterface {
	return &CarsRepository{
		db,
	}
}

// Create implements insert action of Car entity.
func (r *CarsRepository) Create(c dto.Car) dto.Car {
	return r.Db.Create(c)
}

// UpdateById implements update action of Car entity by id.
func (r *CarsRepository) Update(c dto.Car) {
	r.Db.Update(c)
}

// FindById implements exist action of Cars entity by id.
// Returns true if successful and false otherwise.
func (r *CarsRepository) FindById(id string) (bool, dto.Car) {
	return r.Db.FindById(id)
}

// FindAll returns a list of all Car entities.
func (r *CarsRepository) FindAll() []dto.Car {
	return r.Db.FindAll()
}

// GetNewId() generate a new id
func (s *CarsRepository) GetNewId() string {
	return fmt.Sprintf("%s", ksuid.New())
}
