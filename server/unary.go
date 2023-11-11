package main

import (
	"context"

	pb "github.com/tanvir/grpc-demo-yt/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello(From Unary response.)",
	}, nil
}
