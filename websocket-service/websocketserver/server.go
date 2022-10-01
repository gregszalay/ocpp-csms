package websocketserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gregszalay/ocpp-csms-common-types/QueuedMessage"
	"github.com/gregszalay/ocpp-csms/websocket-service/authentication"
	"github.com/gregszalay/ocpp-csms/websocket-service/messageprocessing"
	"github.com/gregszalay/ocpp-csms/websocket-service/pubsub"
	"github.com/gregszalay/ocpp-messages-go/wrappers"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var connections map[string]*websocket.Conn = map[string]*websocket.Conn{}

var outgoingMessages map[string]string = map[string]string{}

func Ocpp(w http.ResponseWriter, r *http.Request) {
	// parse charging station identity
	id, ok := mux.Vars(r)["chargerId"]
	if !ok {
		fmt.Println("id is missing in parameters")
		return
	}
	fmt.Printf("Websocket connection initiated by charging station %s \n", id)

	// authenticate connection
	err := authentication.AuthenticateChargingStation(id, r)
	if err != nil {
		fmt.Println("error: failed to authenticate charging station")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("failed to authenticate charging station!"))
		return
	}
	fmt.Printf("charger successfully authenticated\n")

	// upgrade to websocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to authenticate charging station!"))
		return
	}
	fmt.Println("successfully established websocket connection")

	// Save ws conn
	connections[id] = ws
	fmt.Printf("number of connections: %d\n", len(connections))

	// Create response queue for charging station
	pubsub.ToChargerQueue[id] = make(chan *QueuedMessage.QueuedMessage)

	// Start message handlers
	go writeOutgoingMessages(id, ws)
	go handleIncomingMessages(id, ws)
}

func handleIncomingMessages(id string, ws *websocket.Conn) {
	for {
		_, receivedMessage, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		//fmt.Printf("Got message from connection %d\n%+v\n", numberOfLatestConnection, charger)
		fmt.Println("websocket service received a message: \n", string(receivedMessage))
		messageprocessing.ProcessableMessages <- &messageprocessing.ProcessableMessage{
			ChargerId: id,
			Message:   receivedMessage,
		}
	}
}

func writeOutgoingMessages(device_id string, ws *websocket.Conn) {
	for qm := range pubsub.ToChargerQueue[device_id] {
		fmt.Println("Got message from ToChargerQueue")

		callresult := wrappers.CALLRESULT{
			MessageTypeId: wrappers.CALLRESULT_TYPE,
			MessageId:     qm.MessageId,
			Payload:       qm.Payload,
		}
		if err := ws.WriteMessage(1, callresult.Marshal()); err != nil {
			log.Println(err)
			return
		}
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OCPP websocket service homepage.")
}

func Start() {
	fmt.Println("Starting Websocket server...")
	router := NewRouter()
	go messageprocessing.Start()
	log.Fatal(http.ListenAndServe(":3000", router))

}
