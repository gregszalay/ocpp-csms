package chargerauth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	stations "github.com/gregszalay/ocpp-csms-common-types/types"
)

func GetCharger(chargerId string) (stations.Charger, error) {
	//resp, err := http.Get("http://localhost:5000/charger/" + id)
	resp, err := http.Get("http://host.docker.internal:5000/charger/" + chargerId)
	if err != nil {
		return stations.Charger{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return stations.Charger{}, err
	}
	var newCharger stations.Charger
	error := json.Unmarshal(body, &newCharger)
	if error != nil {
		return stations.Charger{}, err
	}
	return newCharger, nil
}
