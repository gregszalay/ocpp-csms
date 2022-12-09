package messagemux

import (
	"encoding/json"

	"github.com/gregszalay/ocpp-csms-common-types/QueuedError"
	"github.com/gregszalay/ocpp-csms-common-types/QueuedMessage"
	"github.com/gregszalay/ocpp-csms/websocket-service/publishing"
	"github.com/gregszalay/ocpp-messages-go/wrappers"
	log "github.com/sirupsen/logrus"
)

var calls_awaiting_response map[string]wrappers.CALL = map[string]wrappers.CALL{}

func ProcessAndPublish(stationId string, message []byte) error {
	messageTypeId, err := parseMessageTypeId(stationId, message)
	if err != nil {
		log.Error("could not parse message type id")
		return err
	}

	switch messageTypeId {
	case wrappers.CALL_TYPE:
		return process_as_CALL(stationId, message)
	case wrappers.CALLRESULT_TYPE:
		return process_as_CALLRESULT(stationId, message)
	case wrappers.CALLERROR_TYPE:
		return process_as_CALLERROR(stationId, message)
	}
	return nil
}

// Processes the the incoming message (the receiver type) as a CALL message
func process_as_CALL(stationId string, message []byte) error {
	var call wrappers.CALL
	err := call.UnmarshalJSON([]byte(message))
	if err != nil {
		log.Error("failed to unmarshal OCPP CALL message. Error: %s", err)
		return err
	}
	qm := QueuedMessage.QueuedMessage{
		MessageId: call.MessageId,
		DeviceId:  stationId,
		Payload:   call.Payload,
	}
	call_topic := call.Action + "Request"
	// We publish the CALLRESULT to the relevant pupbsub topic
	if err := publishing.Publish(call_topic, qm); err != nil {
		log.Error(err)
	}
	return nil
}

// Processes the the incoming message (the receiver type) as a CALLRESULT message
// These are messages that are repsonses to CALLs we have sent out to the charging station
func process_as_CALLRESULT(stationId string, message []byte) error {
	var callresult wrappers.CALLRESULT
	err := callresult.UnmarshalJSON(message)
	if err != nil {
		log.Error("failed to unmarshal OCPP CALLRESULT message. Error: ", err)
	}
	qm := QueuedMessage.QueuedMessage{
		MessageId: callresult.MessageId,
		DeviceId:  stationId,
		Payload:   callresult.Payload,
	}
	// We retrieve the original CALL message that we previously sent out to the charging station
	// so that we know which topic to put the response in
	original_call_message := calls_awaiting_response[callresult.MessageId]
	callresult_topic := original_call_message.Action + "Response"
	// We publish the CALLRESULT to the relevant upbsub topic
	if err := publishing.Publish(callresult_topic, qm); err != nil {
		log.Error(err)
	}
	return nil
}

// Processes the the incoming message (the receiver type) as a CALLERROR message
func process_as_CALLERROR(stationId string, message []byte) error {
	var callerror wrappers.CALLERROR
	err := callerror.UnmarshalJSON([]byte(message))
	if err != nil {
		log.Error("failed to unmarshal OCPP CALLERROR message. Error: %s", err)
		return err
	}
	qm := QueuedError.QueuedError{
		MessageId:        callerror.MessageId,
		DeviceId:         stationId,
		ErrorCode:        callerror.ErrorCode,
		ErrorDescription: callerror.ErrorDescription,
		ErrorDetails:     callerror.ErrorDetails,
	}
	// We retrieve the original CALL message that we previously sent out to the charging station
	// so that we know which topic to put the error response in
	original_call_message := calls_awaiting_response[callerror.MessageId]
	callerror_topic := original_call_message.Action + "Error"
	// We publish the CALLRESULT to the relevant upbsub topic
	if err := publishing.Publish(callerror_topic, qm); err != nil {
		log.Error(err)
	}
	return nil

}

func parseMessageTypeId(stationId string, message []byte) (int, error) {
	var data []interface{}
	err := json.Unmarshal([]byte(message), &data)
	if err != nil {
		log.Error("could not unmarshal json", err)
		return 0, err
	}
	messageTypeId, ok := data[0].(float64)
	if !ok {
		log.Error("data[0] is not a uint8", err)
		return 0, err
	}
	return int(messageTypeId), nil
}
