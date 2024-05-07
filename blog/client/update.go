package main

import (
	"context"
	"github.com/4925k/grpc-go/blog/proto"
	"log"
)

func updateBlog(client proto.BlogServiceClient, id string) {
	log.Printf("updating blog with id %s", id)

	newBlog := &proto.Blog{
		Id:       id,
		AuthorId: "Mln",
		Title:    "Thursday",
		Content:  "This is my updated content",
	}

	_, err := client.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Printf("error updating blog: %v", err)
	}

	log.Printf("blog with id %s updated", id)

}
