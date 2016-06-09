package system

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// GetAvailableCommands gets all commands supported by a specified host
func GetAvailableCommands(host string) (Commands, error) {
	var postBody = []byte(`{"method":"getRemoteControllerInfo","params":[],"id":10,"version":"1.0"}`)

	req, err := http.NewRequest("POST", "http://"+host+"/sony/system", bytes.NewBuffer(postBody))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return Commands{}, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Commands{}, err
	}

	json.Unmarshal(body, &Commands)

	allCommands := Commands{}

	// for i := range elasticAllBuildings.Aggregations.FullName.Buckets {
	// 	buildingName := strings.ToUpper(elasticAllBuildings.Aggregations.FullName.Buckets[i].Key)
	// 	building := Building{Building: buildingName}
	//
	// 	allBuildings.Buildings = append(allBuildings.Buildings, building)
	// }

	return allCommands, nil
}
