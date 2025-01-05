package main

import (
	"context"
	"log"
	"net"

	pb "grpc/proto"

	"google.golang.org/grpc"
)

// PaymentServiceServer implements the PaymentService gRPC server
type PaymentServiceServer struct {
	pb.UnimplementedPaymentServiceServer
}

// ProcessPayment handles payment processing logic
func (s *PaymentServiceServer) ProcessPayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	log.Printf("Processing payment for user: %s, amount: %.2f, method: %s", req.GetUserId(), req.GetAmount(), req.GetPaymentMethod())

	if req.GetAmount() > 0 {
		return &pb.PaymentResponse{
			Success: true,
			Message: "Payment processed successfully",
		}, nil
	}

	return &pb.PaymentResponse{
		Success: false,
		Message: "Invalid payment amount",
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
