package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gregszalay/ocpp-csms/user-service/db"
	"github.com/gregszalay/ocpp-csms/user-service/subscribing"
	"github.com/gregszalay/ocpp-messages-go/types/AuthorizeRequest"
	log "github.com/sirupsen/logrus"
)

func main() {

	if LOG_LEVEL := os.Getenv("LOG_LEVEL"); LOG_LEVEL != "" {
		setLogLevel(LOG_LEVEL)
	} else {
		setLogLevel("Info")
	}

	// FOR TESTING PURPOSES ONLY:
	newIdTokenInfo := AuthorizeRequest.IdTokenType{
		IdToken: "abcd",
		Type:    AuthorizeRequest.IdTokenEnumType_1_ISO14443,
	}
	db.CreateIdToken("ID001", newIdTokenInfo)

	fmt.Println("Creating pubsub subscriptions...")
	subscribing.Subscribe()
	for {
		time.Sleep(time.Millisecond * 10)
	}

}

func setLogLevel(levelName string) {
	switch levelName {
	case "Panic":
		log.SetLevel(log.PanicLevel)
	case "Fatal":
		log.SetLevel(log.FatalLevel)
	case "Error":
		log.SetLevel(log.ErrorLevel)
	case "Warn":
		log.SetLevel(log.WarnLevel)
	case "Info":
		log.SetLevel(log.InfoLevel)
	case "Debug":
		log.SetLevel(log.DebugLevel)
	case "Trace":
		log.SetLevel(log.TraceLevel)
	}
}
