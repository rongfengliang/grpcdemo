package main

import (
	"context"
	"log"

	pb "github.com/rongfengliang/grpcdemo/chapter3"
	platform "github.com/rongfengliang/grpcdemo/platform"
	grpc "google.golang.org/grpc"
)

func main() {
	con, err := grpc.Dial("47.52.58.151:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("connect error")
	}
	client := pb.NewUserLoginClient(con)
	result, err := client.Platlogin(context.Background(), &pb.LoginEntityUser{
		User: &pb.User{
			Usertype: pb.UserType_ADMIN,
			Export: &platform.Export{
				Exporttype:      "dalongdemo",
				Exporturl:       "https://github.com/rongfengliang",
				Webhookurl:      "https://github.com/rongfengliang",
				Webhooksecurity: "appdemo",
			},
		},
		Loginentity: &pb.LoginEntity{
			Usertype: pb.UserType_ADMIN,
			Name:     "dalong",
			Password: "appdemo",
		},
	})
	if err != nil {
		log.Println("request is error")
	}
	log.Println(result.Token)

}
