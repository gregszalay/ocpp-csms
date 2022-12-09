package websocketserver

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var OCPP_CONNECTION_URL string = "/ocpp/{id}"
var OCPP_PORT string = ":3000"

func Start() {
	router := newRouter()
	TLS_CERT_FILE := os.Getenv("TLS_CERT_FILE")
	TLS_KEY_FILE := os.Getenv("TLS_KEY_FILE")
	fmt.Println("TLS_CERT_FILE path:")
	fmt.Println(TLS_CERT_FILE)

	fmt.Println("TLS_KEY_FILE path:")
	fmt.Println(TLS_KEY_FILE)
	
	if TLS_CERT_FILE != "" && TLS_KEY_FILE != "" {
		log.Info("TLS certificate found, serving secure TLS")
		log.Fatal(http.ListenAndServeTLS(OCPP_PORT, TLS_CERT_FILE, TLS_KEY_FILE, router))

	} else {
		log.Warn("TLS certificate not found, serving unsecure ws")
		log.Fatal(http.ListenAndServe(OCPP_PORT, router))
	}
}

func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler = route.HandlerFunc
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"ocpp",
		strings.ToUpper("Get"),
		OCPP_CONNECTION_URL,
		ChargingStationHandler,
	},
}
