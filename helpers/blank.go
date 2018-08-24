package helpers

import (
	"encoding/json"
	"fmt"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/nerr"
	"github.com/byuoitav/common/status"
)

// SonyBaseResult struct is the basic SonyRequest kinda
type SonyBaseResult struct {
	ID     int                 `json:"id"`
	Result []map[string]string `json:"result"`
	Error  []interface{}       `json:"error"`
}

// GetBlanked returns the video mute status of the Sony Device
func GetBlanked(address string) (status.Blanked, *nerr.E) {

	var blanked status.Blanked

	payload := SonyRequest{
		Params:  []map[string]interface{}{},
		Method:  "getPowerSavingMode",
		Version: "1.0",
		ID:      1,
	}

	log.L.Debugf("%+v", payload)

	resp, err := PostHTTP(address, payload, "system")
	if err != nil {
		log.L.Errorf("ERROR: %v", err.Error())
		return blanked, err
	}

	re := SonyBaseResult{}
	er := json.Unmarshal(resp, &re)
	if err != nil {
		return blanked, nerr.Create(fmt.Sprintf("failed to unmarshal response from tv: %s", er), "Failed")
	}

	// make sure there is a result
	if len(re.Result) == 0 {
		return blanked, nerr.Create(fmt.Sprintf("error response from tv: %s", re.Error), "Failed")
	}

	if val, ok := re.Result[0]["mode"]; ok {
		if val == "pictureOff" {
			blanked.Blanked = true
		} else {
			blanked.Blanked = false
		}
	}

	return blanked, nil
}
