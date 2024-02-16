FROM golang:1.22

WORKDIR /app
ADD . .

RUN go mod download

RUN go build cmd/main.go
CMD ["main"]