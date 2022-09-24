package chargingstations

import (
	"encoding/json"
	"fmt"

	types "github.com/gregszalay/ocpp-csms-common-types/devices"
	"github.com/gregszalay/ocpp-csms/device-service/devicedb"

	"github.com/fatih/structs"
)

var collection string = "devices"

func GetCharger(id string) *types.Charger {
	result := devicedb.Get(collection, id)
	// Convert map to json string
	jsonStr, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	// Convert json string to struct
	var charger types.Charger
	if err := json.Unmarshal(jsonStr, &charger); err != nil {
		fmt.Println(err)
	}
	return &charger
}

func ListChargers() []types.Charger {
	list := devicedb.ListAll(collection)
	chargerList := []types.Charger{}
	for _, value := range list {
		fmt.Println("***")
		jsonStr, err := json.Marshal(value)
		if err != nil {
			fmt.Println(err)
		}
		// Convert json string to struct
		var charger types.Charger
		if err := json.Unmarshal(jsonStr, &charger); err != nil {
			fmt.Println(err)
		}
		chargerList = append(chargerList, charger)
	}
	return chargerList
}

func CreateCharger(newCharger *types.Charger, newId string) {
	devicedb.Create(structs.Map(newCharger), collection, newId)
}

func UpdateCharger(newCharger *types.Charger, newId string) {
	devicedb.Update(structs.Map(newCharger), collection, newId)
}

func DeleteCharger(newCharger *types.Charger, newId string) {
	devicedb.Delete(structs.Map(newCharger), collection, newId)
}
