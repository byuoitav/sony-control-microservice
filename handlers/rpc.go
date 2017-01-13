package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/byuoitav/sony-control-microservice/helpers"
	"github.com/labstack/echo"
)

func PowerOn(context echo.Context) error {
	log.Printf("Powering on %s...", context.Param("address"))

	err := setPower(context.Param("address"), true)
	if err != nil {
		log.Printf("Error: %v", err.Error())
		return err
	}

	log.Printf("Done.")
	return nil
}

func setPower(address string, status bool) error {
	params := make(map[string]interface{})
	params["status"] = status

	return BuildAndSendPayload(address, "system", "setPowerStatus", params)
}

func Standby(context echo.Context) error {
	log.Printf("Powering off %s...", context.Param("address"))

	err := setPower(context.Param("address"), false)
	if err != nil {
		log.Printf("Error: %v", err.Error())
		return err
	}

	log.Printf("Done.")
	return nil
}

func SwitchInput(context echo.Context) error {
	log.Printf("Switching input for %s to %s ...", context.Param("address"), context.Param("port"))
	address := context.Param("address")
	port := context.Param("port")

	splitPort := strings.Split(port, "!")

	params := make(map[string]interface{})
	params["uri"] = fmt.Sprintf("extInput:%s?port=%s", splitPort[0], splitPort[1])

	err := BuildAndSendPayload(address, "avContent", "setPlayContent", params)
	if err != nil {
		return err
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

	err := BuildAndSendPayload(address, "audio", "setAudioVolume", params)
	if err != nil {
		return err
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
		return err
	}

	log.Printf("Done.")
	return nil
}

func setMute(address string, status bool) error {
	params := make(map[string]interface{})
	params["status"] = status

	return BuildAndSendPayload(address, "audio", "setAudioMute", params)
}

func VolumeMute(context echo.Context) error {
	log.Printf("Muting %s...", context.Param("address"))

	err := setMute(context.Param("address"), true)
	if err != nil {
		log.Printf("Error: %v", err.Error())
		return err
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
	log.Printf("Getting volume for %s...", context.Param("address"))

	payload := helpers.SonyTVRequest{
		Params:  []map[string]interface{}{},
		Method:  "getVolumeInformation",
		Version: "1.0",
		ID:      1,
	}

	log.Printf("%+v", payload)

	resp, err := helpers.PostHTTP(context.Param("address"), payload, "audio")

	parentResponse := helpers.SonyAudioResponse{}

	log.Printf("%s", resp)

	err = json.Unmarshal(resp, &parentResponse)
	if err != nil {
		return err
	}

	log.Printf("%+v", parentResponse)

	b, err := json.Marshal(parentResponse.Result[0])
	if err != nil {
		return err
	}

	context.Response().Write(b)

	log.Printf("Done")
	return nil
}

func BuildAndSendPayload(address string, service string, method string, params map[string]interface{}) error {
	payload := helpers.SonyTVRequest{
		Params:  []map[string]interface{}{params},
		Method:  method,
		Version: "1.0",
		ID:      1,
	}

	_, err := helpers.PostHTTP(address, payload, service)

	return err

}
