# CS425 MP1
## Name
This repository contains the code for running a distributed log querier. It is predominantly written in go, with some testing scripts written in python.
## Description
With this project, given a cluster of machines in a distributed cluster, you will be able to run a distributed grep command that greps all the
individual machines and written a consolidated output efficiently.

## Installation
Refer to setup.sh to install libraries needed to run this repo.
Preqrequisites are
- protobuf-compiler
- protoc-gen-go@v1.28 
- protoc-gen-go-grpc@v1.2

The setup.sh, when run from an appropriate location, will install required libraries and clone the repo from each of the VMs

## Usage
To generate code from gRPC protobufs. Note : This has to be run before any client/servers are started.
```
./generate_protobuf.sh
#FAQ - If you hit issues here you can try to export GO_PATH in PATH using 'export PATH="$PATH:$(go env GOPATH)/bin'
```
To run the workers/server processes from the root of the repo
```
go run worker/worker_server.go
```
To run the coordinator process from the root of the repo
```
go run coordinator/coordinator.go
```
Different query options supported on distributed grep client
```
go run client/client.go c '<pattern_to_search>' '<files_to_search>'
Here c is a grep option to return no. of matches
```
```
go run client/client.go Ec '<regex_to_search>' '<files_to_search>'
Here E is a grep option to run grep on regular expressions
```
## Test Client
You can run integration tests that test a variety of scenarios.
We test the following currently :
TestDistributedGrepInfrequentWord()
TestDistributedGrepFrequentWord()
TestDistributedGrepRegex()
TestDistributedGrepNonExistentWord()
TestDistributedGrepKillRandomWorker()
To run the test client you can run the following commands. 
```
cd tests
python3 integration_test.py
```
Additionally, to just generate test logs, we have a test_client that you can use.
The test script integration_test.py uses these scripts internally to create and delete logs
for every test run, so these need not be run explicitly.
To create test logs , from the root of the repo run
```
go run test_client/test_client.go create
```
To delete the test logs
```
go run test_client/test_client.go delete
```

## CS 425 MP2
This repository also contains the code for running a failure detector system(peer director). It is written in go.
## Description
With this project, we orchestrate a p2p system of nodes that are able to detect when one of them fails. We use a SWIM style algorithm with a suspicion mechanism where each node pings its neighbours and checks if they're alive. 
##Setup
Hostname must be provided to the peer.go file via a file called hostname.
## Usage
### To start Introducer
```go run peer/peer.go i```
### To start Peers
```go run peer/peer.go```

Once the process is up, we can use the cli following cli:
- join : to join the system
- leave : to leave the system
- list_mem : to display membership list
- list_self : to display node details



