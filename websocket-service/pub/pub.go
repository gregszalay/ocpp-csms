package pub

import (
	"encoding/json"
	"fmt"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/pkg/googlecloud"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gregszalay/ocpp-csms-common-types/QueuedMessage"
)

var pub *googlecloud.Publisher = nil

func Publish(qm QueuedMessage.QueuedMessage, topic string) {
	logger := watermill.NewStdLogger(true, true)

	if pub == nil {
		publisher, err := googlecloud.NewPublisher(googlecloud.PublisherConfig{
			ProjectID: "chargerevolutioncloud",
		}, logger)
		if err != nil {
			panic(err)
		}
		pub = publisher
	}

	fmt.Println("topic:")
	fmt.Println(topic)

	qm_json, err := json.Marshal(qm)
	if err != nil {
		fmt.Printf("Error: Could not marshal queue message: %s\n", err)
	}

	fmt.Printf("Marshalled QM: %+v\n", string(qm_json))

	msg := message.NewMessage(watermill.NewUUID(), qm_json)
	if err := pub.Publish(topic, msg); err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
		panic(err)
	}

}
