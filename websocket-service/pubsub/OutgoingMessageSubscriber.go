package pubsub

import (
	"context"
	"fmt"
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/pkg/googlecloud"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gregszalay/ocpp-csms-common-types/QueuedMessage"
)

var out_topics []string = []string{
	"BootNotificationResponse",
}

//var Subs map[string]<-chan *message.Message = map[string]<-chan *message.Message{}
var ToChargerQueue map[string]chan *QueuedMessage.QueuedMessage = map[string]chan *QueuedMessage.QueuedMessage{}

func Subscribe() {

	logger := watermill.NewStdLogger(true, true)
	subscriber, err := googlecloud.NewSubscriber(
		googlecloud.SubscriberConfig{
			// custom function to generate Subscription Name,
			// there are also predefined TopicSubscriptionName and TopicSubscriptionNameWithSuffix available.
			GenerateSubscriptionName: func(topic string) string {
				return "websocket-service_" + topic
			},
			ProjectID: "chargerevolutioncloud",
		},
		logger,
	)
	if err != nil {
		panic(err)
	}

	for _, topic := range out_topics {
		// Subscribe will create the subscription. Only messages that are sent after the subscription is created may be received.
		messages, err := subscriber.Subscribe(context.Background(), topic)
		if err != nil {
			panic(err)
		}
		go process(topic, messages)
	}
}

func process(topic string, messages <-chan *message.Message) {
	for msg := range messages {

		log.Printf("received message: %s, topic: %s, payload: %s", msg.UUID, topic, string(msg.Payload))

		var qm QueuedMessage.QueuedMessage
		err := qm.UnmarshalJSON(msg.Payload)
		if err != nil {
			fmt.Printf("Failed to unmarshal QueuedMessage message. Error: %s", err)
		}

		// fmt.Println("QueuedMessage as an OBJECT:")
		// litter.Dump(qm)
		//
		if ToChargerQueue[qm.DeviceId] == nil {
			msg.Ack()
			continue
		}
		fmt.Println("Putting msg into ToChargerQueue")
		ToChargerQueue[qm.DeviceId] <- &qm

		// we need to Acknowledge that we received and processed the message,
		// otherwise, it will be resent over and over again.
		msg.Ack()
	}
}
