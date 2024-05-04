package main

import (
	"github.com/4925k/grpc-go/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

var addr string = "localhost:50051"

func main() {
	tls := true
	opts := []grpc.DialOption{}
	if tls {
		creds, err := credentials.NewClientTLSFromFile("ssl/ca.crt", "")
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewGreetServiceClient(conn)

	doGreet(client)
	//doGreetManyTimes(client)
	//doLongGreet(client)
	//doGreetEveryone(client)
	//doGreetDeadline(client, 5*time.Second)
	//doGreetDeadline(client, 1*time.Second)

}
