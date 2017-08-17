package helpers

import (
	"errors"
	"log"
	"strings"

	se "github.com/byuoitav/av-api/statusevaluators"
)

func SetPower(address string, status bool) (se.PowerStatus, error) {
	params := make(map[string]interface{})
	params["status"] = status

	response, err := BuildAndSendPayload(address, "system", "setPowerStatus", params)
	if err != nil {
		return se.PowerStatus{}, err
	}

}

func GetPower(address string) (se.PowerStatus, error) {

	var output se.PowerStatus

	payload := SonyTVRequest{
		Params:  []map[string]interface{}{},
		Method:  "getPowerStatus",
		Version: "1.0",
		ID:      1,
	}

	response, err := PostHTTP(address, payload, "system")
	if err != nil {
		return se.PowerStatus{}, err
	}

	powerStatus := string(response)
	log.Printf("Device returned: %s", powerStatus)
	if strings.Contains(powerStatus, "active") {
		output.Power = "on"
	} else if strings.Contains(powerStatus, "standby") {
		output.Power = "standby"
	} else {
		return se.PowerStatus{}, errors.New("Error getting power status")
	}

	return output, nil
}
