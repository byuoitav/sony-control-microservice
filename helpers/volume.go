package helpers

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/byuoitav/av-api/status"
)

func GetVolume(address string) (status.Volume, error) {

	payload := SonyTVRequest{
		Params:  []map[string]interface{}{},
		Method:  "getVolumeInformation",
		Version: "1.0",
		ID:      1,
	}

	log.Printf("%+v", payload)

	resp, err := PostHTTP(address, payload, "audio")

	parentResponse := SonyAudioResponse{}

	log.Printf("%s", resp)

	err = json.Unmarshal(resp, &parentResponse)
	if err != nil {
		return status.Volume{}, err
	}

	log.Printf("%+v", parentResponse)

	var output status.Volume
	for _, outerResult := range parentResponse.Result {

		for _, result := range outerResult {

			if result.Target == "speaker" {

				output.Volume = result.Volume

			}

		}

	}

	if output.Volume == 0 {

		return status.Volume{}, errors.New("Could not find volume")

	} else {

		log.Printf("Done")

	}

	return output, nil
}
