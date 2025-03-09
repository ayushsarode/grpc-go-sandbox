package main

import ("context"
pb "github.com/ayushsarode/grpc-go-sandbox/proto"
)


func(s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello from the server",
	}, nil
}