package main

import (
	"fmt"
	"log"
	"net"

	pb "go-grpc-server/pb"
	"go-grpc-server/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func greet() {
	fmt.Println("Hello World")
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterGreetServiceServer(grpcServer, &server.Server{})

	reflection.Register(grpcServer)

	log.Println("Server is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
