package websocketserver

import (
	"github.com/gorilla/websocket"
	"github.com/gregszalay/ocpp-csms-common-types/QueuedMessage"
	"github.com/gregszalay/ocpp-csms/websocket-service/messagemux"
	"github.com/gregszalay/ocpp-messages-go/wrappers"
	log "github.com/sirupsen/logrus"
)

type ChargingStationConnection struct {
	stationId string
	wsConn    *websocket.Conn
}

var AllMessagesToDeviceMap map[string]chan *QueuedMessage.QueuedMessage = map[string]chan *QueuedMessage.QueuedMessage{}

func (conn *ChargingStationConnection) processIncomingMessages() {
	for {
		_, receivedMessage, err := conn.wsConn.ReadMessage()
		if err != nil {
			log.Error(err)
			return
		}
		log.Debug("received new message from charging station with id ", conn.stationId, " message: ", string(receivedMessage))
		log.Debug("Putting message into channel for processing")
		messagemux.ProcessAndPublish(conn.stationId, receivedMessage)
	}
}

func (conn *ChargingStationConnection) writeMessagesToDevice() {
	stationChannel := AllMessagesToDeviceMap[conn.stationId]
	for qm := range stationChannel {
		log.Debug("Writing message to charging station ", conn.stationId, " via ws connection. message: ", qm)
		callresult := wrappers.CALLRESULT{
			MessageTypeId: wrappers.CALLRESULT_TYPE,
			MessageId:     qm.MessageId,
			Payload:       qm.Payload,
		}
		if err := conn.wsConn.WriteMessage(1, callresult.Marshal()); err != nil {
			log.Error("failed to write message to charging station ", conn.stationId)
			log.Error(err)
			return
		}
	}
}
