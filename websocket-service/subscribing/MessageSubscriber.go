package subscribing

import (
	"context"
	"os"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/pkg/googlecloud"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gregszalay/ocpp-csms-common-types/QueuedMessage"
	"github.com/gregszalay/ocpp-csms/websocket-service/websocketserver"
	log "github.com/sirupsen/logrus"
)

var PROJECT_ID string = os.Getenv("GCP_PROJECT_ID")
var SERVICE_APP_NAME string = os.Getenv("SERVICE_APP_NAME")

var out_topics []string = []string{
	"BootNotificationResponse",
	"AuthorizeResponse",
	"TransactionEventResponse",
	"HeartbeatResponse",
	"StatusNotificationResponse",
}

func Subscribe() {

	logger := watermill.NewStdLogger(true, true)
	subscriber, err := googlecloud.NewSubscriber(
		googlecloud.SubscriberConfig{
			// custom function to generate Subscription Name,
			// there are also predefined TopicSubscriptionName and TopicSubscriptionNameWithSuffix available.
			GenerateSubscriptionName: func(topic string) string {
				return SERVICE_APP_NAME + "_" + topic
			},
			ProjectID: PROJECT_ID,
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

		log.Info("received message: %s, topic: %s, payload: %s", msg.UUID, topic, string(msg.Payload))

		var qm QueuedMessage.QueuedMessage
		err := qm.UnmarshalJSON(msg.Payload)
		if err != nil {
			log.Error("failed to unmarshal QueuedMessage message. Error: %s", err)
		}

		if websocketserver.AllMessagesToDeviceMap[qm.DeviceId] == nil {
			msg.Ack()
			continue
		}
		log.Debug("Putting msg into MessagesToDevice")
		websocketserver.AllMessagesToDeviceMap[qm.DeviceId] <- &qm

		// we need to Acknowledge that we received and processed the message,
		// otherwise, it will be resent over and over again.
		msg.Ack()
	}
}
