package main

import (
	"sync"

	"github.com/gregszalay/ocpp-csms/websocket-service/subscribing"
	"github.com/gregszalay/ocpp-csms/websocket-service/websocketserver"
	log "github.com/sirupsen/logrus"
)

func main() {
	//amqppublisher.Setup()

	log.SetLevel(log.InfoLevel)

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
