package ircc

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func GetCapability(address string) (Capability, error) {
	var postBody = []byte(`{
  "method": "getRemoteControllerInfo",
  "params": [],
  "id": 10,
  "version": "1.0"
}`)

	req, err := http.NewRequest("POST", "http://"+address+"/sony/system", bytes.NewBuffer(postBody))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return Capability{}, err
	}

	capabilities := Capability{}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Capability{}, err
	} else if body != nil {
		err = json.Unmarshal(body, &capabilities)
		if err != nil {
			return Capability{}, err
		}
	} else { // If the body doesn't have anything
		return Capability{}, errors.New("Response from device was blank")
	}

	return capabilities, nil
}
