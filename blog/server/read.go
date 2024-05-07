package main

import (
	"context"
	"github.com/4925k/grpc-go/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) ReadBlog(ctx context.Context, in *proto.BlogId) (*proto.Blog, error) {
	log.Printf("Reading Blog Id %v", in.Id)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	data := &BlogItem{}
	filter := bson.M{"_id": oid}

	res := collection.FindOne(ctx, filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return documentToBlog(data), nil
}
