# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY ./myPackage/*.go ./myPackage/
COPY ./myWebPackage/*.go ./myWebPackage/

RUN go build -o /docker-mypoc

EXPOSE 8080

CMD [ "/docker-mypoc" ]