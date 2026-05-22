package helpers

import (
	"context"
	"encoding/json"

	"github.com/byuoitav/common/log"

	"github.com/byuoitav/common/status"
)

func GetVolume(address string) (status.Volume, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultSonyTimeout)
	defer cancel()

	return GetVolumeWithContext(ctx, address)
}

func GetVolumeWithContext(ctx context.Context, address string) (status.Volume, error) {
	log.L.Infof("Getting volume for %v", address)
	parentResponse, err := getAudioInformation(ctx, address)
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

func getAudioInformation(ctx context.Context, address string) (SonyAudioResponse, error) {
	payload := SonyTVRequest{
		Params:  []map[string]interface{}{},
		Method:  "getVolumeInformation",
		Version: "1.0",
		ID:      1,
	}

	log.L.Infof("%+v", payload)

	resp, err := PostHTTPWithContext(ctx, address, "audio", payload)

	parentResponse := SonyAudioResponse{}

	log.L.Infof("%s", resp)

	err = json.Unmarshal(resp, &parentResponse)
	return parentResponse, err

}

func GetMute(address string) (status.Mute, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultSonyTimeout)
	defer cancel()

	return GetMuteWithContext(ctx, address)
}

func GetMuteWithContext(ctx context.Context, address string) (status.Mute, error) {
	log.L.Infof("Getting mute status for %v", address)
	parentResponse, err := getAudioInformation(ctx, address)
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
