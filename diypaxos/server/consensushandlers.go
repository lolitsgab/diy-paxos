package server

//
//import (
//	"context"
//	pb "diy-paxos/diypaxos/proto"
//	"google.golang.org/grpc"
//	"log"
//	"math"
//	"os"
//	"time"
//)
//
//var logger *log.Logger = log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds)
//
//type Candidate struct {
//	round  float64
//	leader string
//}
//
//func (s *Server) ElectLeader() {
//	for {
//		proposed := s.majorityPromises()
//		if s.majorityAccept(proposed) {
//			s.leaderMu.Lock()
//			if s.LeaderName == "" {
//				s.LeaderName = proposed.leader
//			}
//			s.leaderMu.Unlock()
//			break
//		}
//	}
//	logger.Printf("Elected leader: %v", s.LeaderName)
//}
//
//func (s *Server) majorityAccept(candidate *Candidate) bool {
//	logger.Printf("=================================")
//	logger.Printf("=================================")
//	majority := int((len(s.Replicas) + 1) / 2)
//	for _, addr := range s.Replicas {
//		s.leaderMu.Lock()
//		if s.LeaderName != "" {
//			s.leaderMu.Unlock()
//			return true
//		}
//		s.leaderMu.Unlock()
//		conn, err := grpc.Dial(addr, grpc.WithInsecure())
//		if err != nil {
//			logger.Printf("Failed to dial %v: %v", addr, err)
//			continue
//		}
//		defer conn.Close()
//		client := pb.NewSimpleKvStoreClient(conn)
//		req := &pb.AcceptRequest{Name: s.Name, Round: s.Round, Key: "leader", Val: []byte(candidate.leader)}
//		logger.Printf("sending accept %v to %v", req, addr)
//		resp, err := client.Accept(context.Background(), req)
//		logger.Printf("received %v from %v", resp, addr)
//		if err != nil {
//			continue
//		}
//		if !resp.GetAccepted() {
//			continue
//		}
//		majority -= 1
//		if majority <= 0 {
//			break
//		}
//	}
//	return majority <= 0
//}
//
//func (s *Server) majorityPromises() *Candidate {
//	var candidate *Candidate
//	majorityReached := false
//	postfix := float64(s.Id) / 10
//
//	for !majorityReached {
//		logger.Printf("=================================")
//		logger.Printf("=================================")
//		sleep(s.Id)
//		s.roundMu.Lock()
//		s.Round = math.Trunc(s.Round) + postfix
//
//		candidate = &Candidate{
//			round:  s.Round,
//			leader: s.Name,
//		}
//
//		majority := int((len(s.Replicas) + 1) / 2)
//		for _, addr := range s.Replicas {
//			s.leaderMu.Lock()
//			if s.LeaderName != "" {
//				candidate.leader = s.LeaderName
//				return candidate
//			}
//			s.leaderMu.Unlock()
//			conn, err := grpc.Dial(addr, grpc.WithInsecure())
//			if err != nil {
//				logger.Printf("Failed to dial %v: %v", addr, err)
//				continue
//			}
//			defer conn.Close()
//			client := pb.NewSimpleKvStoreClient(conn)
//			req := &pb.PrepareRequest{Name: s.Name, Round: s.Round}
//			logger.Printf("sending prepare %v to %v", req, addr)
//			resp, err := client.Prepare(context.Background(), req)
//			logger.Printf("received %v from %v", resp, addr)
//			if err != nil {
//				continue
//			}
//			if !resp.GetPromise() {
//				s.Round = math.Trunc(resp.GetHighestRoundSeen()) + 1
//				break
//			}
//			if candidate.round < resp.GetHighestRoundSeen() {
//				candidate.round = resp.GetHighestRoundSeen()
//				candidate.leader = string(resp.GetVal())
//			}
//			majority -= 1
//			if majority <= 0 {
//				majorityReached = true
//				break
//			}
//		}
//		s.roundMu.Unlock()
//	}
//	return candidate
//}
//
//func (s *Server) Prepare(ctx context.Context, in *pb.PrepareRequest) (*pb.PrepareResponse, error) {
//	s.leaderMu.Lock()
//	defer s.leaderMu.Unlock()
//	s.roundMu.Lock()
//	defer s.roundMu.Unlock()
//
//	s.Round = math.Max(s.Round, in.Round)
//	return &pb.PrepareResponse{
//		Name:             s.Name,
//		Promise:          s.Round <= in.Round,
//		HighestRoundSeen: s.Round,
//		Val:              s.getLeaderVal(),
//	}, nil
//}
//
//func (s *Server) Accept(ctx context.Context, in *pb.AcceptRequest) (*pb.AcceptResponse, error) {
//	s.leaderMu.Lock()
//	defer s.leaderMu.Unlock()
//	s.roundMu.Lock()
//	defer s.roundMu.Unlock()
//
//	if in.GetRound() >= s.Round {
//		s.LeaderName = string(in.GetVal())
//		s.Round = in.GetRound()
//	}
//	return &pb.AcceptResponse{
//		Name:     s.Name,
//		Accepted: in.GetRound() >= s.Round,
//	}, nil
//}
//
//func (s *Server) getLeaderVal() []byte {
//	if s.LeaderName == "" {
//		return nil
//	}
//	return []byte(s.LeaderName)
//}
//
//func sleep(n int) {
//	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
//	//n := r.Intn(10)
//	n = n * 2
//	logger.Printf("Sleeping %d seconds...\n", n)
//	time.Sleep(time.Duration(n) * time.Second)
//	logger.Println("Done")
//}
