package main

import (
	"fmt"
	"log"
	"net"

	"go-grpc-server/pb"
	"go-grpc-server/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the GreetService with the gRPC server
	pb.RegisterGreetServiceServer(grpcServer, &service.GreetService{})

	// Register reflection service on gRPC server
	reflection.Register(grpcServer)

	// Start listening for incoming connections
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Println("Server is running on port :50051...")

	// Serve gRPC server
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
