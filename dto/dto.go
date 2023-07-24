// Package dto defines domain objects of
// cars-app microservice.
package dto

type Car struct {
	Id       string  `json:"id" required:"true" minLength:"0" description:"Id is a unique string that determines car."`
	Make     string  `json:"make" required:"true" description:"Make of the car."`
	Model    string  `json:"model" required:"true" description:"Model of the car."`
	Package  string  `json:"package" required:"true" description:"Package of the car."`
	Color    string  `json:"color" required:"true" description:"Color of the car."`
	Year     uint16  `json:"year" required:"true" minimum:"1800" maximum:"9999" description:"Year of the car."`
	Category string  `json:"category" required:"true" description:"Category of the car."`
	Mileage  float64 `json:"milage" minimum:"0" description:"Mileage in miles."`
	Price    float64 `json:"price" minimum:"0" description:"Price in USD cents."`
}
