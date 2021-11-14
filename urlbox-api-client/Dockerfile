FROM golang:alpine3.14

WORKDIR /app

COPY . .

RUN go mod edit -replace ekstrah.com/go-protoBox-grpc=/app/gen/protobuf/ekstrah.com/go-protoBox-grpc

RUN go mod download

EXPOSE 8080

RUN go build main.go

ENTRYPOINT ["./main"]