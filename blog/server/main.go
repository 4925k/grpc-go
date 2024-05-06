package server

import (
	"context"
	"github.com/4925k/grpc-go/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var (
	addr       = "0.0.0.0:5003"
	collection *mongo.Collection
)

type Server struct {
	proto.BlogServiceServer
}

func main() {
	// mongo db client
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatalf("failed to connect to mongodb: %v", err)
	}

	collection = client.Database("blog").Collection("blog")

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("listening on %s", addr)

	s := grpc.NewServer()
	proto.RegisterBlogServiceServer(s, &Server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
