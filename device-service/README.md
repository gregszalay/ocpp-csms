
 # Device Management Service
 ### Main responsibilities of the Service
   
   - via REST API (see REST_API/ folder)
      - Authenticate charging stations (verify that the charging station is in the db)
   - via Google Pub/sub messages:
      - subscribe to device-related topics (e.g. BootNotificationRequest) and publish responses (e.g. BootNotificationResponse) 
