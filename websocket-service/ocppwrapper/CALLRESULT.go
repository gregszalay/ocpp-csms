package ocppwrapper

import (
	"encoding/json"
	"errors"
	"fmt"
)

type CALLRESULT struct {
	// This is a Message Type Number which is used to identify the type of the message.
	MessageTypeId float64
	// This must be the exact same id that is in the call request so that the recipient can match request and result.
	MessageId string
	// JSON Payload of the action, see: JSON Payload for more information.
	Payload interface{}
}

func MakeCALLRESULTMessage(requestMessageId string, payload interface{}) []byte {
	new_call := &CALLRESULT{
		MessageTypeId: CALLRESULT_TYPE,
		MessageId:     requestMessageId,
		Payload:       payload,
	}
	return new_call.wrap()
}

func ParseCALLRESULTMessage(ocpp_message_json string) (CALLRESULT, error) {
	fmt.Println("Making CALL RESULT message object...")

	var data []interface{}
	err := json.Unmarshal([]byte(ocpp_message_json), &data)
	if err != nil {
		fmt.Printf("Error: could not unmarshal json: %s\n", err)
		return CALLRESULT{}, err
	}

	if len(data) != 3 {
		return CALLRESULT{}, errors.New("Error: invalid CALL RESULT message length!")
	}

	messageTypeId, ok := data[0].(float64)
	if !ok {
		return CALLRESULT{}, errors.New("Error: data[0] is not a float64!")
	}

	messageId, ok := data[1].(string)
	if !ok {
		return CALLRESULT{}, errors.New("Error: data[1] is not a string!")
	}

	payload, ok := data[2].(map[string]interface{})
	if !ok {
		return CALLRESULT{}, errors.New("Error: data[3] is not an map[string]interface{}")
	}

	return CALLRESULT{
		MessageTypeId: messageTypeId,
		MessageId:     messageId,
		Payload:       payload,
	}, nil

}

func (c *CALLRESULT) wrap() []byte {
	message_array := [...]interface{}{CALLRESULT_TYPE, c.MessageId, c.Payload}
	message_array_json, err := json.Marshal(message_array)
	if err != nil {
		fmt.Printf("Error: Could not marshal CALLRESULT message: %s\n", err)
		return []byte("")
	}
	fmt.Printf("CALLRESULT message: %s\n", message_array_json)
	return message_array_json
}

func (c *CALLRESULT) print() {
	fmt.Printf("CALLRESULT: %v\n", c)
}
