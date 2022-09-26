// This is the main client for calling the grep command.
package main

import (
	"context"
	"flag"
	"fmt"
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
	start := time.Now()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDistributedGrepClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	searchOptions := args[0]
	searchString := args[1]
	searchFile := args[2]
	r, err := c.CallGrep(ctx, &pb.ClientRequest{QueryString: searchString, Args: searchOptions, QueryFile: searchFile})
	if err != nil {
		log.Printf("could not grep!: %v", err)
	} else {
		duration := time.Since(start)
		for _, item := range r.GetResult() {
			fmt.Printf("%s", item)
		}
		fmt.Printf("\nTotal Matches : %d\n", r.GetTotal())
		fmt.Printf("Query Time : %v\n", duration)
	}
}
