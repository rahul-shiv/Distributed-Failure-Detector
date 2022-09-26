// Package main implements a client for coordinator service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	re "regexp"
	"strconv"
	"time"

	pb "mp1/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var vm_machine_ips = [10]string{"fa22-cs425-0301:50052", "fa22-cs425-0302:50052", "fa22-cs425-0303:50052", "fa22-cs425-0304:50052", "fa22-cs425-0305:50052",
	"fa22-cs425-0306:50052", "fa22-cs425-0307:50052", "fa22-cs425-0308:50052", "fa22-cs425-0309:50052", "fa22-cs425-0310:50052"}

var local_machine_ips = [1]string{"localhost:50052"}

var machine_ips = &vm_machine_ips

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedDistributedGrepServer
}

/* this rpc triggered by the client, calls the create logs function parallely
 */
func (s *server) CallCreateLogs(ctx context.Context, in *pb.CreateLogsRequest) (*pb.CreateLogsResponse, error) {
	log.Printf("Received Create Logs Request")

	returnChan := make(chan error)
	defer close(returnChan)
	for index, ip := range machine_ips {
		go CreateLogs(returnChan, ip, index+1)
	}

	for k := 0; k < len(machine_ips); k++ {
		<-returnChan
	}
	return &pb.CreateLogsResponse{}, nil
}

/* this function triggers the create logs rpc on the workers
 */
func CreateLogs(returnChan chan error, ip string, vmInd int) {
	fmt.Printf("Connecting to %s\n", ip)
	conn, err := grpc.Dial(ip, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGrepInterfaceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	_, err = c.CallWorkerCreateLogs(ctx, &pb.WorkerCreateLogsRequest{
		VmIndex: int32(vmInd)})
	if err != nil {
		log.Printf("could not grep!: %v", err)
	}
	returnChan <- err
}

/* this rpc triggered by the client, calls the delete logs function parallely
 */
func (s *server) CallDeleteLogs(ctx context.Context, in *pb.DeleteLogsRequest) (*pb.DeleteLogsResponse, error) {
	log.Printf("Received Delete Logs Request")

	returnChan := make(chan error)
	defer close(returnChan)
	for _, ip := range machine_ips {
		go DeleteLogs(returnChan, ip)
	}
	for k := 0; k < len(machine_ips); k++ {
		<-returnChan
	}
	return &pb.DeleteLogsResponse{}, nil
}

/* this function triggers the delete logs rpc on the workers
 */
func DeleteLogs(returnChan chan error, ip string) {
	fmt.Printf("Connecting to %s\n", ip)
	conn, err := grpc.Dial(ip, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGrepInterfaceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	_, err = c.CallWorkerDeleteLogs(ctx, &pb.WorkerDeleteLogsRequest{})
	if err != nil {
		log.Printf("could not grep!: %v", err)
	}
	returnChan <- err
}

/* this rpc triggered by the client, calls the connect function parallely
 */
func (s *server) CallGrep(ctx context.Context, in *pb.ClientRequest) (*pb.CoordinatorResponse, error) {
	log.Printf("Received: %v %v %v", in.GetArgs(), in.GetQueryString(), in.GetQueryFile())

	searchString := in.GetQueryString()
	searchOptions := fmt.Sprintf("%s%sH", "-", in.GetArgs())
	searchFile := in.GetQueryFile()

	returnChan := make(chan string)
	defer close(returnChan)

	// Concurrently connect to all VMs and perform grep
	for _, ip := range machine_ips {
		go connect(searchString, searchOptions, searchFile, returnChan, ip)
	}
	var resp pb.CoordinatorResponse
	regex := re.MustCompile(`.*:([0-9]+)`)
	var total_matches int
	for k := 0; k < len(machine_ips); k++ {
		resp.Result = append(resp.Result, <-returnChan)
		for _, match := range regex.FindAllStringSubmatch(resp.Result[k], -1) {
			matches, err := strconv.Atoi(match[1])
			if err != nil {
				log.Printf(err.Error())
			}
			total_matches += matches
		}
	}
	resp.Total = int64(total_matches)
	return &resp, nil
}

/* this function calls the grep rpc on the workers
 */
func connect(searchString string, searchOptions string, searchFile string, returnChan chan string, ip string) {
	fmt.Printf("Connecting to %s\n", ip)
	conn, err := grpc.Dial(ip, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGrepInterfaceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	r, err := c.Grep(ctx, &pb.GrepRequest{
		SearchString:  searchString,
		SearchOptions: searchOptions,
		SearchFile:    searchFile})
	if err != nil {
		log.Printf("could not grep!: %v", err)
	}
	log.Printf("GrepResponse: %s", r.GetGrepOutput())
	returnChan <- r.GetGrepOutput()
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Printf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDistributedGrepServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
	}
}
