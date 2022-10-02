package db

import (
	"encoding/json"

	"github.com/gregszalay/ocpp-csms-common-types/devices"

	"github.com/fatih/structs"
	log "github.com/sirupsen/logrus"
)

var collection string = "chargingstations"

func GetChargingStation(id string) (devices.ChargingStation, error) {
	result, err := get(collection, id)

	jsonStr, err_marshal := json.Marshal(result)
	if err_marshal != nil {
		log.Error("failed to marshal charging station", err_marshal)
	}

	var charger devices.ChargingStation
	if err_unmarshal := json.Unmarshal(jsonStr, &charger); err != nil {
		log.Error("failed to unmarshal charging station json", err_unmarshal)

	}
	return charger, err
}

func ListChargingStations() (*[]devices.ChargingStation, error) {
	list, err := listAll(collection)
	chargerList := []devices.ChargingStation{}
	for index, value := range *list {
		jsonStr, err := json.Marshal(value)
		if err != nil {
			log.Error("failed to marshal charging station list element ", index, " error: ", err)
		}
		var charger devices.ChargingStation
		if err := json.Unmarshal(jsonStr, &charger); err != nil {
			log.Error("failed to unmarshal charging station list element ", index, " error: ", err)
		}
		chargerList = append(chargerList, charger)
	}
	log.Debug("List of charging stations: ", chargerList)
	return &chargerList, err
}

func CreateChargingStation(id string, newCharger devices.ChargingStation) error {
	return create(collection, id, structs.Map(newCharger))
}

func UpdateChargingStation(id string, newCharger devices.ChargingStation) error {
	return update(collection, id, structs.Map(newCharger))
}

func DeleteChargingStation(id string) error {
	return delete(collection, id)
}
