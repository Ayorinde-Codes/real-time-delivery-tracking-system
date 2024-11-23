package service

import (
	"context"

	"github.com/ayorinde-codes/real-time-delivery-tracking/models"
	"github.com/ayorinde-codes/real-time-delivery-tracking/proto/order"
	"gorm.io/gorm"
)

type OrderService struct {
	db *gorm.DB
	order.UnimplementedOrderServiceServer
}

// CreateOrder creates a new order.
func (s *OrderService) CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	newOrder := models.Order{
		CustomerID: req.CustomerId,
		Status:     "Pending",
	}

	if err := s.db.Create(&newOrder).Error; err != nil {
		return nil, err
	}

	return &order.CreateOrderResponse{
		Message: "Order created successfully",
	}, nil
}

// UpdateOrderStatus updates the status of an order.
func (s *OrderService) UpdateOrderStatus(ctx context.Context, req *order.UpdateOrderStatusRequest) (*order.UpdateOrderStatusResponse, error) {
	var existingOrder models.Order
	if err := s.db.First(&existingOrder, req.OrderId).Error; err != nil {
		return nil, err
	}

	existingOrder.Status = req.Status
	if err := s.db.Save(&existingOrder).Error; err != nil {
		return nil, err
	}

	return &order.UpdateOrderStatusResponse{
		Message: "Order status updated successfully",
	}, nil
}

// GetOrderStatus retrieves the status of an order.
func (s *OrderService) GetOrderStatus(ctx context.Context, req *order.GetOrderStatusRequest) (*order.GetOrderStatusResponse, error) {
	var existingOrder models.Order
	if err := s.db.First(&existingOrder, req.OrderId).Error; err != nil {
		return nil, err
	}

	return &order.GetOrderStatusResponse{
		Status: existingOrder.Status,
	}, nil
}
