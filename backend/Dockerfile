FROM golang:1.18

WORKDIR /go/src/loop

COPY . .

RUN go mod download

RUN go run main.go