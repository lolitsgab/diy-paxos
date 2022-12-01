package server

import (
	"context"
	"log"

	pb "diy-paxos/diypaxos/proto"
)

// Server implements the SimpleKvStore server.
type Server struct {
	pb.UnimplementedSimpleKvStoreServer
}

// Get a value by key.
func (s *Server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	log.Printf("Received a Get request.")
	return &pb.GetResponse{}, nil
}

// Insert a KV pair.
func (s *Server) Insert(ctx context.Context, in *pb.InsertRequest) (*pb.InsertResponse, error) {
	log.Printf("Received an Insert request.")
	return &pb.InsertResponse{}, nil
}

// Remove a KV pair.
func (s *Server) Remove(ctx context.Context, in *pb.RemoveRequest) (*pb.RemoveResponse, error) {
	log.Printf("Received a Remove request.")
	return &pb.RemoveResponse{}, nil
}

// Update a KV pair.
func (s *Server) Update(ctx context.Context, in *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	log.Printf("Received an Update request.")
	return &pb.UpdateResponse{}, nil
}

// Upsert updates or inserts a KV pair.
func (s *Server) Upsert(ctx context.Context, in *pb.UpsertRequest) (*pb.UpsertResponse, error) {
	log.Printf("Received an Upsert request.")
	return &pb.UpsertResponse{}, nil
}
