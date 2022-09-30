package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/pkg/googlecloud"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gregszalay/ocpp-csms-common-types/QueuedMessage"
	"github.com/gregszalay/ocpp-csms/device-service/ocpphandlers"
	"github.com/sanity-io/litter"
)

var call_topics map[string]func([]byte, string, string) = map[string]func([]byte, string, string){
	"BootNotificationRequest": ocpphandlers.BootNotificationHandler,
}

var subs []<-chan *message.Message = []<-chan *message.Message{}

func Subscribe() {

	logger := watermill.NewStdLogger(true, true)
	subscriber, err := googlecloud.NewSubscriber(
		googlecloud.SubscriberConfig{
			// custom function to generate Subscription Name,
			// there are also predefined TopicSubscriptionName and TopicSubscriptionNameWithSuffix available.
			GenerateSubscriptionName: func(topic string) string {
				return "test-sub_" + topic
			},
			ProjectID: "chargerevolutioncloud",
		},
		logger,
	)
	if err != nil {
		panic(err)
	}

	for topic, handler := range call_topics {
		// Subscribe will create the subscription. Only messages that are sent after the subscription is created may be received.
		messages, err := subscriber.Subscribe(context.Background(), topic)
		if err != nil {
			panic(err)
		}
		subs = append(subs, messages)
		go process(topic, messages, handler)
	}
}

// func response() {
// 	ocpp_messages.BootNotificationResponseJson
// 	currentTime
// }
// }

func process(topic string, messages <-chan *message.Message, fn func([]byte, string, string)) {
	for msg := range messages {
		log.Printf("received message: %s, topic: %s, payload: %s", msg.UUID, topic, string(msg.Payload))

		var qm QueuedMessage.QueuedMessage
		err := qm.UnmarshalJSON(msg.Payload)
		if err != nil {
			fmt.Printf("Failed to unmarshal OCPP CALLRESULT message. Error: %s", err)
		}

		fmt.Println("QueuedMessage as an OBJECT:")
		litter.Dump(qm)

		fmt.Println("CALL as an OBJECT:")
		litter.Dump(qm.Payload)

		//re-marshal
		result, err := json.Marshal(qm.Payload)
		if err != nil {
			fmt.Printf("Could not re-marshal OCPP payload: %s\n", err)
		}
		fn(result, qm.MessageId, qm.DeviceId)

		// we need to Acknowledge that we received and processed the message,
		// otherwise, it will be resent over and over again.
		msg.Ack()
	}
}
