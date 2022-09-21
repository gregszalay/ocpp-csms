
 # OCPP Device Management Service

 ### Main responsibilities of the Service
   
   - via REST API (see REST_API/ folder)
      - Authorize connections (upon request from the Websocket Service)
   - via Google Pub/sub messages:
      - subscribe to device-related topics (e.g. BootNotificationRequest) and publish responses (e.g. BootNotificationResponse) 

 ## To run:

    docker build . -t gregszalay/ocpp-device-service:v1.0.0

    docker run -e GOOGLE_APPLICATION_CREDENTIALS=/tmp/keys/PRIVATE.json -v C:/repos/ocpp-device-service/credentials/PRIVATE.json:/tmp/keys/PRIVATE.json:ro -p 5000:5000 gregszalay/ocpp-device-service:v1.0.0 