package websocketserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gregszalay/ocpp-csms/websocket-service/MessageIn"
	"github.com/gregszalay/ocpp-csms/websocket-service/chargerauth"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var connections = make(chan *ChargingStationConnection)
var incomingWSMessages = make(chan IncomingMessage, 3)

type ChargingStationConnection struct {
	identity string
	WSConn   *websocket.Conn
}

var outgoingMessages map[string]string = map[string]string{}

func processInbound(incoming MessageIn.MessageIn) {
	//for incoming := range incomingWSMessages {
	log.Printf("received message in goroutine, payload: %s", string(incoming.Message))
	err := incoming.Process()
	if err != nil {
		fmt.Println("processing error ---")
		fmt.Println(err)
		return
		//panic(err)
	}
}

func Ocpp(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["chargerId"]
	if !ok {
		fmt.Println("id is missing in parameters")
		return
	}
	fmt.Println(`id := `, id)

	// /***** AUTH */
	charger, err := chargerauth.GetCharger(id)
	if err != nil {
		fmt.Println("Error: could not get charger data!")
		return
	}
	fmt.Printf("Charger successfully authenticated: %+v\n", charger)

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
		messageType, receivedMessage, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("-------------------------- ")
		//fmt.Printf("Got message from connection %d\n%+v\n", numberOfLatestConnection, charger)
		fmt.Println("messageType: ", messageType)
		fmt.Println("message payload: ", string(receivedMessage))
		fmt.Println("-------------------------- ")
		new_msg := MessageIn.MessageIn{
			chargerId: id,
			Message:   receivedMessage,
		}
		//incomingWSMessages <- new_msg
		processInbound(new_msg)
		response := receivedMessage
		//ocpp_handlers.ProcessIncomingMessage(string(p))
		if err := ws.WriteMessage(1, response); err != nil {
			log.Println(err)
			return
		}

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

	fmt.Println("Hello OCPP WS 3!")
	router := NewRouter()
	//go processInbound()
	//setupRoutes()
	log.Fatal(http.ListenAndServe(":3000", router))

}
