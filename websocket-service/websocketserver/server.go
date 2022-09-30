package websocketserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gregszalay/ocpp-csms-common-types/QueuedMessage"
	"github.com/gregszalay/ocpp-csms/websocket-service/MessageIn"
	"github.com/gregszalay/ocpp-csms/websocket-service/ocppsub"
	"github.com/gregszalay/ocpp-messages-go/wrappers"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var connections = make(chan *ChargingStationConnection)
var incomingWSMessages = make(chan MessageIn.MessageIn, 3)

type ChargingStationConnection struct {
	identity string
	WSConn   *websocket.Conn
}

var outgoingMessages map[string]string = map[string]string{}

func processInbound() {
	for incoming_message := range incomingWSMessages {
		log.Printf("Sending message to processor: \n%s", string(incoming_message.Message))
		err := incoming_message.Process()
		if err != nil {
			fmt.Println("processing error ---")
			fmt.Println(err)
			return
		}
	}
}

func Ocpp(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["chargerId"]
	if !ok {
		fmt.Println("id is missing in parameters")
		return
	}
	fmt.Printf("Websocket connection initiated by charging station with id %s \n", id)

	// /***** AUTH */
	// charger, err := chargerauth.GetCharger(id)
	// if err != nil {
	// 	fmt.Println("Error: could not get charger data!")
	// 	return
	// }
	// fmt.Printf("Charger successfully authenticated: %+v\n", charger)

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	//fmt.Printf("Charging Station WS connected: %+v\n", newCharger)
	// Store charging station connection
	// connections <- &ChargingStationConnection{
	// 	identity: id,
	// 	WSConn:   ws,
	// }

	//numberOfLatestConnection := len(connections) - 1

	for {
		_, receivedMessage, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		//fmt.Printf("Got message from connection %d\n%+v\n", numberOfLatestConnection, charger)
		fmt.Println("websocket service received a message: \n", string(receivedMessage))
		new_msg := MessageIn.MessageIn{
			ChargerId: id,
			Message:   receivedMessage,
		}
		incomingWSMessages <- new_msg

		for topic, queue := range ocppsub.Subs {
			for msg := range queue {
				log.Printf("sending message: %s, topic: %s, payload: %s", msg.UUID, topic, string(msg.Payload))
				var qm QueuedMessage.QueuedMessage
				err := qm.UnmarshalJSON(msg.Payload)
				if err != nil {
					fmt.Printf("Failed to unmarshal OCPP CALLRESULT message. Error: %s", err)
				}
				if qm.DeviceId == id {
					fmt.Println("qm.DeviceId == id, hurray")
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

		}
		//processInbound(new_msg)
		//ocpp_handlers.ProcessIncomingMessage(string(p))

	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OCPP websocket service homepage.")
}

func sendToClient(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Trying to send request to device.")

	// for index, element := range connections {
	// 	fmt.Printf("%d\n%+v\n", index, element)
	// 	if r.URL.Query().Get("serialNumber") == element.ChargingStation.SerialNumber &&
	// 		r.URL.Query().Get("model") == element.ChargingStation.Model &&
	// 		r.URL.Query().Get("vendorName") == element.ChargingStation.VendorName {
	// 		message := []byte("Hello " + element.ChargingStation.VendorName + element.ChargingStation.SerialNumber)
	// 		if err := element.WSConn.WriteMessage(1, message); err != nil {
	// 			log.Println(err)
	// 			return
	// 		}
	// 	}
	// }

}

func Run() {

	fmt.Println("Starting OCPP CSMS server...")
	router := NewRouter()
	go processInbound()
	//setupRoutes()
	log.Fatal(http.ListenAndServe(":3000", router))

}
