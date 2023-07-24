# Cars App

## Project overview

### General concept

Cars App is a API Restful to process Cars developed under standard libraries without any http framework. However, 
this app comprises a clean architecture support that allows to manage Restful error codes simpler and generate
an Open API documentation easier. The persistence operations are implemented with a simulated in-memory database.

### Data definition structure

type Car struct {
```
Id       string  `json:"id" required:"true" minLength:"0" description:"Id is a unique string that determines car."`
```
```
Make     string  `json:"make" required:"true" description:"Make of the car."`
```
```
Model    string  `json:"model" required:"true" description:"Model of the car."`
```
```
Package  string  `json:"package" required:"true" description:"Package of the car."`
```
```
Color    string  `json:"color" required:"true" description:"Color of the car."`
```
```
Year     uint16  `json:"year" required:"true" minimum:"1800" maximum:"9999" description:"Year of the car."`
```
```
Category string  `json:"category" required:"true" description:"Category of the car."`
```
```
Mileage  float64 `json:"milage" minimum:"0" description:"Mileage in miles."`
```
```
Price    float64 `json:"price" minimum:"0" description:"Price in USD cents."`
```
}

### Data preloaded by default

```
{Id: "ZJkWuZy378gsd8lrMXdFTYLvbo0", Make: "Ford", Model: "F10", Package: "Base", Color: "Silver", Year: 2010, Category: "Truck", Mileage: 120123, Price: 1999900},
{Id: "ZJkWuXxgIulNqb8AU1gmUo9HZ9K", Make: "Toyota", Model: "Camry", Package: "SE", Color: "White", Year: 2019, Category: "Segan", Mileage: 3999, Price: 2899000},
{Id: "ZJkWub1uNXDRxEN2PIDBuAxBKqH", Make: "Toyota", Model: "Rav4", Package: "XSE", Color: "Red", Year: 2018, Category: "SUV", Mileage: 24001, Price: 2275000},
{Id: "ZJkWuXveyGAIVGj0BCkYQ9Q3AGi", Make: "Ford", Model: "Bronco", Package: "Badlants", Color: "Bumt Orange", Year: 2022, Category: "SUV", Mileage: 1, Price: 4499000},
```

### Build and Run app

Located in the root directory of the application we carry out the following steps:

```
$ go build
```
```
$ go run main.go
```

### Running Unit Testing app

Located in the root directory of the application we carry out the following steps:

```
$ go test -cover -json ./...
```

### Open API Documentation

We can use the following url to present in the web browser the documentation in Open API format of the application that allows us to intuitively test the application. 

```
http://localhost:8580/docs#/
```

### TODOs

In order to give it a more production-oriented finish, we must complete the following:

1. Add OAuth2 security and OpenID Connect if you want restricted access to API methods
2. Change simulated in-memory database to SQL database
3. Design tests to cover the different use cases that are of interest
4. Create objects that represent business models that will be exposed outside the application. While business entities represent the data in the database.
5. Put the source code in a git repo
6. Use advanced libraries and frameworks that facilitate and organize development
7. Define deployment files to allow application management with docker and kubernetes
8. Use external logging and configuration

