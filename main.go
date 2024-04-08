package main

import (
	"blogging/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	strategy := strategy.NewInMemoryBlogStrategy()

	// Create a server instance with the chosen strategy
	srv := server.GetInstance(strategy)

	// Start the gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterBlogServiceServer(grpcServer, srv)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
