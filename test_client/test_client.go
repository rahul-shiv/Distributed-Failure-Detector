// This is the test client used to create and delete test logs.
package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	pb "mp1/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	local_addr = flag.String("local_addr", "localhost:50051", "the address to connect to")
	vm_addr    = flag.String("vm_addr", "fa22-cs425-0301:50051", "the address to connect to")
)

var addr = vm_addr

func main() {
	flag.Parse()
	args := os.Args[1:]
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDistributedGrepClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	cmd := args[0]
	if cmd == "create" {
		_, err = c.CallCreateLogs(ctx, &pb.CreateLogsRequest{})
	} else if cmd == "delete" {
		_, err = c.CallDeleteLogs(ctx, &pb.DeleteLogsRequest{})
	}
	if err != nil {
		log.Printf("Failed Test Client Call: %v", err)
	}
}
