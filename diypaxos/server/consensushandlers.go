package server

import (
	"context"
	pb "diy-paxos/diypaxos/proto"
	"google.golang.org/grpc"
	"log"
	"math"
	"os"
)

var logger *log.Logger = log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds)

type Candidate struct {
	round  float64
	leader string
}

func (s *Server) ElectLeader() {
	for {
		proposed := s.majorityPromises()
		s.mu.Lock()
		if s.LeaderName != "" || s.majorityAccept(proposed) {
			s.LeaderName = proposed.leader
			s.mu.Unlock()
			break
		}
		s.mu.Unlock()
	}
	logger.Printf("Elected leader: %v", s.LeaderName)
}

func (s *Server) majorityAccept(candidate *Candidate) bool {
	majority := int(len(s.Replicas)/2) + 1
	for _, addr := range s.Replicas {
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			logger.Printf("Failed to dial %v: %v", addr, err)
			continue
		}
		defer conn.Close()
		client := pb.NewSimpleKvStoreClient(conn)
		logger.Printf("sending prepare to %v", addr)
		resp, err := client.Accept(context.Background(), &pb.AcceptRequest{Name: s.Name, Round: s.Round, Key: "leader", Val: []byte(candidate.leader)})
		if err != nil {
			continue
		}
		if !resp.GetAccepted() {
			continue
		}
		majority -= 1
		if majority <= 0 {
			break
		}
	}
	return majority <= 0
}

func (s *Server) majorityPromises() *Candidate {
	var candidate *Candidate
	majorityReached := false

	for !majorityReached {
		sleep()
		postfix := float64(s.Id) / 10
		s.Round = math.Trunc(s.Round) + postfix

		candidate = &Candidate{
			round:  s.Round,
			leader: s.Name,
		}

		majority := int((len(s.Replicas)+1)/2) + 1
		for _, addr := range s.Replicas {
			if s.LeaderName != "" {
				candidate.leader = s.LeaderName
				return candidate
			}
			conn, err := grpc.Dial(addr, grpc.WithInsecure())
			if err != nil {
				logger.Printf("Failed to dial %v: %v", addr, err)
				continue
			}
			defer conn.Close()
			client := pb.NewSimpleKvStoreClient(conn)
			logger.Printf("sending prepare to %v", addr)
			resp, err := client.Prepare(context.Background(), &pb.PrepareRequest{Name: s.Name, Round: s.Round})
			if err != nil {
				continue
			}
			if !resp.GetPromise() {
				s.Round = math.Trunc(resp.GetHighestRoundSeen()) + 1
				break
			}
			if candidate.round < resp.GetHighestRoundSeen() {
				candidate.round = resp.GetHighestRoundSeen()
				candidate.leader = string(resp.GetVal())
			}
			majority -= 1
			if majority <= 0 {
				majorityReached = true
				break
			}
		}
	}
	return candidate
}

func sleep() {
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//n := r.Intn(10)
	//logger.Printf("Sleeping %d seconds...\n", n)
	//time.Sleep(time.Duration(n) * time.Second)
	//logger.Println("Done")
}
