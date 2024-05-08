package main

import (
	"context"
	"github.com/4925k/grpc-go/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (s *Server) DeleteBlog(ctx context.Context, id *proto.BlogId) (*emptypb.Empty, error) {
	log.Printf("DeleteBlog(%v)", id)

	oid, err := primitive.ObjectIDFromHex(id.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if res.DeletedCount == 0 {
		return nil, status.Error(codes.NotFound, "Blog not found")
	}

	return &emptypb.Empty{}, nil
}
