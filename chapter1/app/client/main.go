package main

import (
	"context"
	"log"

	"sync"

	pb "github.com/rongfengliang/grpcdemo/chapter1/rpc"
	"google.golang.org/grpc"
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
	c := pb.NewUserLoginClient(conn)
	r, err := c.AppLogin(context.Background(), &pb.UserLoginRequest{
		Name:     "dalongdemo",
		Password: "dalongdemo",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Token)
	messageclient, err := c.LongMessage(context.Background(), &pb.QueryRequst{
		Messageversion: "v1",
		Userid:         "60667",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	go func() {
		mess, err := messageclient.Recv()
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		messageclient.SendMsg(
			&pb.QueryRequst{
				Userid:         "dddddddd",
				Messageversion: "v1",
			},
		)
		log.Println(mess.Message)
	}()

	wg.Wait()
}
