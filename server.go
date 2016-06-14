package main

import (
	"fmt"

	"github.com/byuoitav/hateoas"
	"github.com/byuoitav/sony-control-microservice/controllers"
	"github.com/byuoitav/sony-control-microservice/packages/wso2jwt"
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
	router.Use(wso2jwt.ValidateJWT())
	// 	router.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 		SigningMethod: "RS256",
	// 		SigningKey: []byte(`-----BEGIN PUBLIC KEY-----
	// MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDdlatRjRjogo3WojgGHFHYLugdUWAY9iR3fy4arWNA1KoS8kVw33cJibXr8bvwUAUparCwlvdbH6dvEOfou0/gCFQsHUfQrSDv+MuSUMAe8jzKE4qW+jK+xQU9a03GUnKHkkle+Q0pX/g6jXZ7r1/xAK5Do2kQ+X5xK9cipRgEKwIDAQAB
	// -----END PUBLIC KEY-----`),
	// 	}))

	// GET requests
	router.Get("/health", health.Check)

	router.Get("/", hateoas.RootResponse)
	router.Get("/command/:address", controllers.GetCommands)

	router.Post("/command", controllers.SendCommand)

	fmt.Printf("Sony Control microservice is listening on %s\n", port)
	router.Run(fasthttp.New(port))
}
