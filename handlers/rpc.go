package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/byuoitav/sony-control-microservice/helpers"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/status"
	"github.com/labstack/echo"
)

func PowerOn(context echo.Context) error {
	log.L.Infof("Powering on %s...", context.Param("address"))

	err := helpers.SetPower(context.Request().Context(), context.Param("address"), true)
	if err != nil {
		log.L.Debugf("Error: %v", err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	log.L.Debugf("Done.")
	return context.JSON(http.StatusOK, status.Power{Power: "on"})
}

func Standby(context echo.Context) error {
	log.L.Infof("Powering off %s...", context.Param("address"))

	err := helpers.SetPower(context.Request().Context(), context.Param("address"), false)
	if err != nil {
		log.L.Debugf("Error: %v", err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	log.L.Infof("Done.")
	return context.JSON(http.StatusOK, status.Power{Power: "standby"})
}

func GetPower(context echo.Context) error {
	log.L.Infof("Getting power status of %s...", context.Param("address"))

	response, err := helpers.GetPower(context.Request().Context(), context.Param("address"))
	if err != nil {
		return context.JSONBlob(http.StatusInternalServerError, []byte(err.Error()))
	}

	return context.JSON(http.StatusOK, response)
}

func SwitchInput(context echo.Context) error {
	log.L.Infof("Switching input for %s to %s ...", context.Param("address"), context.Param("port"))
	address := context.Param("address")
	port := context.Param("port")

	splitPort := strings.Split(port, "!")

	params := make(map[string]interface{})
	if len(splitPort) < 2 {
		return context.JSON(http.StatusBadRequest, fmt.Sprintf("ports configured incorrectly (should follow format \"hdmi!2\"): %s", port))
	}
	params["uri"] = fmt.Sprintf("extInput:%s?port=%s", splitPort[0], splitPort[1])

	err := helpers.BuildAndSendPayload(address, "avContent", "setPlayContent", params)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	log.L.Debugf("Done.")
	return context.JSON(http.StatusOK, status.Input{Input: port})
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

	log.L.Debugf("Setting volume for %s to %v...", address, value)

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

	log.L.Debugf("Done.")
	return context.JSON(http.StatusOK, status.Volume{Volume: volume})
}

func VolumeUnmute(context echo.Context) error {
	address := context.Param("address")
	log.L.Debugf("Unmuting %s...", address)

	err := setMute(context, address, false, 4)
	if err != nil {
		log.L.Debugf("Error: %v", err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	log.L.Debugf("Done.")
	return context.JSON(http.StatusOK, status.Mute{Muted: false})
}

func setMute(context echo.Context, address string, status bool, retryCount int) error {
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

	return fmt.Errorf("Attempted to set mute status %v times, could not", initCount+1)
}

func VolumeMute(context echo.Context) error {
	log.L.Debugf("Muting %s...", context.Param("address"))

	err := setMute(context, context.Param("address"), true, 4)
	if err != nil {
		log.L.Debugf("Error: %v", err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	log.L.Debugf("Done.")
	return context.JSON(http.StatusOK, status.Mute{Muted: true})
}

func BlankDisplay(context echo.Context) error {
	params := make(map[string]interface{})
	params["mode"] = "pictureOff"

	err := helpers.BuildAndSendPayload(context.Param("address"), "system", "setPowerSavingMode", params)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, status.Blanked{Blanked: true})

}

func UnblankDisplay(context echo.Context) error {
	params := make(map[string]interface{})
	params["mode"] = "off"

	err := helpers.BuildAndSendPayload(context.Param("address"), "system", "setPowerSavingMode", params)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, status.Blanked{Blanked: false})
}

func GetVolume(context echo.Context) error {
	response, err := helpers.GetVolume(context.Param("address"))
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

// GetInput gets the input that is currently being shown on the TV
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
}

func GetHardwareInfo(context echo.Context) error {
	response, err := helpers.GetHardwareInfo(context.Param("address"))
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

// GetActiveSignal determines if the current input on the TV is active or no
func GetActiveSignal(context echo.Context) error {
	response, err := helpers.GetActiveSignal(context.Param("address"), context.Param("port"))
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
