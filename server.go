package main

import (
	"log"
	"net/http"

	"github.com/byuoitav/hateoas"
	"github.com/byuoitav/sony-control-microservice/handlers"
	"github.com/byuoitav/wso2jwt"
	"github.com/jessemillar/health"
	"github.com/labstack/echo"
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

	// Use the `secure` routing group to require authentication
	secure := router.Group("", echo.WrapMiddleware(wso2jwt.ValidateJWT))

	router.GET("/", echo.WrapHandler(http.HandlerFunc(hateoas.RootResponse)))
	router.GET("/health", echo.WrapHandler(http.HandlerFunc(health.Check)))

	secure.GET("/:address/list/commands", handlers.GetCommands)
	secure.GET("/:address/command/:command", handlers.SendCommand)
	secure.GET("/:address/command/:command/count/:count", handlers.FloodCommand)

	secure.GET("/:address/power/on", handlers.PowerOn)
	secure.GET("/:address/power/standby", handlers.Standby)
	secure.GET("/:address/input/:port", handlers.SwitchInput)
	secure.GET("/:address/volume/set/:difference", handlers.SetVolume)
	secure.GET("/:address/volume/calibrate/:default", handlers.CalibrateVolume)
	secure.GET("/:address/volume/up", handlers.VolumeUp)
	secure.GET("/:address/volume/down", handlers.VolumeDown)
	secure.GET("/:address/volume/mute", handlers.VolumeMute)
	secure.GET("/:address/volume/unmute", handlers.VolumeUnmute)
	secure.GET("/:address/display/blank", handlers.BlankDisplay)
	secure.GET("/:address/display/unblank", handlers.UnblankDisplay)

	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 1024 * 10,
	}

	router.StartServer(&server)
}
