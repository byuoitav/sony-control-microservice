package main

import (
	"net/http"

	"github.com/byuoitav/common"
	"github.com/byuoitav/sony-control-microservice/handlers"
)

func main() {
	port := ":8007"
	router := common.NewRouter()

	// functionality endpoints
	router.GET("/:address/power/on", handlers.PowerOn)
	router.GET("/:address/power/standby", handlers.Standby)
	router.GET("/:address/input/:port", handlers.SwitchInput)
	router.GET("/:address/volume/set/:value", handlers.SetVolume)
	router.GET("/:address/volume/mute", handlers.VolumeMute)
	router.GET("/:address/volume/unmute", handlers.VolumeUnmute)
	router.GET("/:address/display/blank", handlers.BlankDisplay)
	router.GET("/:address/display/unblank", handlers.UnblankDisplay)

	// status endpoints
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
