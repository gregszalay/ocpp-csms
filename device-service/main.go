package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"

	sw "github.com/gregszalay/ocpp-csms/device-service/http/go"
	"github.com/gregszalay/ocpp-csms/device-service/subscribing"
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
		fmt.Println("Creating http server...")
		router := sw.NewRouter()
		log.Fatal(http.ListenAndServe(":5000", router))
		waitgroup.Done()
	}()

	waitgroup.Add(1)
	go func() {
		fmt.Println("Creating pubsub subscriptions...")
		subscribing.Subscribe()
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
