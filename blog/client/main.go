package main

import (
	"github.com/4925k/grpc-go/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var (
	addr = "0.0.0.0:50053"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewBlogServiceClient(conn)

	id := createBlog(client)
	//readBlog(client, id)
	//updateBlog(client, id)
	//readBlog(client, id)
	listBlogs(client)

	deleteBlog(client, id)
}
