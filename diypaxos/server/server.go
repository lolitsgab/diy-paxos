package server

import (
	"context"
	pb "diy-paxos/diypaxos/proto"
	"diy-paxos/diypaxos/storage"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type KvStoreServer interface {
	string
	Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error)
	Insert(ctx context.Context, in *pb.InsertRequest) (*pb.InsertResponse, error)
	Remove(ctx context.Context, in *pb.RemoveRequest) (*pb.RemoveResponse, error)
	Update(ctx context.Context, in *pb.UpdateRequest) (*pb.UpdateResponse, error)
	Upsert(ctx context.Context, in *pb.UpsertRequest) (*pb.UpsertResponse, error)
	Accept(ctx context.Context, in *pb.AcceptRequest) (*pb.AcceptResponse, error)
	Promise(ctx context.Context, in *pb.PromiseRequest) (*pb.PromiseResponse, error)
}

// Server implements the SimpleKvStore Server.
type Server struct {
	Addr            string
	storage         storage.Storage
	headlessService string
	Replicas        []string
	promises        map[string]storage.Value
}

func LogAndReturnError(code codes.Code, format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args)
	log.Printf(msg)
	return status.New(code, msg).Err()
}

// NewServer generates a new Server using the provided Storage as a Storage backend.
func NewServer(name, headlessServer string, store storage.Storage) *Server {
	if name == "" || store == nil {
		panic("Name and Store required.")
	}
	return &Server{Addr: name, storage: store, headlessService: headlessServer}
}

// Get a value by key.
func (s *Server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	log.Printf("Received a Get(%v) request.", in)
	resp, err := s.storage.Get(in.GetKey())
	if err != nil {
		log.Printf("error in Get(%v, %v): %v", ctx, in, err)
		return nil, err
	}
	return &pb.GetResponse{Val: resp}, nil
}

// Insert a KV pair.
func (s *Server) Insert(ctx context.Context, in *pb.InsertRequest) (*pb.InsertResponse, error) {
	log.Printf("Received an Insert(%v) request.", in)
	err := s.storage.Insert(in.GetKey(), in.GetVal())
	if err != nil {
		log.Printf("error in Insert(%v, %v): %v", ctx, in, err)
		return nil, err
	}
	return &pb.InsertResponse{}, nil
}

// Remove a KV pair.
func (s *Server) Remove(ctx context.Context, in *pb.RemoveRequest) (*pb.RemoveResponse, error) {
	log.Printf("Received a Remove(%v) request.", in)
	if err := s.storage.Remove(in.GetKey()); err != nil {
		log.Printf("error in Remove(%v, %v): %v", ctx, in, err)
		return nil, err
	}
	return &pb.RemoveResponse{}, nil
}

// Update a KV pair.
func (s *Server) Update(ctx context.Context, in *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	log.Printf("Received an Update(%v) request.", in)
	if err := s.storage.Update(in.GetKey(), in.GetVal()); err != nil {
		log.Printf("error in Remove(%v, %v): %v", ctx, in, err)
		return nil, err
	}
	return &pb.UpdateResponse{}, nil
}

// Upsert updates or inserts a KV pair.
func (s *Server) Upsert(ctx context.Context, in *pb.UpsertRequest) (*pb.UpsertResponse, error) {
	log.Printf("Received an Upsert(%v) request.", in)
	if err := s.storage.Upsert(in.GetKey(), in.GetVal()); err != nil {
		log.Printf("error in Remove(%v, %v): %v", ctx, in, err)
		return nil, err
	}
	return &pb.UpsertResponse{}, nil
}
