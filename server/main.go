package main

import "golang.org/x/net/context"
import (
	pb "study_grpc/helloworld"
	"net"
	"log"
	"google.golang.org/grpc"
)

const port  = ":8888"

type server struct {}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message:"Hello" + in.Name, Code:10010}, nil
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}
