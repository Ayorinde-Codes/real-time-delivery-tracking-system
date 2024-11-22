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

// CreateOrder creates a new order
func (s *OrderServiceServer) CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	// Create a new order
	_, err := s.DB.Exec("INSERT INTO orders (customer_id) VALUES ($1)", req.CustomerId)

	if err != nil {
		log.Printf("Error creating order: %v", err)
		return nil, errors.New("Failed to create order")
	}
	return &order.CreateOrderResponse{OrderId: 1}, nil
}

// UpdateOrderStatus updates the status of an order
func (s *OrderServiceServer) UpdateOrderStatus(ctx context.Context, req *order.UpdateOrderStatusRequest) (*order.UpdateOrderStatusResponse, error) {

	// Update the status of the order
	_, err := s.DB.Exec("UPDATE orders SET status = $1, updated_at = NOW() WHERE id = $2", req.Status, req.OrderId)
	if err != nil {
		log.Printf("Error updating order status: %v", err)
		return nil, errors.New("Failed to update order status")
	}

	return &order.UpdateOrderStatusResponse{Message: "Order status updated successfully!"}, nil
}

// GetOrderStatus retrieves the status of an order
func (s *OrderServiceServer) GetOrderStatus(ctx context.Context, req *order.GetOrderStatusRequest) (*order.GetOrderStatusResponse, error) {

	var status string
	err := s.DB.QueryRow("SELECT status FROM orders where id = $1", req.OrderId).Scan(&status)
	if err != nil {
		log.Printf("Error retrieving order status: %v", err)
		return nil, errors.New("Failed to get order status")
	}
	return &order.GetOrderStatusResponse{Status: status}, nil
}
