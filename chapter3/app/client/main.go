package main

import (
	"context"
	"log"

	pb "github.com/rongfengliang/grpcdemo/chapter3"
	grpc "google.golang.org/grpc"
)

func main() {
	con, err := grpc.Dial("127.0.0.1:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("connect error")
	}
	client := pb.NewUserLoginClient(con)
	result, err := client.Platlogin(context.Background(), &pb.LoginEntityUser{})
	if err != nil {
		log.Println("request is error")
	}
	log.Println(result.Token)

}
