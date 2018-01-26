package helpers

import (
	"encoding/json"
	"log"

	se "github.com/byuoitav/av-api/statusevaluators"
)

type SonyBaseResult struct {
	ID     int                 `json:"id"`
	Result []map[string]string `json:"result"`
}

func GetBlankedStatus(address string) (se.BlankedStatus, error) {

	var blanked se.BlankedStatus

	payload := SonyTVRequest{
		Params:  []map[string]interface{}{},
		Method:  "getPowerSavingMode",
		Version: "1.0",
		ID:      1,
	}

	log.Printf("%+v", payload)

	resp, err := PostHTTP(address, payload, "system")
	if err != nil {
		log.Printf("ERROR: %v", err.Error())
		return blanked, err
	}

	re := SonyBaseResult{}
	err = json.Unmarshal(resp, &re)

	if val, ok := re.Result[0]["mode"]; ok {
		if val == "pictureOff" {
			blanked.Blanked = true
		} else {
			blanked.Blanked = false
		}
	}

	return blanked, nil
}
