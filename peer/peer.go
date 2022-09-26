// Package main implements the peer and introducer for a failure detector system.
// NOTE: logs of kind RL/L/U/RU with indices are to check if a lock has been acquired/released at
// a particular point
package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sort"
	"sync"
	"time"

	"github.com/chrusty/go-tableprinter"

	"google.golang.org/grpc/credentials/insecure"

	pb "mp1/protos"

	"golang.org/x/exp/slices"
	"google.golang.org/grpc"
)

const (
	TCP_PORT        = 50053
	UDP_PORT_LISTEN = 50054
	UDP_PORT_DIAL   = 0
	LOG_DIR         = "/mp1/logs/"
	LOCAL_ADDR      = "192.168.64.2:50053"
	VM_ADDR         = "fa22-cs425-0301:50053"
	VM_PREFIX       = "fa22-cs425-03"
)

type server struct {
	pb.UnimplementedIntroInterfaceServer
}

// member structure which is part of the membership list, this structure has
// timers which are important for the suspiciion mechanism timeouts
type member struct {
	HostName    string
	Timestamp   int
	Status      int //0 -> dead 1 -> sus 2 -> alive
	Incarnation int
	Timer       *time.Timer
}

// trimmed down version of the member structure that doesn't contain timers,
// to be sent over udp
type UDPmember struct {
	HostName    string
	Timestamp   int
	Status      int //0 -> dead 1 -> sus 2 -> alive
	Incarnation int
}

// empty struct being sent to ping the udp-server
type UDPSyn struct {
}

// return struct sent by udp-server to the client, containing memberlist
type UDPAck struct {
	MemberList []UDPmember
}

// variable to switch log dir location when using local machinevs vms
var addr = VM_ADDR

// variable to store peers ipaddr
var myHostName string

// membership list that each peer maintains
var memberList []member = make([]member, 0, 20)

// read-write mutex to lock access to membership list, which can be used by
// multiple threads
var lock sync.RWMutex
var leaving = false

// function to find the neighbours of a given host
func findNeighbours(hostname string) []string {
	var returnList []string
	log.Println("IN findNeighbours")
	log.Println("1 RL try")
	lock.RLock()
	log.Println("1 RL done")
	n := len(memberList)
	idx := slices.IndexFunc(memberList, func(c member) bool { return c.HostName == hostname })
	if n != 0 && (idx+1)%n != idx {
		returnList = append(returnList, memberList[(idx+1)%n].HostName)
	}
	if n != 0 && (idx+n-1)%n != idx && memberList[(idx+n-1)%n].HostName != returnList[0] {
		returnList = append(returnList, memberList[(idx+n-1)%n].HostName)
	}
	idx2 := slices.IndexFunc(returnList, func(c string) bool { return c == memberList[(idx+n/2)%n].HostName })
	if n != 0 && idx != (idx+n/2)%n && idx2 == -1 {
		returnList = append(returnList, memberList[(idx+n/2)%n].HostName)
	}
	log.Println("2 RU try")
	lock.RUnlock()
	log.Println("2 RU done")
	log.Println("OUT findNeighbours", returnList)
	return returnList
}

// fucntion to convert member struct info to protobuf peer struct
func toMemberMsg(idx int, n int) *pb.Peer {
	return &pb.Peer{Hostname: memberList[(idx)%n].HostName,
		Timestamp:   int64(memberList[(idx)%n].Timestamp),
		Status:      int64(memberList[(idx)%n].Status),
		Incarnation: int64(memberList[(idx)%n].Incarnation)}
}

// function to convert protobuf peer array into membership list
func toMemberList(inPeers []*pb.Peer) []member {
	var members []member
	for _, ele := range inPeers {
		members = append(members, member{
			ele.Hostname, int(ele.Timestamp), int(ele.Status), int(ele.Incarnation), nil,
		})
	}
	return members
}

// function to calculate the members to be given by the introducer to the
// incoming peer requesting peers
func introNeighbours(hostname string) []*pb.Peer {
	log.Println("IN introneighbours")
	var peers []*pb.Peer
	log.Println("3 RL try")
	lock.RLock()
	log.Println("3 RL DONE")
	n := len(memberList)
	idx := slices.IndexFunc(memberList, func(c member) bool { return c.HostName == hostname })
	peers = append(peers, toMemberMsg(idx, n))
	idx2 := slices.IndexFunc(peers, func(c *pb.Peer) bool { return c.Hostname == memberList[(idx+1)%n].HostName })
	if idx2 == -1 {
		peers = append(peers, toMemberMsg(idx+1, n))
	}
	idx2 = slices.IndexFunc(peers, func(c *pb.Peer) bool { return c.Hostname == memberList[(idx+n-1)%n].HostName })
	if idx2 == -1 {
		peers = append(peers, toMemberMsg(idx+n-1, n))
	}
	idx2 = slices.IndexFunc(peers, func(c *pb.Peer) bool { return c.Hostname == memberList[(idx+n/2)%n].HostName })
	if idx2 == -1 {
		peers = append(peers, toMemberMsg(idx+n/2, n))
	}
	log.Println("4 RU try")
	lock.RUnlock()
	log.Println("4 RU done")
	log.Println("OUT introneighbours")
	return peers
}

// this rpc is called by a peer to request introduction into the p2p system
func (s *server) RequestIntro(ctx context.Context, in *pb.IntroRequest) (*pb.IntroResponse, error) {
	var resp pb.IntroResponse
	log.Println("IN requestintro")
	log.Println("5 L try")
	lock.Lock()
	log.Println("5 L done")
	memberList = append(memberList, member{HostName: in.GetHostName(),
		Timestamp:   int(in.GetTimestamp()),
		Status:      2,
		Incarnation: int(in.GetIncarnation())})
	sort.SliceStable(memberList, func(i, j int) bool {
		return memberList[i].HostName < memberList[j].HostName
	})
	log.Println("6 U try")
	lock.Unlock()
	log.Println("6 U done")
	resp.Peers = introNeighbours(in.GetHostName())
	sort.SliceStable(resp.Peers, func(i, j int) bool {
		return resp.Peers[i].Hostname < resp.Peers[j].Hostname
	})
	log.Printf("PEERS : %v", resp.Peers)
	log.Printf("Returned Intro Response %v\n", resp.Peers)
	log.Println("OUT introneighbours", memberList)
	return &resp, nil
}

// this function is called to remove a given host from the membership list
func RemoveIndex(hostname string) {
	log.Println("IN removeindex", hostname)
	idx := slices.IndexFunc(memberList, func(n member) bool {
		return hostname == n.HostName
	})
	if idx != -1 {
		memberList = append(memberList[:idx], memberList[idx+1:]...)
	}

	// log.Println("%v", findNeighbours(myHostName))
	log.Println("OUT removeindex")
}

// this function works as a statemachine to transition through the various states
// of a peer in the suspicion mechanism
func waitOnTimerAsync(hostname string, timeoutMarker int) {
	log.Println("IN waitOnTimerAsync", timeoutMarker, hostname, memberList)
	// stopTimer(m)
	go func() {
		log.Println("Reached to wait\n")
		log.Println("99 RLOCK try")
		lock.RLock()
		log.Println("99 RLOCK done", memberList)
		idx := slices.IndexFunc(memberList, func(m member) bool {
			return m.HostName == hostname
		})
		log.Println("24 U try", hostname)
		if idx == -1 {
			log.Println("idx -1 exit", hostname)
			lock.RUnlock()
			return
		}
		if memberList[idx].Timer == nil {
			log.Println("timer nil exit", hostname)
			lock.RUnlock()
			return
		}
		lock.RUnlock()
		log.Println("waiting on timer")
		<-memberList[idx].Timer.C
		log.Println("finished waiting on timer")
		if timeoutMarker == 0 { // Delete Item from list
			// Remove the element at index i from m
			log.Println("7 L try")
			lock.Lock()
			log.Println("7 L done")
			log.Printf("Deleting member %v \n", hostname)
			RemoveIndex(hostname)
			log.Println("8 U try")
			lock.Unlock()
			log.Println("8 U done")
		} else if timeoutMarker == 1 { // Dead -> Cleanup
			log.Println("9 L try")
			lock.Lock()
			idx := slices.IndexFunc(memberList, func(m member) bool {
				return m.HostName == hostname
			})

			if idx == -1 {
				log.Println("idx -1 exit", hostname)
				lock.Unlock()
				return
			}
			log.Println("9 L done", memberList, memberList[idx])
			memberList[idx].Status = 0
			log.Printf("Starting cleanup timer for %v \n", memberList[idx].HostName)
			if memberList[idx].Timer != nil {
				memberList[idx].Timer.Stop()
			}
			memberList[idx].Timer = time.NewTimer(1 * time.Second)
			log.Println("10 U try")
			lock.Unlock()
			log.Println("10 U done", memberList, memberList[idx])
			log.Println("ADDRESS %p %p\n", memberList, memberList[idx])
			waitOnTimerAsync(hostname, 0)
		} else if timeoutMarker == 2 { // Sus -> Dead
			log.Println("11 L try")
			lock.Lock()
			idx := slices.IndexFunc(memberList, func(m member) bool {
				return m.HostName == hostname
			})
			if idx == -1 {
				log.Println("idx -1 exit", hostname)
				lock.Unlock()
				return
			}
			log.Println("11 L done")
			log.Printf("Starting dead timer for %v \n", memberList[idx].HostName)
			if memberList[idx].Timer != nil {
				memberList[idx].Timer.Stop()
			}
			memberList[idx].Timer = time.NewTimer(time.Millisecond * 500)
			log.Println("12 U try")
			lock.Unlock()
			log.Println("12 U done")
			waitOnTimerAsync(hostname, 1)
		}
	}()
	log.Println("OUT waitOnTimerAsync", hostname)
}

// function to stop the timer for a member
func stopTimer(m *member) {
	log.Println("IN stoptimer")
	m.Timer.Stop()
	log.Println("OUT stoptimer", m.Timer)
}

// convert udpmember struct into member struct
func toMember(u UDPmember) member {
	return member{HostName: u.HostName,
		Timestamp:   u.Timestamp,
		Status:      u.Status,
		Incarnation: u.Incarnation,
	}
}

// convert member struct into udpmember struct
func toUDPMember(m member) UDPmember {
	return UDPmember{HostName: m.HostName,
		Timestamp:   m.Timestamp,
		Status:      m.Status,
		Incarnation: m.Incarnation,
	}
}

// function to merge peer's membershiplist with incoming membership list
// received using the ping-pong and suspicion mechanism
func mergeMembershipList(inList []UDPmember) {
	log.Println("IN mergeMembershipList")

	var newList []member
	var newMember member
	var timerFlag []bool = make([]bool, len(memberList)+len(inList))
	i := 0
	j := 0

	log.Println("13 L try")
	lock.Lock()
	log.Println("13 L done")
	for i < len(memberList) && j < len(inList) {
		log.Println("member", memberList[i])
		log.Println("inmember", inList[j])
		if memberList[i].HostName == inList[j].HostName {
			if memberList[i].Incarnation < inList[j].Incarnation { //self lower incarnation
				if memberList[i].Status != 0 { //if i am not dead
					if memberList[i].Status != inList[j].Status {
						//start timer to death if inStatus = sus and memberList!=sus or
						if inList[j].Status == 2 && memberList[i].Timer != nil {
							stopTimer(&memberList[i])
						}
						memberList[i].Status = inList[j].Status
						//start remove timer if inStatus = death and memberList!=death
						if inList[j].Status == 0 || inList[j].Status == 1 {
							log.Println("14 U try")
							lock.Unlock()
							log.Println("14 U done")
							memberList[i].Timer = time.NewTimer(0)
							timerFlag[len(newList)] = true
							log.Println("15 L try")
							lock.Lock()
							log.Println("15 L done")
						}
					}
					memberList[i].Incarnation = inList[j].Incarnation
				}
			} else { // self >= in incarnation
				if inList[j].Status == 0 && memberList[i].Status != 0 {
					memberList[i].Status = inList[j].Status
					//start remove timer
					log.Println("16 U try")
					memberList[i].Timer = time.NewTimer(0)
					lock.Unlock()
					log.Println("16 U done")
					timerFlag[len(newList)] = true
					log.Println("17 L try")
					lock.Lock()
					log.Println("17 L done")
				}
			}
			newList = append(newList, memberList[i])
			i += 1
			j += 1
		} else if memberList[i].HostName < inList[j].HostName {

			newList = append(newList, memberList[i])
			i += 1
		} else {
			newMember = toMember(inList[j])
			if newMember.Status == 2 {
				newList = append(newList, newMember)
			}
			if inList[j].Status != 2 {
				// start remove timer if inStatus = death
				log.Println("18 U try")
				newMember.Timer = time.NewTimer(0)
				lock.Unlock()
				log.Println("18 U done")
				timerFlag[len(newList)] = true
				log.Println("19 U try")
				lock.Lock()
				log.Println("19 U done")
			}
			j += 1
		}
	}
	for i < len(memberList) {
		log.Println("member", memberList[i])
		newList = append(newList, memberList[i])
		i += 1
	}
	for j < len(inList) {
		log.Println("inList", inList[j])
		newMember = toMember(inList[j])
		if newMember.Status == 2 {
			newList = append(newList, newMember)
		}
		// start remove timer if inStatus = death
		if newMember.Status == 0 || newMember.Status == 1 {
			log.Println("20 U try", newMember.HostName)
			newMember.Timer = time.NewTimer(0)
			lock.Unlock()
			log.Println("20 U done")
			timerFlag[len(newList)] = true
			log.Println("21 L try")
			lock.Lock()
			log.Println("21 L done")
		}
		j += 1
	}
	for i, ele := range newList {
		if i < len(memberList) {
			memberList[i] = ele
		} else {
			memberList = append(memberList, ele)
		}
		if timerFlag[i] {
			hostname := memberList[i].HostName
			status := memberList[i].Status + 1
			lock.Unlock()
			waitOnTimerAsync(hostname, status)
			lock.Lock()
		}

	}
	log.Println("22 U try")
	lock.Unlock()
	log.Println("22 U done")
	log.Println("OUT mergeMembershipList")
}

/*this rpc is called using the test_client to delete previously created known logs in all of the vms
 */
func (s *server) InformDeath(ctx context.Context, in *pb.DeathInfo) (*pb.DeathAck, error) {
	log.Println("Received Inform Death Request\n")
	return &pb.DeathAck{}, nil
}

// start the introducer server
func startIntroducerServer() {
	log.Println("IN startIntroducerServer")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", TCP_PORT))
	if err != nil {
		log.Printf("failed to listen: %v\n", err)
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterIntroInterfaceServer(s, &server{})
	log.Printf("introducer listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v\n", err)
	}
	log.Println("OUT startIntroducerServer")
}

func list_mem() {
	lock.RLock()
	log.Println("LIST_MEM")
	log.Printf("%+v", memberList)
	if len(memberList) == 0 {
		fmt.Println("Empty Membership List!")
		lock.RUnlock()
		return
	}
	pretty_print(memberList)
	lock.RUnlock()
}

func pretty_print(x interface{}) {
	printer := tableprinter.New().WithBorders(true)
	// Use the custom printer to print the examples:
	if err := printer.Print(x); err != nil {
		panic(err)
	}
}

func list_self() {
	lock.RLock()
	idx := slices.IndexFunc(memberList, func(m member) bool {
		return m.HostName == myHostName
	})
	log.Printf("%d", idx)
	if idx != -1 {
		log.Printf("ID of Self %v", memberList[idx])
	} else {
		log.Printf("Not part of any cluster yet!")
		lock.RUnlock()
		return
	}
	pretty_print(memberList[idx])
	lock.RUnlock()
}

// function to join the p2p network
func join(args []string) {
	hostname, err := os.ReadFile("/mp1/hostname")
	myHostName = string(hostname[:len(hostname)-1])
	fmt.Println("Joining the distributed cluster! ", myHostName)
	timestamp := time.Now().Nanosecond()
	if err != nil {
		panic(err)
	}
	log.Println(string(hostname))
	//args := os.Args[1:]
	if len(args) >= 1 && args[0] == "i" {
		go startIntroducerServer()
		memberList = append(memberList, member{HostName: myHostName,
			Timestamp:   timestamp,
			Status:      2,
			Incarnation: 1})
	} else {
		conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("did not connect: %v\n", err)
		}
		defer conn.Close()
		c := pb.NewIntroInterfaceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.RequestIntro(ctx, &pb.IntroRequest{HostName: myHostName,
			Incarnation: 1,
			Timestamp:   int64(timestamp)})
		if err != nil {
			log.Printf("could not send grpc req response!: %v\n", err)
		} else {
			memberList = toMemberList(r.GetPeers())
			log.Printf("%v %v\n", r.GetPeers(), memberList)
		}
	}
	go startPeerUDPServer()
	startPeerUDPClient()
}

func leave() {
	log.Println("Leaving the distributed cluster! :( ")
	leaving = true
}

func main() {
	flag.Parse()
	fileName := LOG_DIR + "dist_log.log"
	f, err := os.OpenFile(fileName,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)
	args := os.Args[1:]
	go func() {
		for {
			buf := bufio.NewReader(os.Stdin)
			fmt.Print("> ")
			sentence, err := buf.ReadBytes('\n')
			if err != nil {
				log.Println(err)
			} else {
				log.Println(string(sentence))
				if string(sentence) == "list_mem\n" {
					list_mem()
				} else if string(sentence) == "list_self\n" {
					list_self()
				} else if string(sentence) == "join\n" {
					lock.Lock()
					memberList = memberList[:0]
					lock.Unlock()
					go join(args)
				} else if string(sentence) == "leave\n" {
					lock.Lock()
					memberList = memberList[:0]
					lock.Unlock()
					go leave()
				}
			}
		}
	}()
	for {
	}
}

func startPeerUDPServer() {
	log.Println("IN startPeerUDPServer")
	p := make([]byte, 2048)
	log.Printf("PLEASE PRINT ME LISTENING AT %v\n", UDP_PORT_LISTEN)
	udp_addr := net.UDPAddr{
		Port: UDP_PORT_LISTEN,
		IP:   net.ParseIP(myHostName),
	}
	ser, err := net.ListenUDP("udp", &udp_addr)
	if err != nil {
		log.Printf("Some error %v\n", err)
		return
	}
	for {
		if leaving {
			return
		}
		var x UDPSyn
		log.Println("23 L try")
		lock.Lock()
		log.Println("23 L done")
		log.Printf("%v server memlist", memberList)
		idx := slices.IndexFunc(memberList, func(m member) bool {
			return m.HostName == myHostName
		})
		log.Println("24 U try")
		if idx == -1 {
			lock.Unlock()
			continue
		}
		memberList[idx].Incarnation += 1
		memberList[idx].Status = 2
		lock.Unlock()
		log.Println("24 U done")
		var returnList []UDPmember
		log.Println("25 R try")
		lock.RLock()
		log.Println("25 R done")
		for _, ele := range memberList {
			returnList = append(returnList, toUDPMember(ele))
		}
		log.Println("26 RU try")
		lock.RUnlock()
		log.Println("26 RU done")

		y := UDPAck{MemberList: returnList}
		out, err := json.Marshal(&y)
		if err != nil {
			panic(err)
		}

		log.Println("\n\nUDP MembershipList out : %s\n\n", string(out))
		n, remoteaddr, err := ser.ReadFromUDP(p)
		if err := gob.NewDecoder(bytes.NewReader(p[:n])).Decode(&x); err != nil {
			// handle error
			panic(err)

		}
		log.Printf("Read a SYN from %v %v\n", remoteaddr, x)
		if err != nil {
			log.Printf("Some error  %v\n", err)
			continue
		}
		go sendAckFromServer(ser, remoteaddr, out)

	}
	log.Println("OUT startPeerUDPServer")
}

func sendAckFromServer(conn *net.UDPConn, udp_addr *net.UDPAddr, response []byte) {
	_, err := conn.WriteToUDP(response, udp_addr)
	if err != nil {
		log.Printf("Couldn't send response %v\n", err)
	}
}

// function to mark member as suspicious
func markMemberAsSuspicious(neighbour string) {
	log.Println("IN markMemberAsSuspicious", neighbour)
	log.Println("27 L try")
	lock.Lock()
	log.Println("27 L done")
	idx := slices.IndexFunc(memberList, func(m member) bool {
		return m.HostName == neighbour
	})
	if idx == -1 || memberList[idx].Status == 1 || memberList[idx].Status == 0 {
		log.Println("29 U try")
		lock.Unlock()
		log.Println("29 U try")
		return
	}
	hostname := memberList[idx].HostName
	//HERE BREAKS
	if memberList[idx].Status != 1 {
		memberList[idx].Status = 1
		if memberList[idx].Timer != nil {
			memberList[idx].Timer.Stop()
		}
		memberList[idx].Timer = time.NewTimer(time.Millisecond * 500)
		log.Println("sus timer start", memberList[idx])
	}
	log.Println("28 U try", memberList[idx])
	lock.Unlock()
	log.Println("28 U done")
	waitOnTimerAsync(hostname, 2)
	log.Println("OUT markMemberAsSuspicious", neighbour)
}

func startPeerUDPClient() {
	log.Println("IN startPeerUDPClient")
	log.Println("Starting the ticker\n")
	ticker := time.NewTicker(300 * time.Millisecond)
	retChan := make(chan []byte, 10000)
	for _ = range ticker.C {
		if leaving {
			log.Println("Gracefully leaving the process")
			ticker.Stop()
			break
		}
		t := findNeighbours(myHostName)
		lock.RLock()
		log.Println("\n\n\n\nMEMBERSHIP ", memberList)
		lock.RUnlock()
		log.Print("TICKER", t)
		for i := 0; i < len(t); i++ {
			toAddr := net.UDPAddr{
				Port: UDP_PORT_LISTEN,
				IP:   net.ParseIP(t[i]),
			}
			fromAddr := net.UDPAddr{
				Port: UDP_PORT_DIAL,
				IP:   net.ParseIP(myHostName),
			}
			conn, err := net.DialUDP("udp", &fromAddr, &toAddr)
			if err != nil {
				log.Println("UDP FAIL", err.Error())
				// retChan <- []byte("Network Partitioned")
				continue
			}
			log.Println("wait for ACK", t[i])
			go sendSynAndWaitForAck(retChan, t[i], conn)
			select {
			case ret := <-retChan:
				if string(ret) == "Network Partitioned" {
					log.Println("No news in 500 milliseconds due to partition.", t[i])
					markMemberAsSuspicious(t[i])
				} else if string(ret) != "CONN CLOSED" {
					log.Println("Got Ack %s \n", string(ret))
					var response UDPAck
					ret = bytes.Trim(ret, "\x00")
					err := json.Unmarshal(ret, &response)
					if err != nil {
						log.Println("UDP FAIL", err.Error())
						retChan <- []byte("IDK WHAT ")
						continue
					}
					mergeMembershipList(response.MemberList)
					log.Printf("%v\n", t)
				}
			case <-time.After(500 * time.Millisecond):
				log.Println("No news in 500 milliseconds due to timeout.", t[i])
				closeConn(conn, retChan, t[i])
				markMemberAsSuspicious(t[i])
			}
		}
	}
	log.Println("OUT startPeerUDPClient")
}

func closeConn(conn *net.UDPConn, retChan chan []byte, hostname string) {
	log.Println("IN closeConn", hostname)
	if conn != nil {
		err := conn.Close()
		if err != nil {
		}
	}
	log.Println("OUT closeConn", hostname)
}

// function to send a ping and wait for pong
func sendSynAndWaitForAck(retChan chan []byte, hostname string, conn *net.UDPConn) {
	log.Println("IN sendSynAndWaitForAck", hostname)
	defer closeConn(conn, retChan, hostname)
	//Send Syn
	gob.Register(UDPSyn{})
	var x UDPSyn
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(x); err != nil {
		panic(err)
	}
	log.Println("IN writing to buffer")
	_, err := conn.Write(buf.Bytes())
	if err != nil {
		log.Println("PKK", err.Error(), hostname)
		retChan <- []byte("Network Partitioned")
		return
	}
	//Wait For Ack
	p := make([]byte, 2048)
	log.Println("IN reading from buffer")
	_, err = bufio.NewReader(conn).Read(p)
	if err == nil {
		log.Println("\nUDP Membership List in %s\n", p)
	} else {
		log.Println("PKK", err.Error(), hostname)
		retChan <- []byte("Network Partitioned")
		return
	}
	log.Println("OUT sendSynAndWaitForAck")
	retChan <- p
}

func makeUDPConnection(hostname string) *net.UDPConn {
	log.Println("\nMaking Connection")
	toAddr := net.UDPAddr{
		Port: UDP_PORT_LISTEN,
		IP:   net.ParseIP(hostname),
	}
	fromAddr := net.UDPAddr{
		Port: UDP_PORT_DIAL,
		IP:   net.ParseIP(myHostName),
	}
	conn, err := net.DialUDP("udp", &fromAddr, &toAddr)
	if err != nil {
		log.Printf("Some error %v", err)
		return nil
	}
	return conn
}
