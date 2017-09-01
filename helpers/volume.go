package helpers

import (
	"encoding/json"
	"log"

	se "github.com/byuoitav/av-api/statusevaluators"
)

func GetVolume(address string) (se.Volume, error) {
	log.Printf("Getting volume for %v", address)
	parentResponse, err := getAudioInformation(address)
	if err != nil {
		return se.Volume{}, err
	}
	log.Printf("%v", parentResponse)

	var output se.Volume
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

func GetMute(address string) (se.MuteStatus, error) {
	log.Printf("Getting mute status for %v", address)
	parentResponse, err := getAudioInformation(address)
	if err != nil {
		return se.MuteStatus{}, err
	}
	var output se.MuteStatus
	for _, outerResult := range parentResponse.Result {
		for _, result := range outerResult {
			if result.Target == "speaker" {
				log.Printf("local mute: %v", result.Mute)
				output.Muted = result.Mute
			}
		}
	}

	log.Printf("Done")

	return output, nil
}
