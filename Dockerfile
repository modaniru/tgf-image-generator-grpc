FROM golang:alpine as builder

ENV GOPATH=/

WORKDIR /app

COPY pkg pkg
COPY src src
COPY go.mod .
COPY go.sum .

RUN go get ./...
RUN go build src/main.go

FROM alpine:latest

COPY --from=builder ./app/main .
COPY images images
COPY fonts fonts

CMD ["./main"]

