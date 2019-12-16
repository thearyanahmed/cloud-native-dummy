FROM golang:1.8-alpine

ENV SOURCES /go/src/github.com/thearyanahmed/cloud-native-dummy

COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go install

ENV PORT 8080

EXPOSE 8080

ENTRYPOINT microservice-with-go

