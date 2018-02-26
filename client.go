package main

import (
	"fmt"

	"github.com/micro/cli"
	proto "github.com/micro/examples/service/proto"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
)
// Setup and the client
func runClient(service micro.Service) {
	// Create new greeter client
	greeter := proto.NewGreeterClient("greeter", service.Client())

	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "John"})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print response
	fmt.Println(rsp.Greeting)
}

func main(){
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("v1.0"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),

		// Setup some flags. Specify --run_client to run the client

		// Add runtime flags
		// We could do this below too
		micro.Flags(cli.BoolFlag{
			Name:  "run_client",
			Usage: "Launch the client",
		}),
	)
	runClient(service)
}