package authentication

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gregszalay/ocpp-csms-common-types/devices"
	log "github.com/sirupsen/logrus"
)

func AuthenticateChargingStation(chargerId string, r *http.Request) error {
	//Currently, authentication means checking that the charging station
	//is already registered in the database
	station, err := GetChargingStationInfo(chargerId)
	if err != nil {
		return errors.New(fmt.Sprintf("error: authentication failed for charging station %s", chargerId))
	}

	log.Info("Charging station info: ", station)

	return nil
}

func GetChargingStationInfo(stationId string) (devices.ChargingStation, error) {
	deviceServiceHost := "host.docker.internal"
	if d := os.Getenv("DEVICE_SERVICE_HOST"); d != "" {
		deviceServiceHost = d
	}
	resp, err := http.Get(fmt.Sprintf("http://%s:5000/chargingstations/station/%s", deviceServiceHost, stationId))
	if err != nil {
		return devices.ChargingStation{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return devices.ChargingStation{}, err
	}
	var newCharger devices.ChargingStation
	error := json.Unmarshal(body, &newCharger)
	if error != nil {
		return devices.ChargingStation{}, err
	}
	return newCharger, nil
}
