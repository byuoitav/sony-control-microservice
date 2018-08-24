package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/nerr"
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
	URI    string `json:"uri"`
	Source string `json:"source"`
	Title  string `json:"title"`
}

type SonyAVContentResponse struct {
	Result []SonyAVContentSettings `json:"result"`
	ID     int                     `json:"id"`
}

// SonyRequest represents the struct we need to send.
type SonyRequest struct {
	Method  string                   `json:"method"`
	Version string                   `json:"version"`
	ID      int                      `json:"id"`
	Params  []map[string]interface{} `json:"params"`
}

// PostHTTP just sends a request
func PostHTTP(address string, payload interface{}, service string) ([]byte, *nerr.E) {

	postBody, err := json.Marshal(payload)
	if err != nil {
		return []byte{}, nerr.Translate(err).Addf("Failed ")
	}

	// This is the body that gets sent in the HTTP Post
	log.L.Debugf("%s", postBody)

	addr := fmt.Sprintf("http://%s/sony/%s", address, service)

	request, err := http.NewRequest("POST", addr, bytes.NewBuffer(postBody))
	if err != nil {
		return []byte{}, nerr.Translate(err)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Auth-PSK", os.Getenv("SONY_TV_PSK"))

	client := &http.Client{Timeout: 3 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		return []byte{}, nerr.Translate(err)
	}

	body, err := ioutil.ReadAll(response.Body)

	// This is the response back from the HTTP Post
	log.L.Debugf("Body: %s", body)

	if err != nil {
		return []byte{}, nerr.Translate(err)
	} else if response.StatusCode != http.StatusOK {
		return []byte{}, nerr.Create(string(body), "There was an error")
	} else if body == nil {
		return []byte{}, nerr.Create(string(body), "Response from device was blank")
	}

	defer response.Body.Close()
	return body, nil
}

// BuildAndSendPayload makes the payload which is the SonyRequest
func BuildAndSendPayload(address string, service string, method string, params map[string]interface{}) *nerr.E {
	payload := SonyRequest{
		Params:  []map[string]interface{}{params},
		Method:  method,
		Version: "1.0",
		ID:      1,
	}

	_, err := PostHTTP(address, payload, service)
	if err != nil {
		return nerr.Translate(err)
	}

	return nil

}
