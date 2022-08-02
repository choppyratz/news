FROM golang:1.17-alpine as build-stage

RUN mkdir -p /app

WORKDIR /app

RUN go mod download

COPY . /app

RUN go build -o news main.go


EXPOSE 9993

ENTRYPOINT [ "news" ]