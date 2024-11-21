package main

import (
	"log"
	"net"

	pb "github.com/ayorinde-codes/real-time-delivery-tracking/proto/user"
	"github.com/ayorinde-codes/real-time-delivery-tracking/server"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listener("tcp", ":50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userService := &server.UserServiceServer{}

	pb.RegisterUserServiceServer(grpcServer, userService)

	log.Println("Server is running on port 50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
