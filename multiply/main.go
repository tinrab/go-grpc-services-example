package main

import (
	"log"
	"net"

	"github.com/tinrab/go-grpc-services-example/pb"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	rpcServer := grpc.NewServer()
	pb.RegisterMultiplicationServiceServer(rpcServer, &server{})
	reflection.Register(rpcServer)

	if err := rpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to server %v", err)
	}
}

func (s *server) Multiply(ctx context.Context, r *pb.MultiplyRequest) (*pb.MultiplyResponse, error) {
	return &pb.MultiplyResponse{Result: r.A * r.B}, nil
}
