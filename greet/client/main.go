package main

import (
	"github.com/4925k/grpc-go/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewGreetServiceClient(conn)

	//doGreet(client)
	//doGreetManyTimes(client)
	//doLongGreet(client)
	doGreetEveryone(client)
}
