package helpers

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

func PostHTTP(address string, postBody []byte) ([]byte, error) {
	request, err := http.NewRequest("POST", address, bytes.NewBuffer(postBody))
	request.Header.Set("Content-Type", "application/json")

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
