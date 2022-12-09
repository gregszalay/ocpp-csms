package websocketserver

import (
	"net/http"

	"github.com/gregszalay/ocpp-csms-common-types/QueuedMessage"
	"github.com/gregszalay/ocpp-csms/websocket-service/authentication"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:    1024,
	WriteBufferSize:   1024,
	EnableCompression: true,
}

var openConnections map[string]*ChargingStationConnection = map[string]*ChargingStationConnection{}

func ChargingStationHandler(w http.ResponseWriter, r *http.Request) {

	// parse charging station identity
	id, ok := mux.Vars(r)["id"]
	if !ok {
		log.Error("cannot find id in connection url. refusing connection")
		return
	}
	log.Info("websocket connection initiated by charging station with id ", id)

	// authenticate connection
	err := authentication.AuthenticateChargingStation(id, r)
	if err != nil {
		log.Error("failed to authenticate charging station with id ", id, ". refusing connection.")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Failed to authenticate charging station."))
		return
	}
	log.Info("charger successfully authenticated")

	// upgrade to websocket connection
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("failed to establish websocket connection on the server, error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to establish websocket connection on the server"))
		return
	}
	_ = ws.SetCompressionLevel(9)
	log.Info("successfully established websocket connection")

	// Create and save ws connection object
	new_connection := ChargingStationConnection{
		stationId: id,
		wsConn:    ws,
	}
	openConnections[id] = &new_connection
	log.Info("number of open connections: ", len(openConnections))

	// Create channel for messages bound for the charging station
	AllMessagesToDeviceMap[id] = make(chan *QueuedMessage.QueuedMessage)

	// Start message handlers for new connection
	go new_connection.writeMessagesToDevice()
	go new_connection.processIncomingMessages()

	ws.SetCloseHandler(func(code int, text string) error {
		delete(openConnections, id)
		log.Info("connection closed to charging station ", id)
		log.Info("number of remaining open connections: ", len(openConnections))
		return nil
	})
}
