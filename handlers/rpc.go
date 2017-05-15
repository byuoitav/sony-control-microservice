package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/byuoitav/sony-control-microservice/helpers"
	"github.com/labstack/echo"
)

func PowerOn(context echo.Context) error {
	log.Printf("Powering on %s...", context.Param("address"))

	err := helpers.SetPower(context.Param("address"), true)
	if err != nil {
		log.Printf("Error: %v", err.Error())
		return context.JSONBlob(http.StatusInternalServerError, []byte(err.Error()))
	}

	log.Printf("Done.")
	return nil
}

func Standby(context echo.Context) error {
	log.Printf("Powering off %s...", context.Param("address"))

	err := helpers.SetPower(context.Param("address"), false)
	if err != nil {
		log.Printf("Error: %v", err.Error())
		return context.JSONBlob(http.StatusInternalServerError, []byte(err.Error()))
	}

	log.Printf("Done.")
	return nil
}

func GetPower(context echo.Context) error {
	log.Printf("Getting power status of %s...", context.Param("address"))

	response, err := helpers.GetPower(context.Param("address"))
	if err != nil {
		return context.JSONBlob(http.StatusInternalServerError, []byte(err.Error()))
	}

	return context.JSON(http.StatusOK, response)
}

func SwitchInput(context echo.Context) error {
	log.Printf("Switching input for %s to %s ...", context.Param("address"), context.Param("port"))
	address := context.Param("address")
	port := context.Param("port")

	splitPort := strings.Split(port, "!")

	params := make(map[string]interface{})
	params["uri"] = fmt.Sprintf("extInput:%s?port=%s", splitPort[0], splitPort[1])

	err := helpers.BuildAndSendPayload(address, "avContent", "setPlayContent", params)
	if err != nil {
		return context.JSONBlob(http.StatusInternalServerError, []byte(err.Error()))
	}

	log.Printf("Done.")
	return nil
}

func SetVolume(context echo.Context) error {
	address := context.Param("address")
	value := context.Param("value")

	log.Printf("Setting volume for %s to %v...", address, value)

	params := make(map[string]interface{})
	params["target"] = "speaker"
	params["volume"] = value

	err := helpers.BuildAndSendPayload(address, "audio", "setAudioVolume", params)
	if err != nil {
		return context.JSONBlob(http.StatusInternalServerError, []byte(err.Error()))
	}

	log.Printf("Done.")
	return nil
}

func VolumeUnmute(context echo.Context) error {
	address := context.Param("address")
	log.Printf("Unmuting %s...", address)

	err := setMute(address, false)
	if err != nil {
		log.Printf("Error: %v", err.Error())
		return context.JSONBlob(http.StatusInternalServerError, []byte(err.Error()))
	}

	log.Printf("Done.")
	return nil
}

func setMute(address string, status bool) error {
	params := make(map[string]interface{})
	params["status"] = status

	return helpers.BuildAndSendPayload(address, "audio", "setAudioMute", params)
}

func VolumeMute(context echo.Context) error {
	log.Printf("Muting %s...", context.Param("address"))

	err := setMute(context.Param("address"), true)
	if err != nil {
		log.Printf("Error: %v", err.Error())
		return context.JSONBlob(http.StatusInternalServerError, []byte(err.Error()))
	}

	log.Printf("Done.")
	return nil
}

func BlankDisplay(context echo.Context) error {
	return Standby(context)
}

func UnblankDisplay(context echo.Context) error {
	return PowerOn(context)
}

func GetVolume(context echo.Context) error {
	response, err := helpers.GetVolume(context.Param("address"))
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func GetInput(context echo.Context) error {
	return nil
}

func GetInputList(context echo.Context) error {
	return nil
}

func GetMute(context echo.Context) error {
	return nil
}

func GetBlank(context echo.Context) error {
	return nil
}
