package helpers

import (
	"encoding/json"
	"strings"
)

type Commands struct {
	ID     int           `json:"id"`
	Result []interface{} `json:"result"`
}

func GetCommands(address string) (Commands, error) {
	var postBody = []byte(`{
  "method": "getRemoteControllerInfo",
  "params": [],
  "id": 10,
  "version": "1.0"
}`)

	response, err := PostHTTP("http://"+address+"/sony/system", postBody)
	if err != nil {
		return Commands{}, err
	}

	commands := Commands{}
	err = json.Unmarshal(response, &commands)
	if err != nil {
		return Commands{}, err
	}

	return commands, nil
}

func SendCommand(address string, command string) (string, error) {
	allCommands := map[string]string{
		"*AD":              "AAAAAgAAABoAAAA7Aw==",
		"ActionMenu":       "AAAAAgAAAMQAAABLAw==",
		"Analog":           "AAAAAgAAAHcAAAANAw==",
		"Analog2":          "AAAAAQAAAAEAAAA4Aw==",
		"AnalogRgb1":       "AAAAAQAAAAEAAABDAw==",
		"Assists":          "AAAAAgAAAMQAAAA7Aw==",
		"Audio":            "AAAAAQAAAAEAAAAXAw==",
		"AudioMixDown":     "AAAAAgAAABoAAAA9Aw==",
		"AudioMixUp":       "AAAAAgAAABoAAAA8Aw==",
		"AudioQualityMode": "AAAAAgAAAJcAAAB7Aw==",
		"BS":               "AAAAAgAAAJcAAAAsAw==",
		"BSCS":             "AAAAAgAAAJcAAAAQAw==",
		"Blue":             "AAAAAgAAAJcAAAAkAw==",
		"CS":               "AAAAAgAAAJcAAAArAw==",
		"ChannelDown":      "AAAAAQAAAAEAAAARAw==",
		"ChannelUp":        "AAAAAQAAAAEAAAAQAw==",
		"ClosedCaption":    "AAAAAgAAAKQAAAAQAw==",
		"Component1":       "AAAAAgAAAKQAAAA2Aw==",
		"Component2":       "AAAAAgAAAKQAAAA3Aw==",
		"Confirm":          "AAAAAQAAAAEAAABlAw==",
		"CursorDown":       "AAAAAgAAAJcAAABQAw==",
		"CursorLeft":       "AAAAAgAAAJcAAABNAw==",
		"CursorRight":      "AAAAAgAAAJcAAABOAw==",
		"CursorUp":         "AAAAAgAAAJcAAABPAw==",
		"DOT":              "AAAAAgAAAJcAAAAdAw==",
		"DUX":              "AAAAAgAAABoAAABzAw==",
		"Ddata":            "AAAAAgAAAJcAAAAVAw==",
		"DemoMode":         "AAAAAgAAAJcAAAB8Aw==",
		"DemoSurround":     "AAAAAgAAAHcAAAB7Aw==",
		"Digital":          "AAAAAgAAAJcAAAAyAw==",
		"DigitalToggle":    "AAAAAgAAAHcAAABSAw==",
		"Display":          "AAAAAQAAAAEAAAA6Aw==",
		"Down":             "AAAAAQAAAAEAAAB1Aw==",
		"DpadCenter":       "AAAAAgAAAJcAAABKAw==",
		"EPG":              "AAAAAgAAAKQAAABbAw==",
		"Enter":            "AAAAAQAAAAEAAAALAw==",
		"Exit":             "AAAAAQAAAAEAAABjAw==",
		"FlashMinus":       "AAAAAgAAAJcAAAB5Aw==",
		"FlashPlus":        "AAAAAgAAAJcAAAB4Aw==",
		"FootballMode":     "AAAAAgAAABoAAAB2Aw==",
		"Forward":          "AAAAAgAAAJcAAAAcAw==",
		"GGuide":           "AAAAAQAAAAEAAAAOAw==",
		"Green":            "AAAAAgAAAJcAAAAmAw==",
		"Hdmi1":            "AAAAAgAAABoAAABaAw==",
		"Hdmi2":            "AAAAAgAAABoAAABbAw==",
		"Hdmi3":            "AAAAAgAAABoAAABcAw==",
		"Hdmi4":            "AAAAAgAAABoAAABdAw==",
		"Help":             "AAAAAgAAAMQAAABNAw==",
		"Home":             "AAAAAQAAAAEAAABgAw==",
		"iManual":          "AAAAAgAAABoAAAB7Aw==",
		"Input":            "AAAAAQAAAAEAAAAlAw==",
		"Jump":             "AAAAAQAAAAEAAAA7Aw==",
		"Left":             "AAAAAQAAAAEAAAA0Aw==",
		"Media":            "AAAAAgAAAJcAAAA4Aw==",
		"MediaAudioTrack":  "AAAAAQAAAAEAAAAXAw==",
		"Mode3D":           "AAAAAgAAAHcAAABNAw==",
		"Mute":             "AAAAAQAAAAEAAAAUAw==",
		"Netflix":          "AAAAAgAAABoAAAB8Aw==",
		"Next":             "AAAAAgAAAJcAAAA9Aw==",
		"Num0":             "AAAAAQAAAAEAAAAJAw==",
		"Num1":             "AAAAAQAAAAEAAAAAAw==",
		"Num11":            "AAAAAQAAAAEAAAAKAw==",
		"Num12":            "AAAAAQAAAAEAAAALAw==",
		"Num2":             "AAAAAQAAAAEAAAABAw==",
		"Num3":             "AAAAAQAAAAEAAAACAw==",
		"Num4":             "AAAAAQAAAAEAAAADAw==",
		"Num5":             "AAAAAQAAAAEAAAAEAw==",
		"Num6":             "AAAAAQAAAAEAAAAFAw==",
		"Num7":             "AAAAAQAAAAEAAAAGAw==",
		"Num8":             "AAAAAQAAAAEAAAAHAw==",
		"Num9":             "AAAAAQAAAAEAAAAIAw==",
		"OneTouchTimeRec":  "AAAAAgAAABoAAABkAw==",
		"OneTouchView":     "AAAAAgAAABoAAABlAw==",
		"Options":          "AAAAAgAAAJcAAAA2Aw==",
		"PAP":              "AAAAAgAAAKQAAAB3Aw==",
		"Pause":            "AAAAAgAAAJcAAAAZAw==",
		"PhotoFrame":       "AAAAAgAAABoAAABVAw==",
		"PicOff":           "AAAAAQAAAAEAAAA+Aw==",
		"PictureMode":      "AAAAAQAAAAEAAABkAw==",
		"PictureOff":       "AAAAAQAAAAEAAAA+Aw==",
		"Play":             "AAAAAgAAAJcAAAAaAw==",
		"PopUpMenu":        "AAAAAgAAABoAAABhAw==",
		"PowerOff":         "AAAAAQAAAAEAAAAvAw==",
		"Prev":             "AAAAAgAAAJcAAAA8Aw==",
		"Rec":              "AAAAAgAAAJcAAAAgAw==",
		"Red":              "AAAAAgAAAJcAAAAlAw==",
		"Return":           "AAAAAgAAAJcAAAAjAw==",
		"Rewind":           "AAAAAgAAAJcAAAAbAw==",
		"Right":            "AAAAAQAAAAEAAAAzAw==",
		"ShopRemoteControlForcedDynamic": "AAAAAgAAAJcAAABqAw==",
		"Sleep":             "AAAAAQAAAAEAAAAvAw==",
		"SleepTimer":        "AAAAAQAAAAEAAAA2Aw==",
		"Stop":              "AAAAAgAAAJcAAAAYAw==",
		"SubTitle":          "AAAAAgAAAJcAAAAoAw==",
		"SyncMenu":          "AAAAAgAAABoAAABYAw==",
		"Teletext":          "AAAAAQAAAAEAAAA\\/Aw==",
		"TenKey":            "AAAAAgAAAJcAAAAMAw==",
		"TopMenu":           "AAAAAgAAABoAAABgAw==",
		"Tv":                "AAAAAQAAAAEAAAAkAw==",
		"TvAnalog":          "AAAAAQAAAAEAAAA4Aw==",
		"TvAntennaCable":    "AAAAAQAAAAEAAAAqAw==",
		"TvInput":           "AAAAAQAAAAEAAAAlAw==",
		"TvPower":           "AAAAAQAAAAEAAAAVAw==",
		"TvSatellite":       "AAAAAgAAAMQAAABOAw==",
		"Tv_Radio":          "AAAAAgAAABoAAABXAw==",
		"Up":                "AAAAAQAAAAEAAAB0Aw==",
		"Video1":            "AAAAAQAAAAEAAABAAw==",
		"Video2":            "AAAAAQAAAAEAAABBAw==",
		"VolumeDown":        "AAAAAQAAAAEAAAATAw==",
		"VolumeUp":          "AAAAAQAAAAEAAAASAw==",
		"WakeUp":            "AAAAAQAAAAEAAAAuAw==",
		"Wide":              "AAAAAgAAAKQAAAA9Aw==",
		"WirelessSubwoofer": "AAAAAgAAAMQAAAB+Aw==",
		"Yellow":            "AAAAAgAAAJcAAAAnAw==",
	}

	for name, key := range allCommands {
		if strings.ToLower(name) == strings.ToLower(command) {
			command = key
		}
	}

	var postBody = []byte(`<?xml version="1.0"?><s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/" s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"><s:Body><u:X_SendIRCC xmlns:u="urn:schemas-sony-com:service:IRCC:1"><IRCCCode>` + command + `</IRCCCode></u:X_SendIRCC></s:Body></s:Envelope>`)

	response, err := PostSOAP("http://"+address+"/sony/IRCC", postBody)
	if err != nil {
		return "", err
	}

	return string(response), nil
}

func FloodCommand(address string, command string, count int) {
	for i := 1; i <= count; i++ {
		go SendCommand(address, command)
	}
}
