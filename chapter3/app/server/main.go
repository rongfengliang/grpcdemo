package main

import (
	"context"
	"log"
	"net"

	pb "github.com/rongfengliang/grpcdemo/chapter3"
	grpc "google.golang.org/grpc"
)

type server struct{}

func (s *server) Platlogin(ctx context.Context, in *pb.LoginEntityUser) (*pb.TokenResponse, error) {
	return &pb.TokenResponse{
		Code:    1,
		Message: "isok",
		Token:   "this is user demo project token generate",
	}, nil
}
func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:9080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserLoginServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal("server start error")
	}
}
