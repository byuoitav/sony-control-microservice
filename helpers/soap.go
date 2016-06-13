package helpers

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
)

func PostSOAP(address string, postBody []byte) ([]byte, error) {
	request, err := http.NewRequest("POST", address, bytes.NewBuffer(postBody))
	request.Header.Set("Content-Type", "text/xml; charset=UTF-8")
	request.Header.Set("Soapaction", `"urn:schemas-sony-com:service:IRCC:1#X_SendIRCC"`)
	request.Header.Set("X-Auth-Psk", os.Getenv("SONY_TV_PSK"))

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
