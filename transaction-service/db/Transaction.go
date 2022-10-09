package db

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gregszalay/ocpp-messages-go/types/TransactionEventRequest"
	log "github.com/sirupsen/logrus"
)

type Transaction struct {
	StationId                string                                   `json:"stationId" yaml:"stationId"`
	EnergyTransferInProgress bool                                     `json:"energyTransferInProgress" yaml:"energyTransferInProgress"`
	EnergyTransferStarted    string                                   `json:"energyTransferStarted" yaml:"energyTransferStarted"`
	EnergyTransferStopped    string                                   `json:"energyTransferStopped" yaml:"energyTransferStopped"`
	MeterValues              []TransactionEventRequest.MeterValueType `json:"meterValues" yaml:"meterValues"`
}

func (j *Transaction) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	err := json.Unmarshal(b, &raw)
	if err != nil {
		log.Error("could not unmarshal jso")
		fmt.Printf("could not unmarshal json: %s\n", err)
		return err
	}

	device_id, err_device_id := raw["deviceId"].(string)
	if !err_device_id {
		return errors.New("field deviceId is not a string")
	}
	energyTransferInProgress, err_energyTransferInProgress := raw["energyTransferInProgress"].(bool)
	if !err_energyTransferInProgress {
		return errors.New("field EnergyTransferInProgress is not a bool")
	}

	energyTransferStarted, err_energyTransferStarted := raw["energyTransferStarted"].(string)
	if !err_energyTransferStarted {
		return errors.New("field EnergyTransferStarted is not a string")
	}

	energyTransferStopped, err_energyTransferStopped := raw["energyTransferStopped"].(string)
	if !err_energyTransferStopped {
		return errors.New("field energyTransferStopped is not a string")
	}

	meterValues, err_meterValues := raw["meterValues"].([]interface{})
	if !err_meterValues {
		return errors.New("field MeterValues is not a []interface{}")
	}

	var decoded_meterValues []TransactionEventRequest.MeterValueType
	for _, value := range meterValues {
		var meter_value_element TransactionEventRequest.MeterValueType
		marshalled_element, marshal_element_err := json.Marshal(value)
		if marshal_element_err != nil {
			log.Error(marshal_element_err)
			return errors.New("failed to re-marshal raw MeterValues element")
		}
		err_elem := meter_value_element.UnmarshalJSON(marshalled_element)
		if err_elem != nil {
			log.Error(err_elem)
			return errors.New("MeterValues element is not a TransactionEventRequest.MeterValueType")
		}
		decoded_meterValues = append(decoded_meterValues, meter_value_element)
	}

	*j = Transaction{
		StationId:                device_id,
		EnergyTransferInProgress: energyTransferInProgress,
		EnergyTransferStarted:    energyTransferStarted,
		EnergyTransferStopped:    energyTransferStopped,
		MeterValues:              decoded_meterValues,
	}

	return nil
}
