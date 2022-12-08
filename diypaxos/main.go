package main

import (
	"diy-paxos/diypaxos/server"
	"diy-paxos/diypaxos/storage"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	pb "diy-paxos/diypaxos/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var port = flag.Int("port", 8080, "The server port")

func main() {
	flag.Parse()
	hostname, err := os.Hostname()
	if hostname == "" {
		panic(fmt.Sprintf("No hostname provided: {%v}", hostname))
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	srv := server.NewServer(hostname, "headless-kvstore", storage.NewInMemoryStorage())
	pb.RegisterSimpleKvStoreServer(s, srv)
	reflection.Register(s)
	if err := srv.GetReplicaIPs(10, time.Second*2); err != nil {
		log.Panicf("could not fetch replicas: %v", err)
	}
	log.Println("++======================++")
	log.Printf("server %v listening at %v", hostname, *port)
	log.Println("++======================++")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
