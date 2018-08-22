package helpers

import (
	"errors"
	"strings"
	"time"

	"github.com/byuoitav/common/nerr"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/status"
)

// SetPower will set the projector status to on or standby.
func SetPower(address string, status bool) error {
	params := make(map[string]interface{})
	params["status"] = status

	//check to see if it's off, if it is we need to wait if we're turning the thing on after we return
	preStatus, err := GetPower(address)
	if err != nil {
		return err
	}

	err = BuildAndSendPayload(address, "system", "setPowerStatus", params)
	if err != nil {
		return err
	}

	postStatus, err := GetPower(address)
	if err != nil {
		return err
	}

	log.L.Infof("%v", postStatus)

	if status && postStatus.Power != "on" {
		// do we want to retry the command
		return errors.New("Power wasn't set successfully")
	} else if !status && postStatus.Power != "standby" {
		return errors.New("Power wasn't set successfully")
	}

	//we need to wait for a little bit to let the tv finish so it doesn't override

	if preStatus.Power == "standby" && status {
		log.L.Infof("Waiting....")
		time.Sleep(1750 * time.Millisecond)
	}

	return nil
}

// GetPower gets the power status and returns that, or an error.
func GetPower(address string) (status.Power, error) {

	var output status.Power

	payload := SonyTVRequest{
		Params: []map[string]interface{}{},
		Method: "getPowerStatus", Version: "1.0",
		ID: 1,
	}

	response, err := PostHTTP(address, payload, "system")
	if err != nil {
		return status.Power{}, err
	}

	powerStatus := string(response)
	log.L.Infof("Device returned: %s", powerStatus)
	if strings.Contains(powerStatus, "active") {
		output.Power = "on"
	} else if strings.Contains(powerStatus, "standby") {
		output.Power = "standby"
	} else {
		return status.Power{}, errors.New("Error getting power status")
	}

	return output, nil
}

// WEB API CODE IS IN AUSTRALIA (down undah...)

// this will be the command struct that will help set up the body that will be sent
type sonyCommand struct {
	Method  string                   `json:"method"`
	Version string                   `json:"version"`
	ID      int                      `json:"id"`
	Params  []map[string]interface{} `json:"params"`
}

// SetPowerStatus sets the power status through the testing Sony API
func SetPowerStatus(address string, status bool) error {

	// Make params which is what will contain the power on/off command
	params := make(map[string]interface{})
	params["status"] = status

	// currentStatus will get the current power state of the projector
	currentStatus, err := GetPowerStatus(address)
	if err != nil {
		return err
	}

	err = BuildAndSendPayload(address, "system", "setPowerStatus", params)
	if err != nil {
		return err
	}

	postStatus, err := GetPower(address)
	if err != nil {
		return err
	}

	log.L.Infof("%v", postStatus)

	if status && postStatus.Power != "on" {
		// do we want to retry the command
		return errors.New("Power wasn't set successfully")
	} else if !status && postStatus.Power != "standby" {
		return errors.New("Power wasn't set successfully")
	}

	//we need to wait for a little bit to let the tv finish so it doesn't override

	if currentStatus.Power == "standby" && status {
		log.L.Infof("Waiting....")
		time.Sleep(1750 * time.Millisecond)
	}

	return nil
}

// GetPowerStatus retrieves the power status through the testing Sony API
func GetPowerStatus(address string) (status.Power, error) {

	// powerOutput is the status.Power JSON thingy...
	var powerOutput status.Power

	request := sonyCommand{
		Method:  "getPowerStatus",
		Version: "1.0",
		ID:      1,
		Params:  []map[string]interface{}{},
	}

	response, err := PostHTTP(address, request, "system")
	if err != nil {
		return status.Power{}, err
	}

	powerStatusAPI := string(response)
	log.L.Infof("Device returned: %s", powerStatusAPI)
	if strings.Contains(powerStatusAPI, "active") {
		powerOutput.Power = "on"
	} else if strings.Contains(powerStatusAPI, "standby") {
		powerOutput.Power = "standby"
	} else {
		return status.Power{}, nerr.Translate(err).Addf("There was an error getting power status")
	}

	return powerOutput, nil
}
