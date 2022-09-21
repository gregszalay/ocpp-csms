package main

import (
	"ocpp-websocket-service/websocketserver"
)

func main() {
	//amqppublisher.Setup()
	websocketserver.Run()

}
