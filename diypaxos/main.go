package main

import (
	"diy-paxos/diypaxos/server"
	"diy-paxos/diypaxos/storage"
	"flag"
	"fmt"
	"log"
	"net"

	pb "diy-paxos/diypaxos/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 8080, "The server port")
	name = flag.String("name", "", "The server hostname")
)

func main() {
	flag.Parse()
	if *name == "" {
		panic("--name flag required")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	srv := server.NewServer("server1", storage.NewInMemoryStorage())
	pb.RegisterSimpleKvStoreServer(s, srv)
	reflection.Register(s)
	log.Println("++======================++")
	log.Printf("server %v listening at %v", *name, *port)
	log.Println("++======================++")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
