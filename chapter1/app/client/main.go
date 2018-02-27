package main

import (
	"context"
	"log"

	pb "github.com/rongfengliang/grpcdemo/chapter1/rpc"
	"google.golang.org/grpc"
)

const (
	address = "127.0.0.1:5001"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserLoginClient(conn)
	r, err := c.AppLogin(context.Background(), &pb.UserLoginRequest{
		Name:     "dalongdemo",
		Password: "dalongdemo",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Token)
}
