// Package db compraises in-memory database definition of
// cars-app microservice.
package db

import (
	"github.com/oeboada/cars-app/dto"
	"golang.org/x/exp/maps"
)

var cars = map[string]dto.Car{
	"ZJkWuZy378gsd8lrMXdFTYLvbo0": {Id: "ZJkWuZy378gsd8lrMXdFTYLvbo0", Make: "Ford", Model: "F10", Package: "Base", Color: "Silver", Year: 2010, Category: "Truck", Mileage: 120123, Price: 1999900},
	"ZJkWuXxgIulNqb8AU1gmUo9HZ9K": {Id: "ZJkWuXxgIulNqb8AU1gmUo9HZ9K", Make: "Toyota", Model: "Camry", Package: "SE", Color: "White", Year: 2019, Category: "Segan", Mileage: 3999, Price: 2899000},
	"ZJkWub1uNXDRxEN2PIDBuAxBKqH": {Id: "ZJkWub1uNXDRxEN2PIDBuAxBKqH", Make: "Toyota", Model: "Rav4", Package: "XSE", Color: "Red", Year: 2018, Category: "SUV", Mileage: 24001, Price: 2275000},
	"ZJkWuXveyGAIVGj0BCkYQ9Q3AGi": {Id: "ZJkWuXveyGAIVGj0BCkYQ9Q3AGi", Make: "Ford", Model: "Bronco", Package: "Badlants", Color: "Bumt Orange", Year: 2022, Category: "SUV", Mileage: 1, Price: 4499000},
}

type DbInterface interface {
	FindAll() []dto.Car
	Create(c dto.Car) dto.Car
	Update(c dto.Car)
	FindById(id string) (bool, dto.Car)
}

// Db struct represents the Cars database.
type Db struct {
	Data map[string]dto.Car // Data
}

// NewDb initializes Db.
func NewDb() *Db {
	return &Db{
		cars,
	}
}
func (db *Db) FindAll() []dto.Car {
	return maps.Values(db.Data)
}
func (db *Db) Create(c dto.Car) dto.Car {
	db.Data[c.Id] = c
	return db.Data[c.Id]
}
func (db *Db) Update(c dto.Car) {
	db.Data[c.Id] = c
}

func (db *Db) FindById(id string) (bool, dto.Car) {

	if value, isPresent := db.Data[id]; isPresent {
		return true, value
	}

	return false, dto.Car{}
}
