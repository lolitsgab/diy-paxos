package server

import (
	"context"
	pb "diy-paxos/diypaxos/proto"
	"diy-paxos/diypaxos/storage"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math"
	"net"
	"os"
	"sync"
)

type KvStoreServer interface {
	string
	Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error)
	Insert(ctx context.Context, in *pb.InsertRequest) (*pb.InsertResponse, error)
	Remove(ctx context.Context, in *pb.RemoveRequest) (*pb.RemoveResponse, error)
	Update(ctx context.Context, in *pb.UpdateRequest) (*pb.UpdateResponse, error)
	Upsert(ctx context.Context, in *pb.UpsertRequest) (*pb.UpsertResponse, error)
	Accept(ctx context.Context, in *pb.AcceptRequest) (*pb.AcceptResponse, error)
	Prepare(ctx context.Context, in *pb.PrepareRequest) (*pb.PrepareResponse, error)
}

// Server implements the SimpleKvStore Server.
type Server struct {
	Addr               string
	Name               string
	Id                 int
	Round              float64
	LeaderName         string
	storage            storage.Storage
	headlessService    string
	Replicas           []string
	ReplicaConnections map[string]pb.SimpleKvStoreClient
	promises           map[string]storage.Value
	leaderMu           sync.Mutex
	roundMu            sync.Mutex
}

var logger *log.Logger = log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds)

func (s *Server) Prepare(ctx context.Context, in *pb.PrepareRequest) (*pb.PrepareResponse, error) {
	logger.Printf("received prepare %v", in)
	s.leaderMu.Lock()
	defer s.leaderMu.Unlock()
	s.roundMu.Lock()
	defer s.roundMu.Unlock()

	resp := &pb.PrepareResponse{
		Name:             s.Name,
		Promise:          s.Round <= in.Round,
		HighestRoundSeen: math.Max(s.Round, in.Round),
		Val:              s.getLeaderVal(),
	}
	s.Round = math.Max(s.Round, in.Round)
	logger.Printf("responded to prepare %v", resp)
	return resp, nil
}

func (s *Server) Accept(ctx context.Context, in *pb.AcceptRequest) (*pb.AcceptResponse, error) {
	logger.Printf("received accept %v", in)
	s.leaderMu.Lock()
	defer s.leaderMu.Unlock()
	s.roundMu.Lock()
	defer s.roundMu.Unlock()

	if in.GetRound() >= s.Round {
		if s.LeaderName == "" {
			s.LeaderName = string(in.GetVal())
		}
		s.Round = in.GetRound()
	}
	resp := &pb.AcceptResponse{
		Name:     s.Name,
		Accepted: in.GetRound() >= s.Round,
	}
	logger.Printf("responded to accept: %v and leader is %v", resp, s.LeaderName)
	return resp, nil
}

func (s *Server) getLeaderVal() []byte {
	if s.LeaderName == "" {
		return nil
	}
	return []byte(s.LeaderName)
}

//
//
//
//
//
//

func LogAndReturnError(code codes.Code, format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args)
	log.Printf(msg)
	return status.New(code, msg).Err()
}

// NewServer generates a new Server using the provided Storage as a Storage backend.
func NewServer(hostname string, id, port int, headlessServer string, store storage.Storage) *Server {
	if hostname == "" || store == nil {
		panic("Hostname not found.")
	}
	addrs, _ := net.LookupIP(hostname)
	var ipv4 string
	for _, addr := range addrs {
		if ip := addr.To4(); ip != nil {
			ipv4 = fmt.Sprintf("%s:%d", ip.String(), port)
		}
	}
	return &Server{Name: hostname, Id: id, Round: 0, Addr: ipv4, storage: store, headlessService: headlessServer}
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
