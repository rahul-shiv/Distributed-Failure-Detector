sudo mkdir /mp1
chmod +777 /mp1
git clone https://gitlab.engr.illinois.edu/rahul11/cs425-mp1.git
cd cs425-mp1
sudo yum install protobuf-compiler
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
./generate_protobufs.sh
go run worker/worker_server.go
go run coordinator/coordinator.go
go run client/client.go