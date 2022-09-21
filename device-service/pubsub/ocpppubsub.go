package pubsub

import (
	"context"
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/pkg/googlecloud"
	"github.com/ThreeDotsLabs/watermill/message"
)

var relevant_topics []string = []string{
	"BootNotificationRequest",
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

	for _, topic := range relevant_topics {
		// Subscribe will create the subscription. Only messages that are sent after the subscription is created may be received.
		messages, err := subscriber.Subscribe(context.Background(), topic)
		if err != nil {
			panic(err)
		}
		subs = append(subs, messages)
		go process(topic, messages)
	}
}

// func response() {
// 	ocpp_messages.BootNotificationResponseJson
// 	currentTime
// }
// }

func process(topic string, messages <-chan *message.Message) {
	for msg := range messages {
		log.Printf("received message: %s, topic: %s, payload: %s", msg.UUID, topic, string(msg.Payload))

		// we need to Acknowledge that we received and processed the message,
		// otherwise, it will be resent over and over again.
		msg.Ack()
	}
}
