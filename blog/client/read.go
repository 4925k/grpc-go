package main

import (
	"context"
	"github.com/4925k/grpc-go/blog/proto"
	"log"
)

func readBlog(client proto.BlogServiceClient, id string) *proto.Blog {
	log.Printf("reading blogID %s\n", id)

	req := proto.BlogId{Id: id}
	res, err := client.ReadBlog(context.Background(), &req)
	if err != nil {
		log.Fatalf("read blog err %v", err)
	}

	log.Printf("read blog %v\n", res)

	return res
}
