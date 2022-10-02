package websocketserver

import (
	"net/http"
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
	log.Fatal(http.ListenAndServe(OCPP_PORT, router))
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
