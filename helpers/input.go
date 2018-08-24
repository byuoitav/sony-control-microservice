package helpers

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/status"
)

func GetInput(address string) (status.Input, error) {
	var output status.Input

	pwrState, err := GetPower(address)
	if err != nil {
		return output, err
	}
	if pwrState.Power != "on" {
		return output, nil
	}

	payload := SonyRequest{
		Params:  []map[string]interface{}{},
		Method:  "getPlayingContentInfo",
		ID:      1,
		Version: "1.0",
	}

	response, err := PostHTTP(address, payload, "avContent")
	if err != nil {
		return output, err
	}

	var outputStruct SonyAVContentResponse
	er := json.Unmarshal(response, &outputStruct)
	if err != nil || len(outputStruct.Result) < 1 {
		return output, er
	}
	//we need to parse the response for the value

	log.L.Infof("%+v", outputStruct)

	regexStr := `extInput:(.*?)\?port=(.*)`
	re := regexp.MustCompile(regexStr)

	matches := re.FindStringSubmatch(outputStruct.Result[0].URI)
	output.Input = fmt.Sprintf("%v!%v", matches[1], matches[2])

	return output, nil
}
