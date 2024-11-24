package server

import (
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/ayorinde-codes/real-time-delivery-tracking/proto/order"
	"github.com/ayorinde-codes/real-time-delivery-tracking/proto/tracking"
	"github.com/ayorinde-codes/real-time-delivery-tracking/proto/user"
)

func RegisterAllServices(grpcServer *grpc.Server, db *gorm.DB) {
	order.RegisterOrderServiceServer(grpcServer, &OrderService{DB: db})
	user.RegisterUserServiceServer(grpcServer, &UserServiceServer{DB: db})
	tracking.RegisterTrackingServiceServer(grpcServer, &TrackingServiceServer{DB: db})
}
