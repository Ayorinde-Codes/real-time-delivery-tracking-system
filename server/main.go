package main

import (
	"log"
	"net"

	"github.com/ayorinde-codes/real-time-delivery-tracking/config"
	"github.com/ayorinde-codes/real-time-delivery-tracking/db"
	"github.com/ayorinde-codes/real-time-delivery-tracking/proto/order"
	"github.com/ayorinde-codes/real-time-delivery-tracking/proto/tracking"
	"github.com/ayorinde-codes/real-time-delivery-tracking/proto/user"
	"github.com/ayorinde-codes/real-time-delivery-tracking/server"
	"google.golang.org/grpc"
)

func main() {
	// Load configuration

	cfg := config.LoadConfig()

	// connect to the database
	database, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// set up gRPC server
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	grpcServer := grpc.NewServer()

	//Register services

	user.RegisterUserServiceServer(grpcServer, &server.UserServiceServer{DB: database})
	order.RegisterOrderServiceServer(grpcServer, &server.OrderServiceServer{DB: database})
	tracking.RegisterTrackingServiceServer(grpcServer, &server.TrackingServiceServer{DB: database})

	log.Println("gRPC server running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
