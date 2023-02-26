package main

import (
	pb "diy-paxos/diypaxos/proto"
	"diy-paxos/diypaxos/server"
	"diy-paxos/diypaxos/storage"
	"diy-paxos/diypaxos/utils"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"time"
)

var port = flag.Int("port", 8080, "The server port")
var replicas = flag.String("replicas", "", "list of replicas to use, ignoring discovery-host: ip1:port")
var discovery_host = flag.String("discovery-host", "headless-kvstore", "hostname of discovery service")
var singleton = flag.Bool("singleton", false, "enable if this is a single node with no replicas.")

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

	srv := server.NewServer(hostname, *port, *discovery_host, storage.NewInMemoryStorage())
	s := grpc.NewServer(grpc.UnaryInterceptor(server.TeeInterceptor(srv)))
	pb.RegisterSimpleKvStoreServer(s, srv)
	reflection.Register(s)
	if err := initReplicas(srv); err != nil {
		panic("could not load replicas" + err.Error())
	}
	log.Println("++======================++")
	log.Printf("server %v listening at %v", hostname, srv.Addr)
	log.Println("++======================++")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initReplicas(srv *server.Server) error {
	if !*singleton && *replicas != "" {
		reps, err := utils.ParseReplicaString(*replicas)
		srv.Replicas = reps
		return err
	}
	if !*singleton && *discovery_host != "" {
		reps, err := utils.DiscoverReplicas(*discovery_host, srv.Addr, 10, time.Millisecond*10)
		srv.Replicas = reps
		return err
	}
	log.Printf("using sigleton mode")
	return nil
}
