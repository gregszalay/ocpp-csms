package ocpphandlers

import (
	"fmt"
	"time"

	"github.com/gregszalay/ocpp-csms-common-types/QueuedMessage"
	"github.com/gregszalay/ocpp-csms/device-service/db"
	"github.com/gregszalay/ocpp-csms/device-service/publishing"
	"github.com/gregszalay/ocpp-messages-go/types/BootNotificationRequest"
	"github.com/gregszalay/ocpp-messages-go/types/BootNotificationResponse"
	"github.com/sanity-io/litter"
	log "github.com/sirupsen/logrus"
)

func BootNotificationHandler(request_json []byte, messageId string, deviceId string) {

	var req BootNotificationRequest.BootNotificationRequestJson
	payload_unmarshal_err := req.UnmarshalJSON(request_json)
	if payload_unmarshal_err != nil {
		fmt.Printf("Failed to unmarshal BootNotificationRequest message payload. Error: %s", payload_unmarshal_err)
	} else {
		fmt.Println("Payload as an OBJECT:")
		litter.Dump(req)
	}

	resp_status := BootNotificationResponse.RegistrationStatusEnumType_1_Accepted
	station, err_get := db.GetChargingStation(deviceId)
	if err_get != nil {
		log.Error("Failed to get charging station from db. error: ", err_get)
		resp_status = BootNotificationResponse.RegistrationStatusEnumType_1_Rejected
	} else {
		station.LastBoot = time.Now().Format(time.RFC3339)

		err_update := db.UpdateChargingStation(deviceId, station)
		if err_update != nil {
			log.Error("Failed to update charging station. error: ", err_update)
		}
	}

	resp := BootNotificationResponse.BootNotificationResponseJson{
		CurrentTime: "",
		Interval:    60,
		Status:      resp_status,
	}

	qm := QueuedMessage.QueuedMessage{
		MessageId: messageId,
		DeviceId:  deviceId,
		Payload:   resp,
	}

	if err := publishing.Publish("BootNotificationResponse", qm); err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
		panic(err)
	}

}
