package server

import (
	"context"
	pb "diy-paxos/diypaxos/proto"
	"diy-paxos/diypaxos/storage"
	"google.golang.org/grpc/codes"
	"log"
)

func (s *Server) Promise(ctx context.Context, in *pb.PromiseRequest) (*pb.PromiseResponse, error) {
	if val, seen := s.promises[in.GetKey()]; seen && val.Meta.Version >= in.GetVersion() {
		return &pb.PromiseResponse{Name: s.name, Promise: false}, nil
	}
	s.promises[in.GetKey()] = storage.Value{Value: in.GetValue(), Meta: storage.Metadata{Version: in.GetVersion()}}
	return &pb.PromiseResponse{Name: s.name, Promise: true}, nil
}

func (s *Server) Accept(ctx context.Context, in *pb.AcceptRequest) (*pb.AcceptResponse, error) {
	if val, ok := s.promises[in.GetKey()]; !ok {
		return nil, LogAndReturnError(codes.NotFound, "no promise found for %v: cannot accept before promise", in.GetKey())

	} else if val.Meta.Version > in.GetVersion() {
		// Technically not an error, just not something we can do.
		log.Printf("cannot accept a version lower (%v) than promised version (%v).", val.Meta.Version, in.GetVersion())
		return &pb.AcceptResponse{Name: s.name, Committed: false}, nil
	}

	delete(s.promises, in.GetKey())
	return &pb.AcceptResponse{Name: s.name, Committed: true}, nil
}
