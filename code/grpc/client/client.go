package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "grpc/proto"

	"google.golang.org/grpc"
)
const (
	address     = "localhost:8080"//在put-forward上的本机端口
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMessageSenderClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Send(ctx, &pb.MessageRequest{SaySomething: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetResponseSomething())
	/*r, err = c.SayHelloAgain(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
			log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())*/
}
