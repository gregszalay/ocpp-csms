package main

import (
	"os"
	"sync"

	"github.com/gregszalay/ocpp-csms/websocket-service/subscribing"
	"github.com/gregszalay/ocpp-csms/websocket-service/websocketserver"
	log "github.com/sirupsen/logrus"
)

func main() {

	if LOG_LEVEL := os.Getenv("LOG_LEVEL"); LOG_LEVEL != "" {
		setLogLevel(LOG_LEVEL)
	} else {
		setLogLevel("Info")
	}

	var waitgroup sync.WaitGroup

	waitgroup.Add(1)
	go func() {
		log.Info("Creating pubsub subscriptions...")
		subscribing.Subscribe()
		waitgroup.Done()
	}()

	waitgroup.Add(1)
	go func() {
		log.Info("Starting Websocket server...")
		websocketserver.Start()
		waitgroup.Done()
	}()

	waitgroup.Wait()

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
