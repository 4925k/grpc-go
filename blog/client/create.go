package main

import (
	"context"
	"github.com/4925k/grpc-go/blog/proto"
	"log"
)

func createBlog(c proto.BlogServiceClient) string {
	log.Println("Creating blog client")

	blog := proto.Blog{
		AuthorId: "Dbk",
		Title:    "First Blog",
		Content:  "Hello world",
	}

	res, err := c.CreateBlog(context.Background(), &blog)
	if err != nil {
		log.Fatalf("could not create blog: %v", err)
	}

	log.Printf("Blog created: %v", res.Id)

	return res.Id
}
