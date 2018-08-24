package helpers

import (
	"encoding/json"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/nerr"
	"github.com/byuoitav/common/status"
)

func GetVolume(address string) (status.Volume, *nerr.E) {
	log.L.Infof("Getting volume for %v", address)
	parentResponse, err := getAudioInformation(address)
	if err != nil {
		return status.Volume{}, err
	}
	log.L.Infof("%v", parentResponse)

	var output status.Volume
	for _, outerResult := range parentResponse.Result {

		for _, result := range outerResult {

			if result.Target == "speaker" {

				output.Volume = result.Volume
			}
		}
	}
	log.L.Infof("Done")

	return output, nil
}

func getAudioInformation(address string) (SonyAudioResponse, *nerr.E) {
	payload := SonyRequest{
		Params:  []map[string]interface{}{},
		Method:  "getVolumeInformation",
		Version: "1.0",
		ID:      1,
	}

	log.L.Infof("%+v", payload)

	resp, err := PostHTTP(address, payload, "audio")
	if err != nil {
		return SonyAudioResponse{}, err.Addf("Failed to post")
	}

	parentResponse := SonyAudioResponse{}

	log.L.Infof("%s", resp)

	er := json.Unmarshal(resp, &parentResponse)
	if er != nil {
		return SonyAudioResponse{}, nerr.Translate(er).Addf("Failed to get audio information")
	}

	return parentResponse, nil
}

func GetMute(address string) (status.Mute, *nerr.E) {
	log.L.Infof("Getting mute status for %v", address)
	parentResponse, err := getAudioInformation(address)
	if err != nil {
		return status.Mute{}, err
	}
	var output status.Mute
	for _, outerResult := range parentResponse.Result {
		for _, result := range outerResult {
			if result.Target == "speaker" {
				log.L.Infof("local mute: %v", result.Mute)
				output.Muted = result.Mute
			}
		}
	}

	log.L.Infof("Done")

	return output, nil
}
