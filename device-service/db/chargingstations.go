package db

import (
	"encoding/json"

	"github.com/gregszalay/firestore-go/firego"
	"github.com/gregszalay/ocpp-csms-common-types/devices"

	log "github.com/sirupsen/logrus"
)

var collection string = "chargingstations"

func GetChargingStation(id string) (devices.ChargingStation, error) {
	result, err := firego.Get(collection, id)

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
	list, err := firego.ListAll(collection)
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
	marshalled, marshal_err := json.Marshal(newCharger)
	if marshal_err != nil {
		log.Error("CreateTransaction marshal error: ", marshal_err)
	}
	var unmarshalled map[string]interface{}
	unmarshal_err := json.Unmarshal(marshalled, &unmarshalled)
	if unmarshal_err != nil {
		log.Error("CreateTransaction unmarshal error: ", unmarshal_err)
	}
	return firego.Create(collection, id, unmarshalled)
}

func UpdateChargingStation(id string, newCharger devices.ChargingStation) error {
	marshalled, marshal_err := json.Marshal(newCharger)
	if marshal_err != nil {
		log.Error("CreateTransaction marshal error: ", marshal_err)
	}
	var unmarshalled map[string]interface{}
	unmarshal_err := json.Unmarshal(marshalled, &unmarshalled)
	if unmarshal_err != nil {
		log.Error("CreateTransaction unmarshal error: ", unmarshal_err)
	}
	return firego.Update(collection, id, unmarshalled)
}

func DeleteChargingStation(id string) error {
	return firego.Delete(collection, id)
}
