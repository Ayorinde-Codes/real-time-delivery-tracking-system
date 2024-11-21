package server

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/ayorinde-codes/real-time-delivery-tracking/proto/order"
)

type OrderServiceServer struct {
	order.UnimplementedOrderServiceServer
	DB *sql.DB
}

func (s *OrderServiceServer) CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	// Create a new order
	_, err := s.DB.Exec("INSERT INTO orders (customer_id) VALUES ($1)", req.CustomerId)

	if err != nil {
		log.Printf("Error creating order: %v", err)
		return nil, errors.NewError("Failed to create order")
	}
	return &order.CreateOrderResponse{OrderId: 1}, nil
}
