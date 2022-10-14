# ocpp-csms

Charging Station Management System for electric vehicles - OCPP 2.0.1. compliant

## Functionality working so far:

- BootNotificationRequest and response
- Transaction and response

## Quick start

    Generate & save GCP credential JSON file with Pub/Sub and Firestore access into the following locations/filenames:
        - device-service/credentials/PRIVATE.json
        - websocket-service/credentials/PRIVATE.json
        - user-service/credentials/PRIVATE.json
        - transaction-service/credentials/PRIVATE.json

    docker compose up --build

    Create charging station records on port 5000 (REST API description in the `device-service/http/api` folder). Make note of the unique station id you created

    Connect to port 3000 as a websocket client:
        ws://{HOST}:3000/ocpp/{stationid}
