package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"go-grpc-client/helpers"
	"go-grpc-client/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(helpers.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(helpers.StreamClientInterceptor()),
	)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMyServiceClient(conn)

	// Unary
	// 1 request, 1 response
	res, err := client.GetUser(context.Background(), &pb.GetUserRequest{Id: 1})
	if err != nil {
		log.Fatalf("Error calling GetUser: %v", err)
	}
	fmt.Println("Unary GetUser:", res.Username)

	// Server streaming
	// 1 request, multiple response
	fmt.Println("Server streaming ListUsers:")
	stream1, err := client.ListUsers(context.Background(), &pb.ListUsersRequest{})
	if err != nil {
		log.Fatalf("Error calling ListUsers: %v", err)
	}
	for {
		user, err := stream1.Recv()
		if err == io.EOF {
			break
		}
		fmt.Println("  ", user.Username)
	}

	// Client streaming
	// Multiple request, 1 response
	fmt.Println("Client streaming UploadLogs:")
	stream2, err := client.UploadLogs(context.Background())
	if err != nil {
		log.Fatalf("Error calling UploadLogs: %v", err)
	}
	messages := []string{"log1", "log2", "log3"}
	for _, msg := range messages {
		stream2.Send(&pb.Log{Message: msg})
	}
	res2, err := stream2.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}
	fmt.Println("Uploaded logs:", res2.Count)

	// Bidirectional streaming
	// multiple request, multiple response
	fmt.Println("Bidirectional streaming Chat:")
	stream3, err := client.Chat(context.Background())
	if err != nil {
		log.Fatalf("Error calling Chat: %v", err)
	}

	go func() {
		messages := []string{"Hello", "How are you?", "Bye"}
		for _, msg := range messages {
			stream3.Send(&pb.ChatMessage{Text: msg})
			time.Sleep(time.Millisecond * 1000)
		}
		stream3.CloseSend()
	}()

	for {
		res, err := stream3.Recv()
		if err == io.EOF {
			break
		}
		fmt.Println("  Received:", res.Text)
	}
}
