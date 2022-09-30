package MessageIn

import (
	"encoding/json"
	"fmt"

	"github.com/gregszalay/ocpp-csms-common-types/QueuedError"
	"github.com/gregszalay/ocpp-csms-common-types/QueuedMessage"
	"github.com/gregszalay/ocpp-csms/websocket-service/pub"
	"github.com/gregszalay/ocpp-messages-go/wrappers"
)

var calls_awaiting_response map[string]wrappers.CALL = map[string]wrappers.CALL{}

type MessageIn struct {
	ChargerId string
	Message   []byte
}

func (in *MessageIn) Process() error {
	fmt.Println("Processing incoming message...")

	messageTypeId, err := in.parseMessageTypeId()
	if err != nil {
		fmt.Printf("error: could not parse message type id\n")
		return err
	}

	switch messageTypeId {
	case wrappers.CALL_TYPE:
		return in.process_as_CALL()
	case wrappers.CALLRESULT_TYPE:
		return in.process_as_CALLRESULT()
	case wrappers.CALLERROR_TYPE:
		return in.process_as_CALLERROR()
	}
	return nil
}

// Processes the the incoming message (the receiver type) as a CALL message
func (in *MessageIn) process_as_CALL() error {
	var call wrappers.CALL
	err := call.UnmarshalJSON([]byte(in.Message))
	if err != nil {
		fmt.Printf("Failed to unmarshal OCPP CALL message. Error: %s", err)
		return err
	}
	qm := QueuedMessage.QueuedMessage{
		MessageId: call.MessageId,
		DeviceId:  in.ChargerId,
		Payload:   call.Payload,
	}
	call_topic := call.Action + "Request"
	// We publish the CALLRESULT to the relevant upbsub topic
	if err := pub.Publish(call_topic, qm); err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
		panic(err)
	}
	return nil
}

// Processes the the incoming message (the receiver type) as a CALLRESULT message
// These are messages that are repsonses to CALLs we have sent out to the charging station
func (in *MessageIn) process_as_CALLRESULT() error {
	var callresult wrappers.CALLRESULT
	err := callresult.UnmarshalJSON(in.Message)
	if err != nil {
		fmt.Printf("Failed to unmarshal OCPP CALLRESULT message. Error: %s", err)
	}
	qm := QueuedMessage.QueuedMessage{
		MessageId: callresult.MessageId,
		DeviceId:  in.ChargerId,
		Payload:   callresult.Payload,
	}
	// We retrieve the original CALL message that we previously sent out to the charging station
	// so that we know which topic to put the response in
	original_call_message := calls_awaiting_response[callresult.MessageId]
	callresult_topic := original_call_message.Action + "Response"
	// We publish the CALLRESULT to the relevant upbsub topic
	if err := pub.Publish(callresult_topic, qm); err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
		panic(err)
	}
	return nil
}

// Processes the the incoming message (the receiver type) as a CALLERROR message
func (in *MessageIn) process_as_CALLERROR() error {
	var callerror wrappers.CALLERROR
	err := callerror.UnmarshalJSON([]byte(in.Message))
	if err != nil {
		fmt.Printf("Failed to unmarshal OCPP CALLERROR message. Error: %s", err)
		return err
	}
	qm := QueuedError.QueuedError{
		MessageId:        callerror.MessageId,
		DeviceId:         in.ChargerId,
		ErrorCode:        callerror.ErrorCode,
		ErrorDescription: callerror.ErrorDescription,
		ErrorDetails:     callerror.ErrorDetails,
	}
	// We retrieve the original CALL message that we previously sent out to the charging station
	// so that we know which topic to put the error response in
	original_call_message := calls_awaiting_response[callerror.MessageId]
	callerror_topic := original_call_message.Action + "Error"
	// We publish the CALLRESULT to the relevant upbsub topic
	if err := pub.Publish(callerror_topic, qm); err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
		panic(err)
	}
	return nil

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
