package main

import (
	"fmt"

	"github.com/byuoitav/hateoas"
	"github.com/byuoitav/sony-control-microservice/handlers"
	"github.com/byuoitav/wso2jwt"
	"github.com/jessemillar/health"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
)

func main() {
	err := hateoas.Load("https://raw.githubusercontent.com/byuoitav/sony-control-microservice/master/swagger.json")
	if err != nil {
		fmt.Println("Could not load swagger.json file. Error: " + err.Error())
		panic(err)
	}

	port := ":8007"
	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())

	router.Get("/", hateoas.RootResponse)
	router.Get("/health", health.Check)

	router.Get("/command/:address", handlers.GetCommands, wso2jwt.ValidateJWT())

	router.Post("/command", handlers.SendCommand, wso2jwt.ValidateJWT())

	fmt.Printf("Sony Control microservice is listening on %s\n", port)
	server := fasthttp.New(port)
	server.ReadBufferSize = 1024 * 10 // Needed to interface properly with WSO2
	router.Run(server)
}
