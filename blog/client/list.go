package main

import (
	"context"
	"github.com/4925k/grpc-go/blog/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"io"
	"log"
)

func listBlogs(client proto.BlogServiceClient) {
	log.Printf("listBlogs was invoked")

	stream, err := client.ListBlog(context.Background(), &empty.Empty{})
	if err != nil {
		log.Fatalf("ListBlogs failed: %v", err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("ListBlogs failed: %v", err)
		}
		log.Printf("ListBlogs: %v", resp)
	}
}
