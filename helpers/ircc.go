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

	response, err := PostHTTP("http://"+address+"/sony/system", postBody)
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
	var postBody = []byte(`<?xml version="1.0"?><s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/" s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"><s:Body><u:X_SendIRCC xmlns:u="urn:schemas-sony-com:service:IRCC:1"><IRCCCode>` + command + `</IRCCCode></u:X_SendIRCC></s:Body></s:Envelope>`)

	response, err := PostSOAP("http://"+address+"/sony/IRCC", postBody)
	if err != nil {
		return "", err
	}

	return string(response), nil
}
