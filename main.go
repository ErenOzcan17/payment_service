package main

import (
	"context"
	"errors"
	"log"
	"net"
	"time"

	pb "grpc/proto"

	"google.golang.org/grpc"
)

// PaymentServiceServer implements the PaymentService gRPC server
type PaymentServiceServer struct {
	pb.UnimplementedPaymentServiceServer
}

// validateCreditCard validates the credit card number and expiration date
func validateCreditCard(cardNumber string, expiryDate string) error {
	// Check if the card number is 12 digits
	if len(cardNumber) != 16 {
		return errors.New("invalid credit card number: must be 16 digits")
	}

	// Parse the expiration date in MM/YY format
	expiry, err := time.Parse("01/06", expiryDate)
	if err != nil {
		return errors.New("invalid expiration date format: must be MM/YY")
	}

	// Check if the expiration date is in the past
	now := time.Now()
	if expiry.Before(time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)) {
		return errors.New("credit card has expired")
	}

	return nil
}

// ProcessPayment handles payment processing logic
func (s *PaymentServiceServer) ProcessPayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	log.Printf("Processing payment for user: %s, amount: %.2f, method: %s", req.GetUserId(), req.GetAmount(), req.GetPaymentMethod())

	// Validate amount
	if req.GetAmount() <= 0 {
		return &pb.PaymentResponse{
			Success: false,
			Message: "Invalid payment amount",
		}, nil
	}

	// Validate credit card details
	err := validateCreditCard(req.GetCardNumber(), req.GetCardExpiry())
	if err != nil {
		return &pb.PaymentResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	// Payment processing logic
	return &pb.PaymentResponse{
		Success: true,
		Message: "Payment processed successfully",
	}, nil
}

func main() {
	// Create a new gRPC server
	server := grpc.NewServer()
	pb.RegisterPaymentServiceServer(server, &PaymentServiceServer{})

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	log.Println("Payment service is running on port 50051")

	// Start the gRPC server
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
