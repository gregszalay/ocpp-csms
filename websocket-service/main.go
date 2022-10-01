package main

import (
	"fmt"
	"sync"

	"github.com/gregszalay/ocpp-csms/websocket-service/pubsub"
	"github.com/gregszalay/ocpp-csms/websocket-service/websocketserver"
)

func main() {
	//amqppublisher.Setup()

	var waitgroup sync.WaitGroup

	waitgroup.Add(1)
	go func() {
		fmt.Println("Creating pubsub subscriptions...")
		pubsub.Subscribe()
		waitgroup.Done()
	}()

	waitgroup.Add(1)
	go func() {
		fmt.Println("Starting Websocket server...")
		websocketserver.Start()
		waitgroup.Done()
	}()

	waitgroup.Wait()

}
