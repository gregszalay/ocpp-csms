package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gregszalay/ocpp-csms/transaction-service/subscribing"
	log "github.com/sirupsen/logrus"
)

func main() {

	if LOG_LEVEL := os.Getenv("LOG_LEVEL"); LOG_LEVEL != "" {
		setLogLevel(LOG_LEVEL)
	} else {
		setLogLevel("Info")
	}

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
