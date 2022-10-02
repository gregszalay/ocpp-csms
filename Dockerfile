FROM golang:1.18-alpine as builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN mkdir build
RUN go build -o build/device-service ./device-service
RUN go build -o build/websocket-service ./websocket-service
RUN go build -o build/user-service ./user-service


FROM alpine as runtime

RUN apk add --no-cache bash

SHELL ["/bin/bash", "-c"]

# Copy executable binary file from the 'builder' image to this 'runtime' image
COPY --from=builder /app/build/device-service /app/
COPY --from=builder /app/build/websocket-service /app/
COPY --from=builder /app/build/user-service /app/