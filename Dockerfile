FROM golang:1.13-alpine as dev

RUN apk add --no-cache make git curl build-base

COPY go.mod go.sum main.go models.go overfit-1334-pubsub.json /app/

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /go/bin/app/cloudbuild cloudbuild

ENTRYPOINT ["/go/bin/app/cloudbuild"]
