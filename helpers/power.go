package helpers

import (
	"errors"
	"log"
	"strings"

	"github.com/byuoitav/av-api/status"
)

func SetPower(address string, status bool) error {
	params := make(map[string]interface{})
	params["status"] = status

	return BuildAndSendPayload(address, "system", "setPowerStatus", params)
}

func GetPower(address string) (status.PowerStatus, error) {

	var output status.PowerStatus

	payload := SonyTVRequest{
		Params:  []map[string]interface{}{},
		Method:  "getPowerStatus",
		Version: "1.0",
		ID:      1,
	}

	response, err := PostHTTP(address, payload, "system")
	if err != nil {
		return status.PowerStatus{}, err
	}

	powerStatus := string(response)
	log.Printf("Device returned: %s", powerStatus)
	if strings.Contains(powerStatus, "active") {
		output.Power = "on"
	} else if strings.Contains(powerStatus, "standby") {
		output.Power = "standby"
	} else {
		return status.PowerStatus{}, errors.New("Error getting power status")
	}

	return output, nil
}
