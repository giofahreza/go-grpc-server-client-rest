package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go-grpc-client/pb"

	"google.golang.org/grpc"
)

func greet(name string) (result string, error error) {
	// Set up a connection to the server
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
		return "", err
	}
	defer conn.Close()
	// Create a new client
	client := pb.NewGreetServiceClient(conn)
	// Set a timeout for the request
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// Create a request
	req := &pb.GreetRequest{Name: name}
	// Call the Greet method
	res, err := client.Greet(ctx, req)
	if err != nil {
		log.Fatalf("Error calling Greet: %v", err)
		return "", err
	}

	return res.GetResult(), nil
}

func main() {
	result, err := greet("Budi")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Print the response
	fmt.Printf("Response from server: %s\n", result)
}
