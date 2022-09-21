

 # Websocket Service
 ### Main responsibilities of the Service
   
   - via Websockets
      - accept & validate device connections
   - via Google Pub/sub messages:
      - publish all incoming WS OCPP messages to relevant topics (e.g. HeartbeatRequest) 
      - subscibe to outgoing WS OCPP message topics and send them to the connected devices