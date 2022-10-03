package db

import "github.com/gregszalay/ocpp-messages-go/types/TransactionEventRequest"

type Transaction struct {
	EnergyTransferInProgress bool                                     `json:"energyTransferInProgress" yaml:"energyTransferInProgress"`
	EnergyTransferStarted    string                                   `json:"energyTransferStarted" yaml:"energyTransferStarted"`
	EnergyTransferStopped    string                                   `json:"energyTransferStopped" yaml:"energyTransferStopped"`
	MeterValues              []TransactionEventRequest.MeterValueType `json:"meterValues" yaml:"meterValues"`
}
