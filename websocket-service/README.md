

 # OCPP Websocket Service

 ### Main responsibilities of the Service
   
   - via Websockets
      - accept & validate device connections
   - via Google Pub/sub messages:
      - publish all incoming WS OCPP messages to relevant topics (e.g. HeartbeatRequest) 
      - subscibe to outgoing WS OCPP message topics and send them to the connected devices

 ## To run:

    docker build . -t gregszalay/ocpp-websocket-service:v1.0.0

    docker run -e GOOGLE_APPLICATION_CREDENTIALS=/tmp/keys/PRIVATE.json -v C:/repos/ocpp-websocket-service/credentials/PRIVATE.json:/tmp/keys/PRIVATE.json:ro -p 3000:3000 gregszalay/ocpp-websocket-service:v1.0.0