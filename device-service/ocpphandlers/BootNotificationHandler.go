package ocpphandlers

import (
	"fmt"

	"github.com/gregszalay/ocpp-csms-common-types/QueuedMessage"
	"github.com/gregszalay/ocpp-csms/device-service/publishing"
	"github.com/gregszalay/ocpp-messages-go/types/BootNotificationRequest"
	"github.com/gregszalay/ocpp-messages-go/types/BootNotificationResponse"
	"github.com/sanity-io/litter"
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

	resp := BootNotificationResponse.BootNotificationResponseJson{
		CurrentTime: "",
		Interval:    60,
		Status:      BootNotificationResponse.RegistrationStatusEnumType_1_Accepted,
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
