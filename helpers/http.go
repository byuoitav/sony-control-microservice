package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//SonyAudioResponse is the parent struct returned when we query audio state
type SonyAudioResponse struct {
	Result [][]SonyAudioSettings `json:"result"`
	ID     int                   `json:"id"`
}

//SonyAudioSettings is the child struct returned
type SonyAudioSettings struct {
	Target    string `json:"target"`
	Volume    int    `json:"volume"`
	Mute      bool   `json:"mute"`
	MaxVolume int    `json:"maxVolume"`
	MinVolume int    `json:"minVolume"`
}

//SonyTVRequest represents the struct we need to send.
type SonyTVRequest struct {
	Method  string                   `json:"method"`
	Version string                   `json:"version"`
	ID      int                      `json:"id"`
	Params  []map[string]interface{} `json:"params"`
}

//PostHTTP just sends a request
func PostHTTP(address string, payload SonyTVRequest, service string) ([]byte, error) {

	postBody, err := json.Marshal(payload)
	if err != nil {
		return []byte{}, err
	}

	addr := fmt.Sprintf("http://%s/sony/%s", address, service)

	request, err := http.NewRequest("POST", addr, bytes.NewBuffer(postBody))
	if err != nil {
		return []byte{}, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Auth-PSK", os.Getenv("SONY_TV_PSK"))

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return []byte{}, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	} else if body == nil {
		return []byte{}, errors.New("Response from device was blank")
	}

	return body, nil
}
