# Real-Time Delivery Tracking Backend
This repository provides a real-time tracking system for delivery services, built using Go, gRPC, and PostgreSQL. The service allows users to stream, track, and subscribe to live location updates for orders.

## Overview
A backend service for real-time delivery tracking built using GRPC.

## Getting Started


### Tech Stack
- **Go**: Main programming language.
- **gRPC**: For inter-service communication.
- **PostgreSQL**: Database.
- **Protobuf Compiler**: Install `protoc` for generating Go code from proto files.

## Features
- **Location Streaming**: Send and save live location updates for an order.
- **Order Tracking**: Retrieve the latest status of an order.
- **Location Subscriptions**: Subscribe to real-time updates for a specific order.
- **Authentication**: Secured endpoints using JWT.
- **Database Management**: GORM-powered PostgreSQL integration.


### Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/ayorinde-codes/real-time-delivery-tracking.git
   cd real-time-delivery-tracking
   ```
2. ### Install dependencies:
    ```bash
    go mod tidy
    ```


3. Start the services with Docker Compose:
    ```bash
    docker-compose up
    ```

3. Endpoints
- **OrderService**: Create and update orders.
- **UserService**: Manage users.
- **TrackingService**: Add and get tracking data.
