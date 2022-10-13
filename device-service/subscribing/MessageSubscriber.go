package subscribing

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/pkg/googlecloud"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gregszalay/ocpp-csms-common-types/QueuedMessage"
	"github.com/gregszalay/ocpp-csms/device-service/ocpphandlers"
	log "github.com/sirupsen/logrus"
)

var PROJECT_ID string = os.Getenv("GCP_PROJECT_ID")
var SERVICE_APP_NAME string = os.Getenv("SERVICE_APP_NAME")

var call_topics map[string]func([]byte, string, string) = map[string]func([]byte, string, string){
	"BootNotificationRequest":   ocpphandlers.BootNotificationHandler,
	"HeartbeatRequest":          ocpphandlers.HeartbeatRequestHandler,
	"StatusNotificationRequest": ocpphandlers.StatusNotificationHandler,
}

var subs []<-chan *message.Message = []<-chan *message.Message{}

func Subscribe() {

	logger := watermill.NewStdLogger(true, true)
	subscriber, err := googlecloud.NewSubscriber(
		googlecloud.SubscriberConfig{
			// custom function to generate Subscription Name,
			// there are also predefined TopicSubscriptionName and TopicSubscriptionNameWithSuffix available.
			GenerateSubscriptionName: func(topic string) string {
				return SERVICE_APP_NAME + "_" + topic
			},
			ConnectTimeout:    time.Second * 60,
			InitializeTimeout: time.Second * 60,
			ProjectID:         PROJECT_ID,
		},
		logger,
	)
	if err != nil {
		log.Fatal("failed to create gcp subscriber", err)
	}

	for topic, handler := range call_topics {
		// Subscribe will create the subscription. Only messages that are sent after the subscription is created may be received.
		messages, err := subscriber.Subscribe(context.Background(), topic)
		if err != nil {
			log.Fatal("failed to subscribe to topic ", topic, "error: ", err)
		}
		subs = append(subs, messages)
		go process(topic, messages, handler)
	}
}

func process(topic string, messages <-chan *message.Message, callback func([]byte, string, string)) {
	for msg := range messages {
		log.Info("subscriber received message: %s, topic: %s, payload: %s", msg.UUID, topic, string(msg.Payload))

		var qm QueuedMessage.QueuedMessage
		err := qm.UnmarshalJSON(msg.Payload)
		if err != nil {
			fmt.Printf("Failed to unmarshal QueuedMessage message. Error: %s", err)
		}

		log.Debug("received QueuedMessage object: ", qm)

		result, err := json.Marshal(qm.Payload)
		if err != nil {
			fmt.Printf("Could not re-marshal OCPP payload: %s\n", err)
		}

		callback(result, qm.MessageId, qm.DeviceId)

		msg.Ack()
	}
}
