package main

import (
	"log"

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
		log.Fatalln("Could not load swagger.json file. Error: " + err.Error())
	}

	port := ":8007"
	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())
	router.Use(middleware.CORS())

	router.Get("/", hateoas.RootResponse)
	router.Get("/health", health.Check)

	router.Get("/command/:address", handlers.GetCommands, wso2jwt.ValidateJWT())

	router.Post("/command", handlers.SendCommand, wso2jwt.ValidateJWT())

	log.Println("Sony Control microservice is listening on " + port)
	server := fasthttp.New(port)
	server.ReadBufferSize = 1024 * 10 // Needed to interface properly with WSO2
	router.Run(server)
}
