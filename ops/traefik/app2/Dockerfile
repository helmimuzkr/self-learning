FROM golang:1.23.0-alpine

# when creating subdirectories hanging off from a non-existing parent directory(s) you must pass the -p flag to mkdir
RUN mkdir -p /app2

WORKDIR /app2

ADD . /app2

RUN go test -v ./...

RUN go build -o main .

CMD ["/app2/main"]
