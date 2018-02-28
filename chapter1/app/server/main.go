package main

import (
	"context"
	"log"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	pb "github.com/rongfengliang/grpcdemo/chapter1/rpc"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
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
func (s *server) LongMessage(m *pb.QueryRequst, stream pb.UserLogin_LongMessageServer) error {
	for i := 0; i < 10; i++ {
		if err := stream.Send(&pb.MessageResponse{
			Message: "dalongdemodddd",
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

func parseToken(token string) (struct{}, error) {
	return struct{}{}, nil
}
func myfunc(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "token")
	if err != nil {
		return nil, err
	}
	tokenInfo, err := parseToken(token)

	if err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}
	if token != "dalong" {
		return nil, grpc.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}
	grpc_ctxtags.Extract(ctx).Set("auth.sub", "dalongdemo")
	newCtx := context.WithValue(ctx, "tokenInfo", tokenInfo)
	return newCtx, nil
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	/*
			grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_auth.StreamServerInterceptor(myfunc),
		)),
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				grpc_auth.UnaryServerInterceptor(myfunc),
			))
	*/
	s := grpc.NewServer(grpc.StreamInterceptor(grpc_auth.StreamServerInterceptor(myfunc)),
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(myfunc)))
	pb.RegisterUserLoginServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
