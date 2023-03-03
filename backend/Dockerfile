FROM golang:1.18

WORKDIR /go/src/loop

COPY . .

RUN go mod download

CMD [ "go build -o . ." ]