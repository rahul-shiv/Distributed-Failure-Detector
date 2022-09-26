FROM golang:1.16-alpine

WORKDIR /app

RUN apk update && apk add make protobuf-dev=3.18.1-r1
RUN GO111MODULE=on \
        go get google.golang.org/protobuf/cmd/protoc-gen-go@v1.26.0 \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0 \ 
        github.com/golang/protobuf/proto \
        github.com/golang/protobuf/ptypes

COPY go.mod ./
COPY go.sum ./
RUN go mod download
ADD protos ./protos
COPY coordinator/*.go ./

RUN protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative protos/coordinator.proto
RUN protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative protos/distributed_grep.proto

RUN go build -o /client

EXPOSE 50051

CMD [ "/client" ]