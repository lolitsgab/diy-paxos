package server

import (
	"context"
	pb "diy-paxos/diypaxos/proto"
	"fmt"
	"google.golang.org/grpc"
	"math"
	"math/rand"
	"time"
)

type Candidate struct {
	round  float64
	leader string
}

func (s *Server) ElectLeader() {
	for {
		sleepIdSeconds(s.Id)
		if s.isLeaderElected() {
			break
		}
		candidate, err := proposeLeader(s)
		if err != nil {
			logger.Printf("unable to get proposal consensus: %v", err.Error())
			continue
		}
		updateLeaderFromCandidate(s, candidate)
		err = getAcceptFromMajority(s, candidate)
		if err == nil {
			break
		}
		logger.Printf("unable to get leader accepted: %v", err.Error())
	}
}

func (s *Server) isLeaderElected() bool {
	s.leaderMu.Lock()
	defer s.leaderMu.Unlock()
	return s.LeaderName != ""
}

func (s *Server) isLeader() bool {
	s.leaderMu.Lock()
	defer s.leaderMu.Unlock()
	return s.LeaderName == s.Addr
}

func updateLeaderFromCandidate(s *Server, candidate *Candidate) {
	s.leaderMu.Lock()
	defer s.leaderMu.Unlock()
	if s.LeaderName != "" {
		candidate.leader = s.LeaderName
	}
	s.LeaderName = candidate.leader
}

func calculateMajority(replicaCount int) int {
	return int(math.Ceil(float64(replicaCount+1) / 2))
}

func getAcceptFromMajority(s *Server, candidate *Candidate) error {
	majority := calculateMajority(len(s.Replicas))
	for _, addr := range s.Replicas {
		if sendAccept(addr, candidate, s, majority) {
			majority--
		}
	}
	if majority > 0 {
		return fmt.Errorf("no majority accepted")
	}
	return nil
}

func sendAccept(addr string, candidate *Candidate, s *Server, majority int) bool {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		logger.Printf("Failed to dial %v: %v", addr, err)
		return false
	}
	client := pb.NewSimpleKvStoreClient(conn)
	req := &pb.AcceptRequest{Name: s.Name, Round: s.Round, Key: "leader", Val: []byte(candidate.leader)}
	logger.Printf("sending accept %v to %v", req, addr)
	s.roundMu.Lock()
	resp, err := client.Accept(context.Background(), req)
	if err == nil && resp.GetHighestRoundSeen() > s.Round {
		s.Round = resp.GetHighestRoundSeen()
	}
	s.roundMu.Unlock()
	return err == nil && resp.GetAccepted()
}

func proposeLeader(s *Server) (*Candidate, error) {
	postfix := float64(s.Id) / 10
	s.roundMu.Lock()
	s.Round = math.Max(s.Round, math.Trunc(s.Round+1)+postfix)
	candidate := &Candidate{
		round:  s.Round,
		leader: s.Addr,
	}
	s.roundMu.Unlock()

	promises := map[string]int{}
	promises[candidate.leader]++
	for _, addr := range s.Replicas {
		sendProposal(s, addr, candidate, promises)
	}
	if err := setMajorityLeader(promises, candidate, len(s.Replicas)); err != nil {
		logger.Printf("err %v", err.Error())
		return nil, err
	}
	return candidate, nil
}

func sendProposal(s *Server, addr string, candidate *Candidate, promises map[string]int) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		logger.Printf("Failed to dial %v: %v", addr, err)
		return
	}
	client := pb.NewSimpleKvStoreClient(conn)
	req := &pb.PrepareRequest{Name: s.Name, Round: s.Round}
	logger.Printf("sending prepare %v to %v", req, addr)
	resp, err := client.Prepare(context.Background(), req)
	if err != nil {
		logger.Printf("unable to contact %v: %v", addr, err.Error())
		return
	}
	logger.Printf("received %v from %v", resp, resp.Name)
	if resp.Promise {
		if resp.Val == nil {
			promises[candidate.leader]++
		} else {
			promises[string(resp.Val)]++
		}
	} else {
		s.roundMu.Lock()
		s.Round = math.Max(s.Round, resp.HighestRoundSeen)
		s.roundMu.Unlock()
	}
}

func setMajorityLeader(m map[string]int, candidate *Candidate, majority int) error {
	max := -1
	for k, v := range m {
		if max < v {
			max = v
			candidate.leader = k
		}
	}
	if max < majority {
		return fmt.Errorf("no majority. Needed %v got %v, %v", majority, max, m)
	}
	return nil
}

func randSleep() {
	max := 5000
	min := 1000
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(max-min) + min
	logger.Printf("Sleeping %d milliseconds...\n", n)
	time.Sleep(time.Duration(n) * time.Millisecond)
	logger.Println("Done")
}

func sleepIdSeconds(id int) {
	duration := id * 1000
	logger.Printf("Sleeping %d milliseconds...\n", duration)
	time.Sleep(time.Duration(duration) * time.Millisecond)
	logger.Println("Woke up.")
}
