package server

import (
	"context"
	"fmt"

	pb "go-grpc-server/pb"
)

type Server struct {
	pb.UnimplementedGreetServiceServer
}

func (s *Server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	firstName := req.GetGreeting()
	result := "Hello " + firstName
	res := &pb.GreetResponse{
		Result: result,
	}
	return res, nil
}
