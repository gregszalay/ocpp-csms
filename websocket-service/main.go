package main

import (
	"fmt"
	"sync"

	"github.com/gregszalay/ocpp-csms/websocket-service/ocppsub"
	"github.com/gregszalay/ocpp-csms/websocket-service/websocketserver"
)

func main() {
	//amqppublisher.Setup()

	var waitgroup sync.WaitGroup

	waitgroup.Add(1)
	go func() {
		fmt.Println("Creating Websocket server...")
		websocketserver.Run()
		waitgroup.Done()
	}()

	waitgroup.Add(1)
	go func() {
		fmt.Println("Creating pubsub subscriptions...")
		ocppsub.Subscribe()
		waitgroup.Done()
	}()

	waitgroup.Wait()

}
