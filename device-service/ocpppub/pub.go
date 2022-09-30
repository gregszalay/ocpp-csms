package ocpppub

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/pkg/googlecloud"
	"github.com/ThreeDotsLabs/watermill/message"
)

var gcp_pub *googlecloud.Publisher = nil

func Publish(topic string, qm interface{}) error {
	logger := watermill.NewStdLogger(true, true)

	if gcp_pub == nil {
		publisher, err := googlecloud.NewPublisher(googlecloud.PublisherConfig{
			ProjectID: "chargerevolutioncloud",
		}, logger)
		if err != nil {
			panic(err)
		}
		gcp_pub = publisher
	}

	qm_json, err := json.Marshal(qm)
	if err != nil {
		return err
	}

	//fmt.Printf("Marshalled QM: %+v\n", string(qm_json))

	msg := message.NewMessage(watermill.NewUUID(), qm_json)
	if err := gcp_pub.Publish(topic, msg); err != nil {
		return err
	}
	return nil

}
