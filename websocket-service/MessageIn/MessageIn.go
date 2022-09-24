package MessageIn

import (
	"encoding/json"
	"fmt"

	"github.com/gregszalay/ocpp-csms-common-types/QueuedMessage"
	"github.com/gregszalay/ocpp-csms/websocket-service/pub"
	"github.com/gregszalay/ocpp-messages-go/wrappers"
)

var calls_awaiting_response map[string]wrappers.CALL = map[string]wrappers.CALL{}

type MessageIn struct {
	chargerId string
	Message   []byte
}

func (in *MessageIn) Process() error {
	fmt.Println("Processing incoming message...")
	fmt.Println(in.Message)

	messageTypeId, err := in.parseMessageTypeId()
	if err != nil {
		fmt.Printf("error: could not parse message type id\n")
		return err
	}

	switch messageTypeId {
	case wrappers.CALL_TYPE:
		in.process_as_CALL()
	case wrappers.CALLRESULT_TYPE:
		in.process_as_CALLRESULT()
	case wrappers.CALLERROR_TYPE:
		in.process_as_CALLERROR()
	}
}

func (in *MessageIn) process_as_CALL() error {
	var call wrappers.CALL
	call_unmarshal_err := call.UnmarshalJSON([]byte(in.Message))
	if call_unmarshal_err != nil {
		fmt.Printf("Failed to unmarshal OCPP CALL message. Error: %s", call_unmarshal_err)
		return call_unmarshal_err
	}
	qm := QueuedMessage.QueuedMessage{
		MessageId: call.MessageId,
		DeviceId:  in.chargerId,
		Payload:   call.Payload,
	}
	pub.Publish(qm, call.Action+"Request")
}

func (in *MessageIn) process_as_CALLRESULT() error {
	var callresult wrappers.CALLRESULT
	call_result_unmarshal_err := callresult.UnmarshalJSON(in.Message)
	if call_result_unmarshal_err != nil {
		fmt.Printf("Failed to unmarshal OCPP CALLRESULT message. Error: %s", call_result_unmarshal_err)
	}
	qm := QueuedMessage.QueuedMessage{
		MessageId: callresult.MessageId,
		DeviceId:  in.chargerId,
		Payload:   callresult.Payload,
	}
	original_call_m := calls_awaiting_response[callresult.MessageId]
	pub.Publish(qm, original_call_m.Action+"Response")
}

func (in *MessageIn) process_as_CALLERROR() error {

}

func (in *MessageIn) parseMessageTypeId() (int, error) {
	var data []interface{}
	err := json.Unmarshal([]byte(in.Message), &data)
	if err != nil {
		fmt.Printf("Error: could not unmarshal json: %s\n", err)
		return 0, err
	}
	messageTypeId, ok := data[0].(float64)
	if !ok {
		fmt.Printf("Error: data[0] is not a uint8\n")
		return 0, err
	}
	return int(messageTypeId), nil
}
