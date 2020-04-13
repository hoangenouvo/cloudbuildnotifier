FROM golang:1.13-alpine
RUN apk add --no-cache make git curl build-base
COPY . /app/
WORKDIR /app
RUN mkdir build && cp .env credential.json build/ && CGO_ENABLED=0 GOOS=linux go build -a -o build/cloudbuild github.com/lxhoang97/cloudbuildnotifier

FROM alpine:latest as app
COPY --from=0 app/build .
CMD ["./cloudbuild"]
