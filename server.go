package main

import (
	"net/http"

	"github.com/byuoitav/common"
	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/v2/auth"
	"github.com/byuoitav/sony-control-microservice/handlers"
)

func main() {
	port := ":8007"
	router := common.NewRouter()

	// functionality endpoints
	write := router.Group("", auth.AuthorizeRequest("write-state", "room", auth.LookupResourceFromAddress))
	write.GET("/:address/power/on", handlers.PowerOn)
	write.GET("/:address/power/standby", handlers.Standby)
	write.GET("/:address/input/:port", handlers.SwitchInput)
	write.GET("/:address/volume/set/:value", handlers.SetVolume)
	write.GET("/:address/volume/mute", handlers.VolumeMute)
	write.GET("/:address/volume/unmute", handlers.VolumeUnmute)
	write.GET("/:address/display/blank", handlers.BlankDisplay)
	write.GET("/:address/display/unblank", handlers.UnblankDisplay)

	// status endpoints
	read := router.Group("", auth.AuthorizeRequest("read-state", "room", auth.LookupResourceFromAddress))
	read.GET("/:address/power/status", handlers.GetPower)
	read.GET("/:address/input/current", handlers.GetInput)
	read.GET("/:address/input/list", handlers.GetInputList)
	read.GET("/:address/input/active", handlers.GetActiveInput)
	read.GET("/:address/volume/level", handlers.GetVolume)
	read.GET("/:address/volume/mute/status", handlers.GetMute)
	read.GET("/:address/display/status", handlers.GetBlank)
	read.GET("/:address/hardware", handlers.GetHardwareInfo)

	// log level endpoints
	router.PUT("/log-level/:level", log.SetLogLevel)
	router.GET("/log-level", log.GetLogLevel)

	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 1024 * 10,
	}

	router.StartServer(&server)
}
