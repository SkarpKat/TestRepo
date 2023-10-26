package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/SkarpKat/TestRepo/"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type ChatServiceServer struct {
	pb.UnimplementedChatServiceServer
}

func (s *ChatServiceServer) Join(ctx context.Context, in *pb.JoinRequest) (*pb.ChatMessage, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.ChatMessage{Body: "Welcome " + in.GetName()}, nil
}

func (s *ChatServiceServer) Say(ctx context.Context, in *pb.ChatMessage) (*pb.ChatMessage, error) {
	log.Printf("Received: %v", in.GetBody())
	return &pb.ChatMessage{Body: in.GetBody()}, nil
}

func (s *ChatServiceServer) Leave(ctx context.Context, in *pb.LeaveRequest) (*pb.ChatMessage, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.ChatMessage{Body: "Goodbye " + in.GetName()}, nil
}

func main() {
	fmt.Println("Starting server...")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Println("Listening on port " + port)

	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &ChatServiceServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
