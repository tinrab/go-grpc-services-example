package main

import (
	"context"
	"log"

	pb "github.com/tinrab/go-grpc-services-example/session"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("session:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	client := pb.NewSessionClient(conn)

	res, err := client.SignIn(context.Background(), &pb.SignInRequest{User: "tin"})
	if err != nil {
		log.Fatalf("Sign in failed: %v", err)
	}

	log.Println(res.Token)
}
