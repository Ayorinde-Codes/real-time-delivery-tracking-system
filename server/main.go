package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/ayorinde-codes/real-time-delivery-tracking/db"
	"github.com/ayorinde-codes/real-time-delivery-tracking/proto/order"
	"github.com/ayorinde-codes/real-time-delivery-tracking/proto/tracking"
	"github.com/ayorinde-codes/real-time-delivery-tracking/proto/user"
	"github.com/ayorinde-codes/real-time-delivery-tracking/server"
)

func main() {
	// Connect to the database
	db.ConnectDB()
	defer func() {
		sqlDB, err := db.DB.DB()
		if err == nil {
			sqlDB.Close()
		}
	}()

	// Get the server port from environment variables
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080" // Default to 8080
	}

	// Start listening on the specified port
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register gRPC services
	order.RegisterOrderServiceServer(grpcServer, &server.OrderServiceServer{DB: db.DB})
	user.RegisterUserServiceServer(grpcServer, &server.UserServiceServer{DB: db.DB})
	tracking.RegisterTrackingServiceServer(grpcServer, &server.TrackingServiceServer{DB: db.DB})

	log.Printf("Server is running on port %s...", port)

	// Start serving requests
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
