package server

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/ayorinde-codes/real-time-delivery-tracking/github.com/ayorinde-codes/real-time-delivery-tracking/proto/tracking"
)

type TrackingServiceServer struct {
	tracking.UnimplementedTrackingServiceServer
	DB *sql.DB
}

// SendLocationStream receives a stream of locations for a given order
func (s *TrackingServiceServer) SendLocationStream(stream tracking.TrackingService_SendLocationStreamServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}

		_, err = s.DB.Exec("INSERT INTO tracking (order_id, latitude, longitude) VALUES ($1, $2, $3)",
			req.OrderId, req.Latitude, req.Longitude)
		if err != nil {
			log.Printf("Error saving location: %v", err)
			return errors.New("Failed to save location")
		}
	}
}

// TrackOrder retrieves the latest location of an order
func (s *TrackingServiceServer) TrackOrder(ctx context.Context, req *tracking.TrackOrderRequest) (*tracking.TrackOrderResponse, error) {
	var latitude, longitude float64
	err := s.DB.QueryRow("SELECT latitude, longitude FROM tracking WHERE order_id = $1 ORDER BY updated_at DESC LIMIT 1", req.OrderId).
		Scan(&latitude, &longitude)
	if err != nil {
		log.Printf("Error tracking order: %v", err)
		return nil, errors.New("Failed to track order")
	}
	return &tracking.TrackOrderResponse{
		OrderId: req.OrderId,
		Status:  "Order is in transit",
	}, nil
}

func (s *TrackingServiceServer) SubscribeLocationUpdates(req *tracking.LocationRequest, stream tracking.TrackingService_SubscribeLocationUpdatesServer) error {
	rows, err := s.DB.Query("SELECT latitude, longitude, updated_at FROM tracking WHERE order_id = $1 ORDER BY updated_at DESC", req.OrderId)
	if err != nil {
		log.Printf("Error retrieving location updates: %v", err)
		return errors.New("failed to retrieve location updates")
	}
	defer rows.Close()

	for rows.Next() {
		var latitude, longitude float64
		var updatedAt time.Time

		if err := rows.Scan(&latitude, &longitude, &updatedAt); err != nil {
			return err
		}

		// Send location response stream
		err = stream.Send(&tracking.LocationResponse{
			OrderId:   req.OrderId,
			Latitude:  latitude,
			Longitude: longitude,
			Message:   "Location update",
		})
		if err != nil {
			return err
		}
	}

	return nil
}
