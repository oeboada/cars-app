// Package router defines request urls of microservice api
package router

import (
	"github.com/oeboada/cars-app/controller"
	"net/http"

	"github.com/swaggest/rest/nethttp"
	"github.com/swaggest/rest/web"
)

// Router represents the router layer.
type CarsRouter struct {
	server         *web.Service                       // embedded a http server
	carsController controller.CarsControllerInterface // controller layer
}

// NewRouter initializes router layer
func NewCarsRouter(
	server *web.Service,
	carsController controller.CarsControllerInterface,
) *CarsRouter {
	return &CarsRouter{
		server,
		carsController,
	}
}

// Init implements request urls definition
func (r *CarsRouter) Init() {

	r.server.Get("/cars", r.carsController.GetAllCars())
	r.server.Get("/cars/{id}", r.carsController.GetCarById())
	r.server.Put("/cars/{id}", r.carsController.PutCar())
	r.server.Post("/cars", r.carsController.PostCar(), nethttp.SuccessStatus(http.StatusCreated))
}
