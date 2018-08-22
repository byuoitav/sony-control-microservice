package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/byuoitav/sony-control-microservice/helpers"

	"github.com/byuoitav/common/status"
	"github.com/labstack/echo"
)

func PowerOn(context echo.Context) error {
	log.Printf("Powering on %s...", context.Param("address"))

	err := helpers.SetPower(context.Param("address"), true)
	if err != nil {
		log.Printf("Error: %v", err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	log.Printf("Done.")
	return context.JSON(http.StatusOK, status.Power{"on"})
}

func Standby(context echo.Context) error {
	log.Printf("Powering off %s...", context.Param("address"))

	err := helpers.SetPower(context.Param("address"), false)
	if err != nil {
		log.Printf("Error: %v", err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	log.Printf("Done.")
	return context.JSON(http.StatusOK, status.Power{"standby"})
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
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	log.Printf("Done.")
	return context.JSON(http.StatusOK, status.Input{port})
}

func SetVolume(context echo.Context) error {
	address := context.Param("address")
	value := context.Param("value")

	volume, err := strconv.Atoi(value)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	} else if volume > 100 || volume < 0 {
		return context.JSON(http.StatusBadRequest, "Error: volume must be a value from 0 to 100!")
	}

	log.Printf("Setting volume for %s to %v...", address, value)

	params := make(map[string]interface{})
	params["target"] = "speaker"
	params["volume"] = value

	err = helpers.BuildAndSendPayload(address, "audio", "setAudioVolume", params)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	//do the same for the headphone
	params = make(map[string]interface{})
	params["target"] = "headphone"
	params["volume"] = value

	err = helpers.BuildAndSendPayload(address, "audio", "setAudioVolume", params)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	log.Printf("Done.")
	return context.JSON(http.StatusOK, status.Volume{volume})
}

func VolumeUnmute(context echo.Context) error {
	address := context.Param("address")
	log.Printf("Unmuting %s...", address)

	err := setMute(address, false, 4)
	if err != nil {
		log.Printf("Error: %v", err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	log.Printf("Done.")
	return context.JSON(http.StatusOK, status.Mute{false})
}

func setMute(address string, status bool, retryCount int) error {
	params := make(map[string]interface{})
	params["status"] = status

	initCount := retryCount

	for retryCount >= 0 {
		err := helpers.BuildAndSendPayload(address, "audio", "setAudioMute", params)
		if err != nil {
			return err
		}
		//we need to validate that it was actually muted
		postStatus, err := helpers.GetMute(address)
		if err != nil {
			return err
		}

		if postStatus.Muted == status {
			return nil
		}
		retryCount--

		//wait for a short time
		time.Sleep(10 * time.Millisecond)
	}
	return errors.New(fmt.Sprintf("Attempted to set mute status %v times, could not", initCount+1))
}

func VolumeMute(context echo.Context) error {
	log.Printf("Muting %s...", context.Param("address"))

	err := setMute(context.Param("address"), true, 4)
	if err != nil {
		log.Printf("Error: %v", err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	log.Printf("Done.")
	return context.JSON(http.StatusOK, status.Mute{true})
}

func BlankDisplay(context echo.Context) error {
	params := make(map[string]interface{})
	params["mode"] = "pictureOff"

	err := helpers.BuildAndSendPayload(context.Param("address"), "system", "setPowerSavingMode", params)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, status.Blanked{true})

}

func UnblankDisplay(context echo.Context) error {
	params := make(map[string]interface{})
	params["mode"] = "off"

	err := helpers.BuildAndSendPayload(context.Param("address"), "system", "setPowerSavingMode", params)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, status.Blanked{false})
}

func GetVolume(context echo.Context) error {
	response, err := helpers.GetVolume(context.Param("address"))
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func GetInput(context echo.Context) error {

	response, err := helpers.GetInput(context.Param("address"))
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func GetInputList(context echo.Context) error {
	return nil
}

func GetMute(context echo.Context) error {
	response, err := helpers.GetMute(context.Param("address"))
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func GetBlank(context echo.Context) error {
	response, err := helpers.GetBlanked(context.Param("address"))
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, response)
	return nil
}

////////////////////////////////////////////////////////////////////
// New Web API Calls Below
////////////////////////////////////////////////////////////////////

// GetPowerAPI is the API call retreiving power status
func GetPowerAPI(context echo.Context) error {
	log.Printf("API - Getting power status of %s...", context.Param("address"))

	response, err := helpers.GetPowerStatus(context.Param("address"))
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
