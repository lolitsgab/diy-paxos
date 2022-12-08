package server

import (
	"context"
	pb "diy-paxos/diypaxos/proto"
)

func (s *Server) Promise(ctx context.Context, in *pb.PromiseRequest) (*pb.PromiseResponse, error) {
	return nil, nil
}

func (s *Server) Accept(ctx context.Context, in *pb.AcceptRequest) (*pb.AcceptResponse, error) {
	return nil, nil
}
