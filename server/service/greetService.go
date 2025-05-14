package service

import (
	"context"
	"fmt"

	"go-grpc-server/pb"
)

type GreetService struct {
	pb.UnimplementedGreetServiceServer
}

func (s *GreetService) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	// Extract the name from the request
	name := req.GetName()

	// Create a greeting message
	greeting := fmt.Sprintf("Hello, %s!", name)

	// Create and return the response
	response := &pb.GreetResponse{
		Result: greeting,
	}
	return response, nil
}
