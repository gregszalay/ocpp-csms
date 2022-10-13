package ocpphandlers

import (
	"fmt"

	"github.com/gregszalay/ocpp-csms-common-types/QueuedMessage"
	"github.com/gregszalay/ocpp-csms/device-service/publishing"
	"github.com/gregszalay/ocpp-messages-go/types/StatusNotificationRequest"
	"github.com/gregszalay/ocpp-messages-go/types/StatusNotificationResponse"
	"github.com/sanity-io/litter"
)

func StatusNotificationHandler(request_json []byte, messageId string, deviceId string) {

	var req StatusNotificationRequest.StatusNotificationRequestJson
	payload_unmarshal_err := req.UnmarshalJSON(request_json)
	if payload_unmarshal_err != nil {
		fmt.Printf("Failed to unmarshal StatusNotificationRequest message payload. Error: %s", payload_unmarshal_err)
	} else {
		fmt.Println("Payload as an OBJECT:")
		litter.Dump(req)
	}

	resp := StatusNotificationResponse.StatusNotificationResponseJson{}

	qm := QueuedMessage.QueuedMessage{
		MessageId: messageId,
		DeviceId:  deviceId,
		Payload:   resp,
	}

	if err := publishing.Publish("StatusNotificationResponse", qm); err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
		panic(err)
	}

}
