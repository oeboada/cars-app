// Package controller implements api layer for external clients.
// Through implementation of the CarsControllerInterface methods,
// it is possible to define validation and management of parameters,
// it also invokes service layer.
package controller

import (
	"context"
	"errors"
	"github.com/oeboada/cars-app/dto"
	"github.com/oeboada/cars-app/service"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"

	"github.com/Masterminds/log-go"
)

// CarsControllerInterface is the interface of Cars controller layer. Contains definition of
// methods to manage the microservice apis.
type CarsControllerInterface interface {
	GetAllCars() usecase.Interactor
	PostCar() usecase.Interactor
	PutCar() usecase.Interactor
	GetCarById() usecase.Interactor
}

// CarsController represents the Cars controller layer.
type CarsController struct {
	carsService service.CarsServiceInterface // Cars service interface
}

// NewCarsController initializes Cars controller layer.
func NewCarsController(carsService service.CarsServiceInterface) CarsControllerInterface {
	return &CarsController{
		carsService,
	}
}

func (ctr *CarsController) GetAllCars() usecase.Interactor {

	u := usecase.NewInteractor(func(ctx context.Context, _ struct{}, output *[]dto.Car) error {
		log.Info("GetAllCars() Started")
		result := ctr.carsService.GetAll()
		*output = result

		log.Infof("GetAllCars() Finished Successful, cars size %d", len(result))
		return nil
	})
	u.SetTags("Cars")

	return u
}

func (ctr *CarsController) PostCar() usecase.Interactor {
	u := usecase.NewInteractor(func(ctx context.Context, input dto.Car, output *dto.Car) error {
		log.Info("PostCar() Started")
		input.Id = ctr.carsService.GetNewId()

		// Check if car id is unique.
		if ok, _ := ctr.carsService.GetById(input.Id); ok {
			log.Infof("PostCar() Finished Successful, car id %s already exists", input.Id)
			return status.Wrap(errors.New("car id already exists"), status.AlreadyExists)
		}

		// Add the new car to the database.
		c := ctr.carsService.Create(input)

		*output = c

		log.Infof("PostCar() Finished Successful, created car id %s", c.Id)
		return nil
	})
	u.SetTags("Cars")
	u.SetExpectedErrors(status.AlreadyExists)

	return u
}

func (ctr *CarsController) PutCar() usecase.Interactor {
	type updateCarByIdInput struct {
		Id  string  `path:"id"  description:"Parameter in resource path, car id."`
		Car dto.Car `json:"body"  description:"Parameter in resource body, car."`
	}

	u := usecase.NewInteractor(func(ctx context.Context, input updateCarByIdInput, output *dto.Car) error {

		log.Info("PutCar() Started")

		// Check if id is unique.
		if ok, _ := ctr.carsService.GetById(input.Id); ok {
			// Update the car.
			ctr.carsService.Update(input.Car)

			*output = input.Car
			log.Infof("PutCar() Finished Successful, updated car id %s", input.Id)
			return nil
		}

		log.Infof("PutCar() Finished Successful, car id %s not found", input.Id)
		return status.Wrap(errors.New("car id not found"), status.NotFound)
	})
	u.SetTags("Cars")
	u.SetExpectedErrors(status.NotFound)

	return u
}

func (ctr *CarsController) GetCarById() usecase.Interactor {
	type getCarByIdInput struct {
		Id string `path:"id"  description:"Parameter in resource path, car id."`
	}

	u := usecase.NewInteractor(func(ctx context.Context, input getCarByIdInput, output *dto.Car) error {

		log.Info("GetCarById() Started")
		if ok, existingCar := ctr.carsService.GetById(input.Id); ok {
			*output = existingCar

			log.Infof("GetCarById() Finished Successful, car id %s found", input.Id)
			return nil
		}

		log.Infof("GetCarById() Finished Successful, car id %s not found", input.Id)
		return status.Wrap(errors.New("car id not found"), status.NotFound)
	})
	u.SetTags("Cars")
	u.SetExpectedErrors(status.NotFound)

	return u
}
