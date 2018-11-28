package helpers

import (
	"errors"
	"strings"
	"time"

	"github.com/byuoitav/common/log"

	"github.com/byuoitav/common/status"
)

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
