package main

import (
	"context"
	"log"
	"net"

	pb "github.com/rongfengliang/grpcdemo/chapter1/rpc"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) AppLogin(ctx context.Context, in *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	return &pb.UserLoginResponse{
		Code:    0,
		Message: "sucess",
		Token:   "dalongdmeo",
	}, nil
}

func (s *server) TokenVerify(ctx context.Context, in *pb.UserTokenRequest) (*pb.TokenVerifyResponse, error) {
	return &pb.TokenVerifyResponse{
		Message:  "ok",
		Code:     0,
		Isverify: true,
	}, nil
}
func (s *server) LongMessage(stream pb.UserLogin_LongMessageServer) error {
	for {
		if err := stream.Send(&pb.MessageResponse{
			Message: "dalongdemo",
			Code:    0,
		}); err != nil {
			return err
		}
	}
	return nil
}

const (
	port = "0.0.0.0:5001"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserLoginServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
