package service

import (
	"github.com/oeboada/cars-app/db"
	"github.com/oeboada/cars-app/repository"
	"github.com/oeboada/cars-app/testutils"

	"testing"
)

func TestCreate(t *testing.T) {
	nameTest := "TestCreate"
	db := db.NewDb()

	carsRepository := repository.NewCarsRepository(db)
	carsService := NewCarsService(carsRepository)

	carInstance := testutils.GetCarInstance()
	carInstance.Id = carsService.GetNewId()

	newCarInstance := carsService.Create(carInstance)

	if newCarInstance.Id == carInstance.Id {
		t.Logf("%s Successful, car created", nameTest)
	} else {
		t.Logf("%s Failed, car not created", nameTest)
	}
}

func TestUpdate(t *testing.T) {
	nameTest := "TestUpdate"
	db := db.NewDb()

	carsRepository := repository.NewCarsRepository(db)
	carsService := NewCarsService(carsRepository)

	carInstance := testutils.GetCarInstance()
	carInstance.Id = carsService.GetNewId()

	newCarInstance := carsService.Create(carInstance)

	colorBefore := newCarInstance.Color

	newCarInstance.Color = "Yellow"
	carsService.Update(newCarInstance)

	if ok, existingCar := carsService.GetById(newCarInstance.Id); ok {
		if existingCar.Color != colorBefore {
			t.Logf("%s Successful, car updated ", nameTest)
		} else {
			t.Logf("%s Failed, car not updated", nameTest)
		}
	} else {
		t.Logf("%s Failed, car id not found", nameTest)
	}
}

func TestFindById(t *testing.T) {
	nameTest := "TestFindById"
	db := db.NewDb()

	carsRepository := repository.NewCarsRepository(db)
	carsService := NewCarsService(carsRepository)

	carInstance := testutils.GetCarInstance()
	carInstance.Id = carsService.GetNewId()

	newCarInstance := carsService.Create(carInstance)

	if ok, _ := carsService.GetById(newCarInstance.Id); ok {
		t.Logf("%s Successful, car id found", nameTest)
	} else {
		t.Logf("%s Failed, car id not found", nameTest)
	}
}

func TestFindById_ErrorIdNotFound(t *testing.T) {
	nameTest := "TestFindById_ErrorIdNotFound"
	db := db.NewDb()

	carsRepository := repository.NewCarsRepository(db)
	carsService := NewCarsService(carsRepository)

	carInstance := testutils.GetCarInstance()
	carInstance.Id = carsService.GetNewId()

	carsService.Create(carInstance)

	if ok, _ := carsService.GetById(""); ok {
		t.Logf("%s Failed, car id found", nameTest)
	} else {
		t.Logf("%s Successful, car id not found", nameTest)
	}
}

func TestFindAll(t *testing.T) {
	nameTest := "TestFindAll"
	db := db.NewDb()

	carsRepository := repository.NewCarsRepository(db)
	carsService := NewCarsService(carsRepository)

	cars := carsService.GetAll()

	if len(cars) > 0 {
		t.Logf("%s Successful, all car list found", nameTest)
	} else {
		t.Logf("%s Failed, car list is empty", nameTest)
	}
}
