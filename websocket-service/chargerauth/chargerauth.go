package chargerauth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	stations "github.com/gregszalay/ocpp-csms-common-types/devices"
)

func GetCharger(chargerId string) (stations.Charger, error) {
	//resp, err := http.Get("http://localhost:5000/charger/" + id)
	deviceServiceHost := "host.docker.internal"
	if d := os.Getenv("DEVICE_SERVICE_HOST"); d != "" {
		deviceServiceHost = d
	}
	resp, err := http.Get(fmt.Sprintf("http://%s:5000/charger/%s", deviceServiceHost, chargerId))
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
