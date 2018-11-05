package main

import (
	"net/http"

	"github.com/byuoitav/common/status"
	"github.com/byuoitav/hateoas"
	"github.com/byuoitav/sony-control-microservice/handlers"
	"github.com/jessemillar/health"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	port := ":8007"
	router := echo.New()

	router.Pre(middleware.RemoveTrailingSlash())
	router.Use(middleware.CORS())

	// Use the `secure` routing group to require authentication
	//secure := router.Group("", echo.WrapMiddleware(authmiddleware.Authenticate))

	router.GET("/", echo.WrapHandler(http.HandlerFunc(hateoas.RootResponse)))
	router.GET("/health", echo.WrapHandler(http.HandlerFunc(health.Check)))
	router.GET("/status", status.DefaultStatusHandler)
	router.GET("/mstatus", status.DefaultStatusHandler)

	//functionality endpoints
	router.GET("/:address/power/on", handlers.PowerOn)
	router.GET("/:address/power/standby", handlers.Standby)
	router.GET("/:address/input/:port", handlers.SwitchInput)
	router.GET("/:address/volume/set/:value", handlers.SetVolume)
	router.GET("/:address/volume/mute", handlers.VolumeMute)
	router.GET("/:address/volume/unmute", handlers.VolumeUnmute)
	router.GET("/:address/display/blank", handlers.BlankDisplay)
	router.GET("/:address/display/unblank", handlers.UnblankDisplay)

	//status endpoints
	router.GET("/:address/power/status", handlers.GetPower)
	router.GET("/:address/input/current", handlers.GetInput)
	router.GET("/:address/input/list", handlers.GetInputList)
	router.GET("/:address/volume/level", handlers.GetVolume)
	router.GET("/:address/volume/mute/status", handlers.GetMute)
	router.GET("/:address/display/status", handlers.GetBlank)

	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 1024 * 10,
	}

	router.StartServer(&server)
}
