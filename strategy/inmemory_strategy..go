package strategy

import (
	pb "blogging/proto"
	"context"

	"google.golang.org/grpc"
)

type InMemoryPostStrategy struct {
	posts map[string]*pb.BlogPost
}

func NewInMemoryBlogStrategy() *InMemoryPostStrategy {
	return &InMemoryPostStrategy{
		posts: make(map[string]*pb.BlogPost),
	}
}

func (s *InMemoryPostStrategy) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	post := req.Post
	s.posts[post.PostId] = post
	return &pb.CreatePostResponse{PostId: post.PostId}, nil
}

func (s *InMemoryPostStrategy) ReadPost(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	post, exists := s.posts[req.PostId]
	if !exists {
		return nil, grpc.Errorf(404, "Post not found")
	}
	return &pb.GetPostResponse{Post: post}, nil
}

func (s *InMemoryPostStrategy) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.UpdatePostResponse, error) {
	postId := req.PostId
	if _, exists := s.posts[postId]; !exists {
		return nil, grpc.Errorf(404, "Post not found")
	}
	s.posts[postId] = req.Post
	return &pb.UpdatePostResponse{Message: "Post updated successfully"}, nil
}

func (s *InMemoryPostStrategy) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	postId := req.PostId
	if _, exists := s.posts[postId]; !exists {
		return nil, grpc.Errorf(404, "Post not found")
	}
	delete(s.posts, postId)
	return &pb.DeletePostResponse{Message: "Post deleted successfully"}, nil
}
