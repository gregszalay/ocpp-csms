package websocketserver

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func Index(w http.ResponseWriter, r *http.Request) {
	device_endpoint := "ws://" + "<SERVER_URL>" + "/ocpp/CHARGER_ID"
	fmt.Fprintf(w, "Please use the "+device_endpoint+" URL to connect your device.")
	log.Info("Incoming request to index page")
}
