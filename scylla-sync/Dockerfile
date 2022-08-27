FROM golang:1.18.2-alpine as base

ENV APP_HOME /go/src/app

WORKDIR ${APP_HOME}
COPY . .

RUN go mod download
RUN go install github.com/cosmtrek/air@latest
RUN go mod verify
RUN go build -o scylla-sync

FROM base as dev
CMD ["air"]

FROM alpine as prod
ENV APP_HOME /go/src/app
RUN mkdir -p "$APP_HOME"
WORKDIR ${APP_HOME}

COPY --from=base ${APP_HOME}/scylla-sync ${APP_HOME}
COPY --from=base ${APP_HOME}/config ${APP_HOME}/config

EXPOSE 8081
CMD ["./scylla-sync"]
