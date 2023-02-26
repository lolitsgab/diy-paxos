package server

import (
	"context"
	pb "diy-paxos/diypaxos/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
)

// first draft of the interceptor
func TeeInterceptor(s *Server) grpc.UnaryServerInterceptor {
	// Define the interceptor function
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, _ := metadata.FromIncomingContext(ctx)
		log.Printf("%v", md)
		id := md.Get("id")
		for _, addr := range s.Replicas {

			if id != nil && id[0] == addr {
				continue
			}
			conn, err := grpc.Dial(addr, grpc.WithInsecure())
			if err != nil {
				log.Printf("Failed to dial %v: %v", addr, err)
				return nil, nil
			}
			defer conn.Close()
			client := pb.NewSimpleKvStoreClient(conn)
			log.Printf("Teeing request to %v: %v", addr, req)

			// Call the corresponding method on the client
			newCtx := metadata.AppendToOutgoingContext(ctx, "id", s.Addr)
			_, err = client.Get(newCtx, req.(*pb.GetRequest))
			if err != nil {
				log.Printf("Failed to tee request to %v: %v", addr, err)
			}
		}
		return handler(ctx, req)
	}
}
