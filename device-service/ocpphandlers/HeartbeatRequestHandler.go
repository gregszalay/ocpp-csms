package ocpphandlers

import (
	"fmt"
	"time"

	"github.com/gregszalay/ocpp-csms-common-types/QueuedMessage"
	"github.com/gregszalay/ocpp-csms/device-service/publishing"
	"github.com/gregszalay/ocpp-messages-go/types/HeartbeatResponse"
)

func HeartbeatRequestHandler(request_json []byte, messageId string, deviceId string) {

	//TODO: implement unmarshallign of Heartbeatrequest, if needed

	time_now_RFC3339 := time.Now().Format(time.RFC3339)
	//time, err := time.Parse( time.RFC3339, "2012-11-01T22:08:41+00:00")
	resp := HeartbeatResponse.HeartbeatResponseJson{
		CurrentTime: time_now_RFC3339,
	}

	qm := QueuedMessage.QueuedMessage{
		MessageId: messageId,
		DeviceId:  deviceId,
		Payload:   resp,
	}

	if err := publishing.Publish("HeartbeatResponse", qm); err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
		panic(err)
	}

}
