FROM golang:1.19.0-alpine as base

ENV APP_HOME /go/src/party-service

WORKDIR ${APP_HOME}
COPY . .

RUN go mod download
RUN go install github.com/cosmtrek/air@latest
RUN go mod verify
RUN go build -o party-service 

FROM base as dev
CMD ["air"]

FROM alpine as prod
ENV APP_HOME /go/src/party-service
RUN mkdir -p "$APP_HOME"
WORKDIR ${APP_HOME}

COPY --from=base ${APP_HOME}/party-service ${APP_HOME}
COPY --from=base ${APP_HOME}/config ${APP_HOME}/config

EXPOSE 8081
CMD ["./party-service"]
