package helpers

import (
	"encoding/json"
	"log"

	"github.com/byuoitav/av-api/status"
)

func GetVolume(address string) (status.Volume, error) {
	log.Printf("Getting volume for %v", address)
	parentResponse, err := getAudioInformation(address)
	if err != nil {
		return status.Volume{}, err
	}
	log.Printf("%v", parentResponse)

	var output status.Volume
	for _, outerResult := range parentResponse.Result {

		for _, result := range outerResult {

			if result.Target == "speaker" {

				output.Volume = result.Volume
			}
		}
	}
	log.Printf("Done")

	return output, nil
}

func getAudioInformation(address string) (SonyAudioResponse, error) {
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
	return parentResponse, err

}

func GetMute(address string) (status.MuteStatus, error) {
	log.Printf("Getting mute status for %v", address)
	parentResponse, err := getAudioInformation(address)
	if err != nil {
		return status.MuteStatus{}, err
	}
	var output status.MuteStatus
	for _, outerResult := range parentResponse.Result {
		for _, result := range outerResult {
			if result.Target == "speaker" {
				output.Muted = result.Mute
			}
		}
	}

	log.Printf("Done")

	return status.MuteStatus{}, nil
}
