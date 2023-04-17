package server

import (
	"context"
	pb "diy-paxos/diypaxos/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
)

// TeeInterceptor returns a new gRPC unary server interceptor that logs information about incoming requests and forwards
// the requests to all replicas except the one that received the original request
func TeeInterceptor(s *Server) grpc.UnaryServerInterceptor {
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
			_, err = dispatchRequest(newCtx, req, client)
			if err != nil {
				log.Printf("Failed to tee request to %v: %v", addr, err)
			}
		}
		return handler(ctx, req)
	}
}

// dispatchRequest calls the corresponding handler for the incoming request.
func dispatchRequest(ctx context.Context, req interface{}, client pb.SimpleKvStoreClient) (interface{}, error) {
	switch r := req.(type) {
	case *pb.GetRequest:
		return client.Get(ctx, r)
	case *pb.InsertRequest:
		return client.Insert(ctx, r)
	case *pb.RemoveRequest:
		return client.Remove(ctx, r)
	case *pb.UpdateRequest:
		return client.Update(ctx, r)
	case *pb.UpsertRequest:
		return client.Upsert(ctx, r)
	case *pb.AcceptRequest:
		return client.Accept(ctx, r)
	case *pb.PrepareRequest:
		return client.Prepare(ctx, r)
	default:
		return nil, status.Errorf(codes.Unimplemented, "method not implemented")
	}
}
