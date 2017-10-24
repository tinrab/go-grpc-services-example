package main

import (
	"crypto/rand"
	"log"
	"net"
	"time"

	pb "github.com/tinrab/go-grpc-services-example/session"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//go:generate protoc -I ../session --go_out=plugins=grpc:../session ../session/session.proto

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSessionServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to server %v", err)
	}
}

func (s *server) GetSession(ctx context.Context, r *pb.SessionRequest) (*pb.SessionResponse, error) {
	return &pb.SessionResponse{
		Token:      "...",
		CreatedAt:  0,
		Expiration: 0,
	}, nil
}

func (s *server) SignIn(ctx context.Context, r *pb.SignInRequest) (*pb.SignInResponse, error) {
	return &pb.SignInResponse{
		Token:      generateToken(),
		Expiration: uint64(time.Now().Add(1 * time.Hour).Unix()),
	}, nil
}

func generateToken() string {
	var (
		alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-")
		n        = 64
	)
	data := make([]byte, n)
	rand.Read(data)
	token := make([]rune, n)
	for i := range data {
		token[i] = alphabet[int(data[i])%len(alphabet)]
	}
	return string(token)
}
