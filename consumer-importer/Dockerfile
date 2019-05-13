FROM golang:latest

WORKDIR /go/src/consumer-importer
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["main"]