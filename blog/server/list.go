package main

import (
	"context"
	"github.com/4925k/grpc-go/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (s *Server) ListBlog(in *emptypb.Empty, stream proto.BlogService_ListBlogServer) error {
	log.Printf("ListBlog called")

	cur, err := collection.Find(context.Background(), primitive.D{{}})
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &BlogItem{}
		err = cur.Decode(data)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		err := stream.Send(documentToBlog(data))
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	if err = cur.Err(); err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}
