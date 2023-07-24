// Package appconfig implements a main routine that is the starting point of
// cars-app microservice.
package appconfig

import (
	"github.com/oeboada/cars-app/controller"
	"github.com/oeboada/cars-app/db"
	"github.com/oeboada/cars-app/repository"
	"github.com/oeboada/cars-app/router"
	"github.com/oeboada/cars-app/service"
	"log"
	"net/http"

	"github.com/swaggest/rest/web"
	"github.com/swaggest/swgui/v4emb"
)

// StartApp implements configuration and start-up of microservice.
func StartApp() {

	server := web.DefaultService()

	server.OpenAPI.Info.Title = "Cars API"
	server.OpenAPI.Info.WithDescription("This service provides API to manage cars.")
	server.OpenAPI.Info.Version = "v1.0.0"

	dbInstance := db.NewDb()
	carsRepository := repository.NewCarsRepository(dbInstance)
	carsService := service.NewCarsService(carsRepository)
	carsController := controller.NewCarsController(carsService)

	r := router.NewCarsRouter(server, carsController)
	r.Init()

	server.Docs("/docs", v4emb.New)

	log.Println("Starting service")
	if err := http.ListenAndServe("localhost:8580", server); err != nil {
		log.Fatal(err)
	}
}
