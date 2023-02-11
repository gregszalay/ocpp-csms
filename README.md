# OCPP CSMS (OCPP 2.0.1.)

Charging station Management system for electric vehicle charging stations.

Based on: OCPP-2.0.1 (Full protocol documentation:
[openchargealliance.org/downloads](https://www.openchargealliance.org/downloads/))

**_ Under development _**

## Working principle

E.g. incoming BootNotificationRequest:

==> REQUEST data flow:

charging station -> websocket-service -> GCP Pub/Sub topic "BootNotificationRequest" -> device-service

<== RESPONSE data flow:

charging station <- websocket-service <- GCP Pub/Sub topic "BootNotificationResponse" <- device-service

## Functional features:

- Implemented OCPP messages:

| Implemented OCPP Messages                             | Microservice name   |
| ----------------------------------------------------- | ------------------- |
| BootNotificationRequest, BootNotificationResponse     | device-service      |
| HeartbeatRequest, HeartbeatResponse                   | device-service      |
| StatusNotificationRequest, StatusNotificationResponse | device-service      |
| TransactionEventRequest, TransactionEventResponse     | transaction-service |
| AuthorizeRequest, AuthorizeResponse                   | user-service        |

## Technical features:

- microservice-based architecture
- GO: managing multiple websocket connections easily with goroutines 
- Mutual TLS authentication
- Google Firestore for simple data persistence
- Google Cloud Pub/Sub for async messaging between services (using [watermill.io](https://github.com/ThreeDotsLabs/watermill))
- REST API for managing charging station information (API description: [swaggerhub.com](https://app.swaggerhub.com/apis/gregszalay/ocpp_device_service/v2.0.0))

## Quick Start

1. Create a **Google Cloud Platform project**. Follow the [Google guide](https://cloud.google.com/resource-manager/docs/creating-managing-projects).

2. Create a **service account** and create **JSON credentials** for it. Follow the [Google guide](https://developers.google.com/workspace/guides/create-credentials) (see the _"Service account credentials"_ section). When you create the service account, under "select a role" choose `Pub/Sub Admin` and `Firebase Admin` (note: for production use, you may want to restrict these to lower roles, needs to be tested).

3. Once you have the JSON credentials, place them in the following directories, under the name 'PRIVATE.json':

   - device-service/credentials/PRIVATE.json
   - websocket-service/credentials/PRIVATE.json
   - user-service/credentials/PRIVATE.json
   - transaction-service/credentials/PRIVATE.json

> Note: you can use the same service account and the same JSON credentials file in all places, or you can create separate ones for extra security if you want

4. Create a Cloud Firestore database Follow the [Google guide](https://firebase.google.com/docs/firestore/quickstart)

5. In the **docker-compose.yml** file, replace `chargerevolutioncloud` name with the **project id** of your own GCP project.

6. Install docker and docker-compose if you haven't already

> If you are deploying on a remote machine or VPS, **open** up port `3000` and port `5000` so you can access the app remotely from your own machine.
> Important: For actual deployment you should add a reverse proxy layer or some other security measure to protect your open port. A good place to start: [How To Deploy a Go Web Application with Docker and Nginx on Ubuntu 18.04](https://www.digitalocean.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker-and-nginx-on-ubuntu-18-04)

7.  Build and run the app on your machine (or VPS)

        docker compose up --build

8.  The device service should now be running on `localhost:5000`. Use Postman or some other tool to send a post request and create a charging station (API description: [swaggerhub.com](https://app.swaggerhub.com/apis/gregszalay/ocpp_device_service/v2.0.0)) For example, you can send something like this in the request payload:

    POST: `localhost:5000/chargingstations/create`

        {
            "id": "CS123",
            "serialNumber": "5",
            "model": "CS-5500",
            "vendorName": "ChargerMaker Inc.",
            "firmwareVersion": "1.8",
            "modem": {
                "iccid": "24",
                "imsi": "24"
            },
            "location": {
                "lat": 41.366446,
                "lng": -38.1854651
            }
        }

    If this is successful, a collection named _"chargingstations"_ and a document named _"CS123"_ should have been created in the Firestore database of your project.

9.  The websocket service should be running on `localhost:3000`. Connect to port 3000 as a websocket client (use the station id you have created in the previous step):

    `ws://{HOST}:3000/ocpp/{stationid}`

    > Tip: Postman now supports websocket connections (beta version), but only on the Windows desktop app I believe. This could be useful for testing.

### To-do

- full implementation of all basic messages (reset, getconfig etc.)
