package main

import (
	"github.com/gregszalay/ocpp-csms/websocket-service/websocketserver"
)

func main() {
	//amqppublisher.Setup()
	websocketserver.Run()

}
