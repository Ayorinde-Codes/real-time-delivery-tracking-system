package server

import (
	"context"
	"errors"
	"io"
	"log"

	"github.com/ayorinde-codes/real-time-delivery-tracking/models"
	"github.com/ayorinde-codes/real-time-delivery-tracking/proto/tracking"
	"gorm.io/gorm"
)

type TrackingServiceServer struct {
	tracking.UnimplementedTrackingServiceServer
	DB *gorm.DB
}

// SendLocationStream receives a stream of locations for a given order
func (s *TrackingServiceServer) SendLocationStream(stream tracking.TrackingService_SendLocationStreamServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&tracking.LocationResponse{
				Message: "Location updates saved successfully",
			})
		}
		if err != nil {
			log.Printf("Error receiving stream: %v", err)
			return err
		}

		locationUpdate := models.Tracking{
			OrderID:   req.OrderId,
			Latitude:  req.Latitude,
			Longitude: req.Longitude,
		}

		if result := s.DB.Create(&locationUpdate); result.Error != nil {
			log.Printf("Error saving location update: %v", result.Error)
			return errors.New("failed to save location update")
		}

		log.Printf("Saved location update: %+v", locationUpdate)
	}
}

// TrackOrder retrieves the latest location of an order
func (s *TrackingServiceServer) TrackOrder(ctx context.Context, req *tracking.TrackOrderRequest) (*tracking.TrackOrderResponse, error) {
	var location models.Tracking

	err := s.DB.Where("order_id = ?", req.OrderId).
		Order("updated_at DESC").
		First(&location).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no tracking information found for this order")
		}
		log.Printf("Error tracking order: %v", err)
		return nil, errors.New("failed to track order")
	}

	return &tracking.TrackOrderResponse{
		OrderId: req.OrderId,
		Status:  "Order is in transit",
	}, nil
}

// SubscribeLocationUpdates streams location updates for a given order
func (s *TrackingServiceServer) SubscribeLocationUpdates(req *tracking.LocationRequest, stream tracking.TrackingService_SubscribeLocationUpdatesServer) error {
	var locationUpdates []models.Tracking

	if result := s.DB.Where("order_id = ?", req.OrderId).Order("created_at ASC").Find(&locationUpdates); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("No location updates found for OrderID %d", req.OrderId)
			return errors.New("no location updates found")
		}
		log.Printf("Error retrieving location updates: %v", result.Error)
		return errors.New("failed to retrieve location updates")
	}

	for _, update := range locationUpdates {
		err := stream.Send(&tracking.LocationResponse{
			OrderId:   update.OrderID,
			Latitude:  update.Latitude,
			Longitude: update.Longitude,
			Message:   "Location update streamed",
		})
		if err != nil {
			log.Printf("Error streaming location update: %v", err)
			return err
		}
	}

	return nil
}
