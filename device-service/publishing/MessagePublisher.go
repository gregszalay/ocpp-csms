package publishing

import (
	"encoding/json"
	"os"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/pkg/googlecloud"
	"github.com/ThreeDotsLabs/watermill/message"
	log "github.com/sirupsen/logrus"
)

var PROJECT_ID string = os.Getenv("GCP_PROJECT_ID")

var gcp_pub *googlecloud.Publisher = nil

func Publish(topic string, qm interface{}) error {
	logger := watermill.NewStdLogger(true, true)

	if gcp_pub == nil {
		publisher, err := googlecloud.NewPublisher(googlecloud.PublisherConfig{
			ProjectID: PROJECT_ID,
		}, logger)
		if err != nil {
			log.Fatal("Failed to create gcp pulisher client")
		}
		gcp_pub = publisher
	}

	qm_json, err := json.Marshal(qm)
	if err != nil {
		return err
	}

	log.Debug("QueuedMessage to be published to topic ", topic, " : ", string(qm_json))

	msg := message.NewMessage(watermill.NewUUID(), qm_json)
	if err := gcp_pub.Publish(topic, msg); err != nil {
		return err
	}
	return nil

}
