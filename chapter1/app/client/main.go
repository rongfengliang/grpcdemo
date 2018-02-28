package main

import (
	"context"
	"log"

	"sync"

	pb "github.com/rongfengliang/grpcdemo/chapter1/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	address = "127.0.0.1:5001"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	md := metadata.Pairs("authorization", "token dalong")
	c := pb.NewUserLoginClient(conn)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := c.AppLogin(ctx, &pb.UserLoginRequest{
		Name:     "dalongdemo",
		Password: "dalongdemo",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Token)

	// test with new context must be fail is pass a new context
	r2, err := c.AppLogin(ctx, &pb.UserLoginRequest{
		Name:     "dalongdemo",
		Password: "dalongdemo",
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	result, err := c.TokenVerify(ctx, &pb.UserTokenRequest{
		Token: r2.Token,
	})
	if result.Isverify {
		log.Println("is ok")
	}
	log.Printf("Greeting: %s", r2.Token)
	client, err := c.LongMessage(ctx, &pb.QueryRequst{
		Userid:         "60667",
		Messageversion: "v1",
	})
	if err != nil {
		log.Fatalf("some wrong with stream")
	}
	mess, _ := client.Recv()
	log.Println(mess.Message)
	wg.Wait()
}
