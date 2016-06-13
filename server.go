package main

import (
	"fmt"

	"github.com/byuoitav/hateoas"
	"github.com/byuoitav/sony-control-microservice/controllers"
	"github.com/jessemillar/health"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
)

func main() {
	err := hateoas.Load("https://raw.githubusercontent.com/byuoitav/sony-control-microservice/master/swagger.yml")
	if err != nil {
		fmt.Println("Could not load swagger.yaml file. Error: " + err.Error())
		panic(err)
	}

	port := ":8007"
	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())

	// GET requests
	router.Get("/health", health.Check)

	router.Get("/", hateoas.RootResponse)
	router.Get("/command/:address", controllers.GetCommands)

	router.Post("/command", controllers.SendCommand)

	fmt.Printf("Sony Control microservice is listening on %s\n", port)
	router.Run(fasthttp.New(port))
}
