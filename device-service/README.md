
 # Device Management Service
 ### Main responsibilities of the Service
   
   - via REST API (see REST_API/ folder)
      - Authorize connections (upon request from the Websocket Service)
   - via Google Pub/sub messages:
      - subscribe to device-related topics (e.g. BootNotificationRequest) and publish responses (e.g. BootNotificationResponse) 
