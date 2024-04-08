package service

import (
    "context"
    pb "blogging/proto"
)
type BlogStrategy interface {
	CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error)
	ReadPost(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostResponse, error)
	UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.UpdatePostResponse, error)
	DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error)
}
