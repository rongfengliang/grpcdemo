package main

import (
	"context"
	"log"
	"net"

	pb "git.yonyou.com/go/grpcdemo/chapter2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) List(ctx context.Context, in *pb.AppRequest) (*pb.AppListResponse, error) {
	applist := []*pb.AppInfo{
		&pb.AppInfo{
			Appid:       "001",
			Appname:     "firstapp",
			Appdescribe: "userlogin firstapp",
		},
		&pb.AppInfo{
			Appid:       "002",
			Appname:     "secondapp",
			Appdescribe: "appsign secondapp",
		},
	}
	return &pb.AppListResponse{
		Applist: applist,
		Code:    0,
		Message: "ok",
	}, nil
}
func (s *server) Delte(ctx context.Context, in *pb.AppRequest) (*pb.OperatorResponse, error) {
	return &pb.OperatorResponse{
		Code: 0, Message: "ok",
	}, nil
}
func (s *server) Echo(ctx context.Context, in *pb.StringMessage) (*pb.StringMessage, error) {
	return &pb.StringMessage{
		Value: "demo echo message",
	}, nil
}

const (
	port = "0.0.0.0:9090"
)

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAppListServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
