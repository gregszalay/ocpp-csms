FROM golang:1.14-alpine as builder

RUN apk add --no-cache git

WORKDIR /app

COPY . .
COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN go build main.go

FROM alpine as runtime

SHELL ["/bin/bash", "-c"]

# Copy executable binary file from the 'builder' image to this 'runtime' image
COPY --from=builder /app/main /app/

# Run the binary executable
CMD ["/app/main"]

