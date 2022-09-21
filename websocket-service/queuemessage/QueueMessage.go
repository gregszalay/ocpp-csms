package queuemessage

import (
	"encoding/json"
	"errors"
	"fmt"
	"ocpp-websocket-service/ocppwrapper"
	"reflect"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/pkg/googlecloud"
	"github.com/ThreeDotsLabs/watermill/message"
)

type QueueMessage struct {
	MessageTypeId   int    `json:"messageTypeId"`
	MessageId       string `json:"messageId"`
	OCPPMessageJSON string `json:"ocppMessageJSONString"`
	ChargerId       string `json:"chargerId"`
}

var calls_awaiting_response map[string]ocppwrapper.CALL = map[string]ocppwrapper.CALL{}

// var amqpURI = "amqp://guest:guest@rabbit-manager:5672/"
// var amqpConfig = amqp.NewDurableQueueConfig(amqpURI)

func MakeQueueMessage(ocpp_message_json_string string, chargerId string) (QueueMessage, error) {
	fmt.Println("Making queue message...")
	fmt.Println(ocpp_message_json_string)

	var data []interface{}
	err := json.Unmarshal([]byte(ocpp_message_json_string), &data)
	if err != nil {
		fmt.Printf("Error: could not unmarshal json: %s\n", err)
		return QueueMessage{}, err
	}

	fmt.Println(reflect.TypeOf(data[0]))
	fmt.Println(data[0])
	fmt.Printf("Unmarshalles json: %+v\n", data)
	fmt.Printf(" data[0] : %+v\n", data)

	messageTypeId, ok := data[0].(float64)
	if !ok {
		fmt.Printf("Error: data[0] is not a uint8\n")
		return QueueMessage{}, errors.New("Not an OCPP message " + "Error: data[0] is not a uint8\n")
	}

	messageId, ok := data[1].(string)
	if !ok {
		fmt.Printf("Error: data[1] is not a string\n")
		return QueueMessage{}, errors.New("Not an OCPP message " + "Error: data[1] is not a string\n")
	}

	return QueueMessage{
		MessageTypeId:   int(messageTypeId),
		MessageId:       messageId,
		OCPPMessageJSON: ocpp_message_json_string,
		ChargerId:       chargerId,
	}, nil
}

func (qm *QueueMessage) Publish() {
	logger := watermill.NewStdLogger(true, true)

	publisher, err := googlecloud.NewPublisher(googlecloud.PublisherConfig{
		ProjectID: "chargerevolutioncloud",
	}, logger)
	if err != nil {
		panic(err)
	}

	topic, err := qm.getMQTopicName()
	if err != nil {
		panic(err)
	}
	fmt.Println("topic:")
	fmt.Println(topic)

	qm_json, err := json.Marshal(qm)
	if err != nil {
		fmt.Printf("Error: Could not marshal queue message: %s\n", err)
	}

	fmt.Printf("Marshalled QM: %+v\n", string(qm_json))

	msg := message.NewMessage(watermill.NewUUID(), qm_json)
	if err := publisher.Publish(topic, msg); err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
		panic(err)
	}

}

func (qm *QueueMessage) getMQTopicName() (string, error) {
	switch qm.MessageTypeId {
	case ocppwrapper.CALL_TYPE:
		call_m, err := ocppwrapper.ParseCALLMessage(qm.OCPPMessageJSON)
		if err != nil {
			return "", err
		}
		return call_m.Action + "Request", nil
	case ocppwrapper.CALLRESULT_TYPE:
		original_call_m := calls_awaiting_response[qm.MessageId]
		return original_call_m.Action + "Response", nil
	case ocppwrapper.CALLERROR_TYPE:
		original_call_m := calls_awaiting_response[qm.MessageId]
		return original_call_m.Action + "Error", nil
	}
	return "", errors.New("Error: wrong message type id!")
}
