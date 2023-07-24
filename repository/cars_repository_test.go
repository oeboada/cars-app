package repository

import (
	"github.com/oeboada/cars-app/db"
	"github.com/oeboada/cars-app/testutils"

	"testing"
)

func TestCreate(t *testing.T) {
	nameTest := "TestCreate"
	db := db.NewDb()

	carsRepository := NewCarsRepository(db)

	carInstance := testutils.GetCarInstance()
	carInstance.Id = carsRepository.GetNewId()

	newCarInstance := carsRepository.Create(carInstance)

	if newCarInstance.Id == carInstance.Id {
		t.Logf("%s Successful, car created", nameTest)
	} else {
		t.Logf("%s Failed, car not created", nameTest)
	}
}

func TestUpdate(t *testing.T) {
	nameTest := "TestUpdate"
	db := db.NewDb()

	carsRepository := NewCarsRepository(db)

	carInstance := testutils.GetCarInstance()
	carInstance.Id = carsRepository.GetNewId()

	newCarInstance := carsRepository.Create(carInstance)

	colorBefore := newCarInstance.Color

	newCarInstance.Color = "Yellow"
	carsRepository.Update(newCarInstance)

	if ok, existingCar := carsRepository.FindById(newCarInstance.Id); ok {
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

	carsRepository := NewCarsRepository(db)

	carInstance := testutils.GetCarInstance()
	carInstance.Id = carsRepository.GetNewId()

	newCarInstance := carsRepository.Create(carInstance)

	if ok, _ := carsRepository.FindById(newCarInstance.Id); ok {
		t.Logf("%s Successful, car id found", nameTest)
	} else {
		t.Logf("%s Failed, car id not found", nameTest)
	}
}

func TestFindById_ErrorIdNotFound(t *testing.T) {
	nameTest := "TestFindById_ErrorIdNotFound"
	db := db.NewDb()

	carsRepository := NewCarsRepository(db)

	carInstance := testutils.GetCarInstance()
	carInstance.Id = carsRepository.GetNewId()

	carsRepository.Create(carInstance)

	if ok, _ := carsRepository.FindById(""); ok {
		t.Logf("%s Failed, car id found", nameTest)
	} else {
		t.Logf("%s Successful, car id not found", nameTest)
	}
}

func TestFindAll(t *testing.T) {
	nameTest := "TestFindAll"
	db := db.NewDb()

	carsRepository := NewCarsRepository(db)

	cars := carsRepository.FindAll()

	if len(cars) > 0 {
		t.Logf("%s Successful, all car list found", nameTest)
	} else {
		t.Logf("%s Failed, car list is empty", nameTest)
	}
}
