package main

import (
	"context"
	"github.com/4925k/grpc-go/blog/proto"
	"log"
)

func deleteBlog(client proto.BlogServiceClient, id string) {
	log.Printf("Deleting Blog %s\n", id)

	_, err := client.DeleteBlog(context.Background(), &proto.BlogId{Id: id})
	if err != nil {
		log.Fatalf("DeleteBlog err: %v", err)
	}

	log.Printf("Deleted Blog %s\n", id)
}
