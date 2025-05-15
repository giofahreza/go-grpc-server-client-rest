package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"go-grpc-server/helpers"
	"go-grpc-server/pb"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMyServiceServer
}

// Unary RPC
// 1 request, 1 response
func (*server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return &pb.GetUserResponse{Username: fmt.Sprintf("user-%d", req.Id)}, nil
}

// Server streaming RPC
// 1 request, multiple response
func (*server) ListUsers(req *pb.ListUsersRequest, stream pb.MyService_ListUsersServer) error {
	for i := 1; i <= 7; i++ {
		stream.Send(&pb.User{Username: fmt.Sprintf("user%d", i)})
		time.Sleep(time.Second * 1)
	}
	return nil
}

// Client streaming RPC
// Multiple request, 1 response
func (*server) UploadLogs(stream pb.MyService_UploadLogsServer) error {
	count := 0
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.UploadSummary{Count: int32(count)})
		}
		if err != nil {
			return err
		}
		count++
	}
}

// Bidirectional streaming RPC
// multiple request, multiple response
func (*server) Chat(stream pb.MyService_ChatServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		response := &pb.ChatMessage{Text: "Echo: " + msg.Text}
		if err := stream.Send(response); err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(helpers.UnaryServerInterceptor()),
		grpc.StreamInterceptor(helpers.StreamServerInterceptor()),
	)
	pb.RegisterMyServiceServer(s, &server{})
	log.Println("gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
