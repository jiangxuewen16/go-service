package main

import (
	"google.golang.org/grpc"
	"log"
	pb "study_grpc/helloworld"
	"os"
	"golang.org/x/net/context"
)

const (
	address = "localhost:8888"
	defualtName = "world"
	defualtCode = 10010
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := defualtName

	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name:name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r)
}
