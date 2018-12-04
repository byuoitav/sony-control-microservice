package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/byuoitav/common/log"
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

type SonyAVContentSettings struct {
	URI        string `json:"uri"`
	Source     string `json:"source"`
	Title      string `json:"title"`
	Status     string `json:"status"`
	Connection bool   `json:"connection"`
}

type SonyAVContentResponse struct {
	Result []SonyAVContentSettings `json:"result"`
	ID     int                     `json:"id"`
}

type SonyMultiAVContentResponse struct {
	Result [][]SonyAVContentSettings `json:"result"`
	ID     int                       `json:"id"`
}

//SonyTVRequest represents the struct we need to send.
type SonyTVRequest struct {
	Method  string                   `json:"method"`
	Version string                   `json:"version"`
	ID      int                      `json:"id"`
	Params  []map[string]interface{} `json:"params"`
}

type SonyTVSystemResponse struct {
	ID     int `json:"id"`
	Result []SonySystemInformation
}

type SonySystemInformation struct {
	Product    string `json:"product"`
	Region     string `json:"region,omitempty"`
	Language   string `json:"language,omitempty"`
	Model      string `json:"model"`
	Serial     string `json:"serial,omitempty"`
	MAC        string `json:"macAddr,omitempty"`
	Name       string `json:"name"`
	Generation string `json:"generation,omitempty"`
	Area       string `json:"area,omitempty"`
	CID        string `json:"cid,omitempty"`
}

type SonyNetworkResponse struct {
	ID     int `json:"id"`
	Result [][]SonyTVNetworkInformation
}

type SonyTVNetworkInformation struct {
	NetworkInterface string   `json:"netif"`
	HardwareAddress  string   `json:"hwAddr"`
	IPv4             string   `json:"ipAddrV4"`
	IPv6             string   `json:"ipAddrV6"`
	Netmask          string   `json:"netmask"`
	Gateway          string   `json:"gateway"`
	DNS              []string `json:"dns"`
}

//PostHTTP just sends a request
func PostHTTP(address string, payload SonyTVRequest, service string) ([]byte, error) {

	postBody, err := json.Marshal(payload)
	if err != nil {
		return []byte{}, err
	}

	log.L.Debugf("%s", postBody)

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

	log.L.Debugf("Body: %s", body)

	if err != nil {
		return []byte{}, err
	} else if response.StatusCode != http.StatusOK {
		return []byte{}, errors.New(string(body))
	} else if body == nil {
		return []byte{}, errors.New("Response from device was blank")
	}

	defer response.Body.Close()
	return body, nil
}

func BuildAndSendPayload(address string, service string, method string, params map[string]interface{}) error {
	payload := SonyTVRequest{
		Params:  []map[string]interface{}{params},
		Method:  method,
		Version: "1.0",
		ID:      1,
	}

	_, err := PostHTTP(address, payload, service)

	return err

}
