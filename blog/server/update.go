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

func (s *Server) UpdateBlog(ctx context.Context, in *proto.Blog) (*emptypb.Empty, error) {
	log.Printf("UpdateBlog: %v", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	data := BlogItem{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}
	res, err := collection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": data})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if res.MatchedCount == 0 {
		return nil, status.Error(codes.NotFound, "Blog not found")
	}

	return &emptypb.Empty{}, nil
}
