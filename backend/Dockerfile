FROM golang AS builder
WORKDIR /loop
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/loop .

FROM alpine:latest
RUN addgroup -S loop && adduser -S loop -G loop
USER loop
WORKDIR /home/loop
COPY --from=builder /bin/loop ./
EXPOSE 3000
ENTRYPOINT ["./loop"]