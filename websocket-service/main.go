package main

import (
	"github.com/gregszalay/ocpp-csms/websocketserver"
)

func main() {
	//amqppublisher.Setup()
	websocketserver.Run()

}
