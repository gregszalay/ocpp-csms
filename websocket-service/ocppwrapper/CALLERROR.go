package ocppwrapper

import (
	"encoding/json"
	"errors"
	"fmt"
)

type CALLERROR struct {
	// This is a Message Type Number which is used to identify the type of the message.
	MessageTypeId float64
	// This must be the exact same id that is in the call request so that the recipient can match request and result.
	MessageId string
	// This field must contain a string from the RPC Framework Error Codes table.
	ErrorCode string
	// Should be filled in if possible, otherwise a clear empty string "".
	ErrorDescription string
	// This JSON object describes error details in an undefined way. If there are no error details you MUST fill in an empty object {}.
	ErrorDetails string
}

func MakeCALLERRORMessage(requestMessageId string, errorCode string, errorDescription string, errorDetails string) []byte {
	new_call := CALLERROR{
		MessageTypeId:    CALLERROR_TYPE,
		MessageId:        requestMessageId,
		ErrorCode:        errorCode,
		ErrorDescription: errorDescription,
		ErrorDetails:     errorDetails,
	}
	return new_call.wrap()
}

func ParseCALLERRORMessage(ocpp_message_json string) (CALLERROR, error) {
	fmt.Println("Making CALL RESULT message object...")

	var data []interface{}
	err := json.Unmarshal([]byte(ocpp_message_json), &data)
	if err != nil {
		fmt.Printf("Error: could not unmarshal json: %s\n", err)
		return CALLERROR{}, err
	}

	if len(data) != 5 {
		return CALLERROR{}, errors.New("Error: invalid CALL ERROR message length!")
	}

	messageTypeId, ok := data[0].(float64)
	if !ok {
		return CALLERROR{}, errors.New("Error: data[0] is not a float64!")
	}

	messageId, ok := data[1].(string)
	if !ok {
		return CALLERROR{}, errors.New("Error: data[1] is not a string!")
	}

	errorCode, ok := data[2].(string)
	if !ok {
		return CALLERROR{}, errors.New("Error: data[2] is not a string!")
	}

	errorDescription, ok := data[3].(string)
	if !ok {
		fmt.Printf("data[3] is not an string\n")
		return CALLERROR{}, errors.New("Error: data[3] is not a string!")
	}

	errorDetails, ok := data[4].(string)
	if !ok {
		fmt.Printf("data[3] is not an string\n")
		return CALLERROR{}, errors.New("Error: data[1] is not a string!")
	}

	return CALLERROR{
		MessageTypeId:    messageTypeId,
		MessageId:        messageId,
		ErrorCode:        errorCode,
		ErrorDescription: errorDescription,
		ErrorDetails:     errorDetails,
	}, nil

}

func (c *CALLERROR) wrap() []byte {
	message_array := [...]interface{}{CALLERROR_TYPE, c.MessageId, c.ErrorCode, c.ErrorDescription, c.ErrorDetails}
	message_array_json, err := json.Marshal(message_array)
	if err != nil {
		fmt.Printf("Error: Could not marshal CALLERROR message: %s\n", err)
		return []byte("")
	}
	fmt.Printf("CALLERROR message: %s\n", message_array_json)
	return message_array_json
}

func (c *CALLERROR) print() {
	fmt.Printf("CALLERROR: %v\n", c)
}

// ************** ENUMS ********************************************************
type rpc_framework_error_codes struct {
	// Payload for Action is syntactically incorrect
	FormatViolation string
	// Any other error not covered by the more specific error codes in this table
	GenericError string
	// An internal error occurred and the receiver was not able to process the requested Action successfully
	InternalError string
	// A message with an Message Type Number received that is not supported by this implementation.
	MessageTypeNotSupported string
	// Requested Action is not known by receiver
	NotImplemented string
	// Requested Action is recognized but not supported by the receiver
	NotSupported string
	// Payload for Action is syntactically correct but at least one of the fields violates occurrence constraints
	OccurrenceConstraintViolation string
	// Payload is syntactically correct but at least one field contains an invalid value
	PropertyConstraintViolation string
	// Payload for Action is not conform the PDU structure
	ProtocolError string
	// Content of the call is not a valid RPC Request, for example: MessageId could not be read.
	RpcFrameworkError string
	// During the processing of Action a security issue occurred preventing receiver from completing the Action successfully
	SecurityError string
	// Payload for Action is syntactically correct but at least one of the fields violates data type constraints (e.g. "somestring": 12)
	TypeConstraintViolation string
}

func init_rpc_framework_error_codes() *rpc_framework_error_codes {
	return &rpc_framework_error_codes{
		FormatViolation:               "FormatViolation",
		GenericError:                  "GenericError",
		InternalError:                 "InternalError",
		MessageTypeNotSupported:       "MessageTypeNotSupported",
		NotImplemented:                "NotImplemented",
		NotSupported:                  "NotSupported",
		OccurrenceConstraintViolation: "OccurrenceConstraintViolation",
		PropertyConstraintViolation:   "PropertyConstraintViolation",
		ProtocolError:                 "ProtocolError",
		RpcFrameworkError:             "RpcFrameworkError",
		SecurityError:                 "SecurityError",
		TypeConstraintViolation:       "TypeConstraintViolation",
	}
}

var RPC_FRAMEWORK_ERROR_CODES = init_rpc_framework_error_codes()
