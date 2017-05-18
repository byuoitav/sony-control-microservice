package helpers

import "github.com/byuoitav/av-api/status"

func GetBlankedStatus(address string) (status.BlankedStatus, error) {

	var blanked status.BlankedStatus

	power, err := GetPower(address)
	if err != nil {
		return blanked, err
	}

	switch power.Power {
	case "on":
		blanked.Blanked = false
	case "standby":
		blanked.Blanked = true
	}
	return blanked, nil
}
