package helpers

import "encoding/json"

type Capability struct {
	ID     int           `json:"id"`
	Result []interface{} `json:"result"`
}

func GetCapability(address string) (Capability, error) {
	var postBody = []byte(`{
  "method": "getRemoteControllerInfo",
  "params": [],
  "id": 10,
  "version": "1.0"
}`)

	response, err := Post("http://"+address+"/sony/system", postBody)
	if err != nil {
		return Capability{}, err
	}

	capabilities := Capability{}
	err = json.Unmarshal(response, &capabilities)
	if err != nil {
		return Capability{}, err
	}

	return capabilities, nil
}

func SendCommand(address string, command string) (string, error) {
	var postBody = []byte(`{
  "method": "getRemoteControllerInfo",
  "params": [],
  "id": 10,
  "version": "1.0"
}`)

	response, err := Post("http://"+address+"/sony/IRCC", postBody)
	if err != nil {
		return "", err
	}

	return string(response), nil
}
