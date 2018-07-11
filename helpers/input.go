package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"

	"github.com/byuoitav/common/structs"
)

func GetInput(address string) (structs.InputStatus, error) {
	var output structs.InputStatus

	pwrState, err := GetPower(address)
	if err != nil {
		return output, err
	}
	if pwrState.Power != "on" {
		return output, nil
	}

	payload := SonyTVRequest{
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
	err = json.Unmarshal(response, &outputStruct)
	if err != nil || len(outputStruct.Result) < 1 {
		return output, err
	}
	//we need to parse the response for the value

	log.Printf("%+v", outputStruct)

	regexStr := `extInput:(.*?)\?port=(.*)`
	re := regexp.MustCompile(regexStr)

	matches := re.FindStringSubmatch(outputStruct.Result[0].URI)
	output.Input = fmt.Sprintf("%v!%v", matches[1], matches[2])

	return output, nil
}
