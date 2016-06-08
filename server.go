package main

import (
	"fmt"

	"github.com/byuoitav/hateoas"
	"github.com/jessemillar/health"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
)

func main() {
	err := hateoas.Load("https://raw.githubusercontent.com/byuoitav/sony-control/master/swagger.yml")
	if err != nil {
		fmt.Println("Could not load Swagger file")
		panic(err)
	}

	port := ":8006"
	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())

	// GET requests
	router.Get("/", hateoas.RootResponse)

	router.Get("/health", health.Check)

	fmt.Printf("Sony Control is listening on %s\n", port)
	router.Run(fasthttp.New(port))
}
