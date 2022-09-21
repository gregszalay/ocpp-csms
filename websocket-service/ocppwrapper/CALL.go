package ocppwrapper

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type CALL struct {
	// This is a Message Type Number which is used to identify the type of the message.
	MessageTypeId float64
	// This must be the exact same id that is in the call request so that the recipient can match request and result.
	MessageId string
	// The name of the remote procedure or action. This field SHALL contain a case-sensitive string.
	// The field SHALL contain the OCPP Message name without the "Request" suffix. For example: For
	// a "BootNotificationRequest", this field shall be set to "BootNotification".
	Action string
	// JSON Payload of the action, see: JSON Payload for more information.
	Payload interface{}
}

func MakeCALLMessage(action string, payload interface{}) []byte {
	new_message_id := uuid.NewString()
	new_call := &CALL{
		MessageTypeId: CALL_TYPE,
		MessageId:     new_message_id,
		Action:        action,
		Payload:       payload,
	}
	return new_call.wrap()
}

func ParseCALLMessage(ocpp_message_json string) (CALL, error) {
	fmt.Println("Making CALL message object...")

	var data []interface{}
	err := json.Unmarshal([]byte(ocpp_message_json), &data)
	if err != nil {
		fmt.Printf("Error: could not unmarshal json: %s\n", err)
		return CALL{}, err
	}

	if len(data) != 4 {
		return CALL{}, errors.New("Error: invalid CALL message length!")
	}

	messageTypeId, ok := data[0].(float64)
	if !ok {
		return CALL{}, errors.New("Error: data[0] is not a float64!")
	}

	messageId, ok := data[1].(string)
	if !ok {
		return CALL{}, errors.New("Error: data[1] is not a string!")
	}

	action, ok := data[2].(string)
	if !ok {
		return CALL{}, errors.New("Error: data[2] is not a string!")

	}

	payload, ok := data[3].(map[string]interface{})
	if !ok {
		return CALL{}, errors.New("Error: data[3] is not an map[string]interface{}")
	}

	return CALL{
		MessageTypeId: messageTypeId,
		MessageId:     messageId,
		Action:        action,
		Payload:       payload,
	}, nil
}

func (c *CALL) wrap() []byte {
	message_array := [...]interface{}{CALL_TYPE, c.MessageId, c.Action, c.Payload}
	message_array_json, err := json.Marshal(message_array)
	if err != nil {
		fmt.Printf("Error: Could not marshal CALL message: %s\n", err)
		return []byte("")
	}
	fmt.Printf("CALL message: %s\n", message_array_json)
	return message_array_json
}

func (c *CALL) print() {
	fmt.Printf("CALL: %v\n", c)
}
