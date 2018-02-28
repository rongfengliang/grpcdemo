package main

import (
	"context"
	"log"

	pb "git.yonyou.com/go/grpcdemo/chapter2"
	"google.golang.org/grpc"
)

const (
	address = "127.0.0.1:9090"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("some wrong", err)
	}
	defer conn.Close()
	c := pb.NewAppListClient(conn)
	applist, err := c.List(context.Background(), &pb.AppRequest{
		Appid: "001",
	})
	if err != nil {
		log.Fatal("some wrong", err)
	}
	if applist.Code != 0 {
		log.Fatal("some wrong", err)
	}
	for _, item := range applist.Applist {
		log.Println(item.Appname)
	}
}
