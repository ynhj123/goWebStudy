package main

import (
	"context"
	pb "grpcExample/proto"
	"log"
	"server/mapper"
)

type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(_ context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	user := mapper.GetByUserName(in.GetName())
	return &pb.HelloReply{Message: "Hello " + user.Nickname}, nil
}
