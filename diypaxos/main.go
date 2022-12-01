package main

import (
	"diy-paxos/diypaxos/server"
	"flag"
	"fmt"
	"log"
	"net"

	pb "diy-paxos/diypaxos/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var port = flag.Int("port", 8080, "The server port")

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSimpleKvStoreServer(s, &server.Server{})
	reflection.Register(s)
	log.Printf("server listening at %v", *port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
