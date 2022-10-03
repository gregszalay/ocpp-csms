package ocpphandlers

import (
	"fmt"

	"github.com/gregszalay/ocpp-csms-common-types/QueuedMessage"
	"github.com/gregszalay/ocpp-csms/transaction-service/db"
	"github.com/gregszalay/ocpp-csms/transaction-service/publishing"
	"github.com/gregszalay/ocpp-messages-go/types/TransactionEventRequest"
	"github.com/gregszalay/ocpp-messages-go/types/TransactionEventResponse"
	"github.com/sanity-io/litter"
	log "github.com/sirupsen/logrus"
)

func TransactionEventHandler(request_json []byte, messageId string, deviceId string) {

	var req TransactionEventRequest.TransactionEventRequestJson
	payload_unmarshal_err := req.UnmarshalJSON(request_json)
	if payload_unmarshal_err != nil {
		fmt.Printf("Failed to unmarshal TransactionEventRequest message payload. Error: %s", payload_unmarshal_err)
	} else {
		if log.GetLevel() == log.DebugLevel {
			log.Info("Payload as an OBJECT: ")
			litter.Dump(req)
		}
	}

	switch req.EventType {
	case TransactionEventRequest.TransactionEventEnumType_1_Started:
		db_id := req.TransactionInfo.TransactionId
		db.CreateTransaction(db_id, db.Transaction{
			EnergyTransferInProgress: true,
			EnergyTransferStarted:    req.Timestamp,
			EnergyTransferStopped:    "",
			MeterValues:              req.MeterValue,
		})
	case TransactionEventRequest.TransactionEventEnumType_1_Updated:
		currentTx, err := db.GetTransaction(req.TransactionInfo.TransactionId)
		if err != nil {
			log.Error("failed to get previous transaction info from db")
		}
		fmt.Println("litter.Dump(currentTx.MeterValues)")
		litter.Dump(currentTx.MeterValues)
		originalMeterValues := currentTx.MeterValues
		fmt.Println("len(originalMeterValues)")
		fmt.Println(len(originalMeterValues))
		fmt.Println("cap(originalMeterValues)")
		fmt.Println(cap(originalMeterValues))
		latestMeterValues := req.MeterValue
		fmt.Println("len(latestMeterValues)")
		fmt.Println(len(latestMeterValues))
		fmt.Println("cap(latestMeterValues)")
		fmt.Println(cap(latestMeterValues))
		newMeterValues := append(originalMeterValues, latestMeterValues...)
		fmt.Println("len(newMeterValues)")
		fmt.Println(len(newMeterValues))
		fmt.Println("cap(newMeterValues)")
		fmt.Println(cap(newMeterValues))

		db_id := req.TransactionInfo.TransactionId
		db.UpdateTransaction(db_id, db.Transaction{
			EnergyTransferInProgress: true,
			EnergyTransferStarted:    currentTx.EnergyTransferStarted,
			EnergyTransferStopped:    currentTx.EnergyTransferStopped,
			MeterValues:              newMeterValues,
		})
	case TransactionEventRequest.TransactionEventEnumType_1_Ended:
		currentTx, err := db.GetTransaction(req.TransactionInfo.TransactionId)
		if err != nil {
			log.Error("failed to get previous transaction info from db")
		}
		originalMeterValues := currentTx.MeterValues
		fmt.Println("len(originalMeterValues)")
		fmt.Println(len(originalMeterValues))
		fmt.Println("cap(originalMeterValues)")
		fmt.Println(cap(originalMeterValues))
		latestMeterValues := req.MeterValue
		fmt.Println("len(latestMeterValues)")
		fmt.Println(len(latestMeterValues))
		fmt.Println("cap(latestMeterValues)")
		fmt.Println(cap(latestMeterValues))
		newMeterValues := append(originalMeterValues, latestMeterValues...)
		fmt.Println("len(newMeterValues)")
		fmt.Println(len(newMeterValues))
		fmt.Println("cap(newMeterValues)")
		fmt.Println(cap(newMeterValues))

		db_id := req.TransactionInfo.TransactionId
		db.UpdateTransaction(db_id, db.Transaction{
			EnergyTransferInProgress: false,
			EnergyTransferStarted:    currentTx.EnergyTransferStarted,
			EnergyTransferStopped:    req.Timestamp,
			MeterValues:              newMeterValues,
		})
	}

	//RFC3339_time := p(time.Now().Format(time.RFC3339))
	//time, err := time.Parse( time.RFC3339, "2012-11-01T22:08:41+00:00")

	resp := TransactionEventResponse.TransactionEventResponseJson{
		UpdatedPersonalMessage: &TransactionEventResponse.MessageContentType{
			Format:  TransactionEventResponse.MessageFormatEnumTypeUTF8,
			Content: "Charging is in progress, your current bill is $5.00",
		},
	}

	qm := QueuedMessage.QueuedMessage{
		MessageId: messageId,
		DeviceId:  deviceId,
		Payload:   resp,
	}

	if err := publishing.Publish("TransactionEventResponse", qm); err != nil {
		log.Error("failed to publish TransactionEventResponse")
		log.Error(err)
	}

}
