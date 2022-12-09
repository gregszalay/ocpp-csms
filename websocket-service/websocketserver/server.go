package websocketserver

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
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
	SERVER_TLS_CERT_FILE := os.Getenv("SERVER_TLS_CERT_FILE")
	SERVER_TLS_KEY_FILE := os.Getenv("SERVER_TLS_KEY_FILE")
	log.Info("SERVER_TLS_CERT_FILE path:")
	log.Info(SERVER_TLS_CERT_FILE)
	log.Info("SERVER_TLS_KEY_FILE path:")
	log.Info(SERVER_TLS_KEY_FILE)

	CLIENT_TLS_CERT_FILE := os.Getenv("CLIENT_TLS_CERT_FILE")
	log.Info("CLIENT_TLS_CERT_FILE path:")
	log.Info(CLIENT_TLS_CERT_FILE)

	if SERVER_TLS_CERT_FILE != "" && SERVER_TLS_KEY_FILE != "" {

		caCert, _ := ioutil.ReadFile(CLIENT_TLS_CERT_FILE)
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		tlsConfig := &tls.Config{
			ClientCAs:  caCertPool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		}
		tlsConfig.BuildNameToCertificate()

		server := &http.Server{
			Handler:   router,
			Addr:      OCPP_PORT,
			TLSConfig: tlsConfig,
		}

		log.Info("TLS certificate found, serving secure TLS")
		log.Fatal(server.ListenAndServeTLS(SERVER_TLS_CERT_FILE, SERVER_TLS_KEY_FILE))
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
