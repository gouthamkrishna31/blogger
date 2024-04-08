package strategy

import (
	"context"
	"testing"

	pb "blogging/proto"

	"github.com/golang/mock/gomock"

	mock_strategy "blogging/strategy/mocks"
)

func TestInMemoryPostStrategy_ReadPost(t *testing.T) {
	// Create a mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock PostStrategy
	mockStrategy := mock_strategy.NewMockBlogStrategy(ctrl)

	// Set up test data
	req := &pb.GetPostRequest{
		PostId: "1",
	}
	expectedResp := &pb.GetPostResponse{
		Post: &pb.BlogPost{
			PostId:          "1",
			Title:           "Test Post",
			Content:         "This is a test post.",
			Author:          "Test Author",
			PublicationDate: "2022-01-01",
			Tags:            []string{"test", "unit", "go"},
		},
	}

	// Set up expectations
	mockStrategy.EXPECT().ReadPost(gomock.Any(), req).Return(expectedResp, nil)

	// Call the GetPost method of the strategy
	resp, err := mockStrategy.ReadPost(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Check if the response matches the expected response
	if resp.Post.PostId != expectedResp.Post.PostId {
		t.Errorf("expected PostId: %s, got: %s", expectedResp.Post.PostId, resp.Post.PostId)
	}
}

func TestInMemoryBlogStrategy_UpdatePost(t *testing.T) {
	// Create a mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock PostStrategy
	mockStrategy := mock_strategy.NewMockBlogStrategy(ctrl)

	// Set up test data
	req := &pb.UpdatePostRequest{
		PostId: "1",
		Post: &pb.BlogPost{
			PostId:          "1",
			Title:           "Updated Post",
			Content:         "This is the updated content of the post.",
			Author:          "Updated Author",
			PublicationDate: "2022-01-02",
			Tags:            []string{"updated", "unit", "go"},
		},
	}
	expectedResp := &pb.UpdatePostResponse{}

	// Set up expectations
	mockStrategy.EXPECT().UpdatePost(gomock.Any(), req).Return(expectedResp, nil)

	// Call the UpdatePost method of the strategy
	resp, err := mockStrategy.UpdatePost(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Check if the response matches the expected response
	if resp.Message != expectedResp.Message {
		t.Errorf("expected Message: %s, got: %s", expectedResp.Message, resp.Message)
	}
}

func TestInMemoryPostStrategy_DeletePost(t *testing.T) {
	// Create a mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock PostStrategy
	mockStrategy := mock_strategy.NewMockBlogStrategy(ctrl)

	// Set up test data
	req := &pb.DeletePostRequest{
		PostId: "1",
	}
	expectedResp := &pb.DeletePostResponse{}

	// Set up expectations
	mockStrategy.EXPECT().DeletePost(gomock.Any(), req).Return(expectedResp, nil)

	// Call the DeletePost method of the strategy
	resp, err := mockStrategy.DeletePost(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Check if the response matches the expected response
	if resp.Message != expectedResp.Message {
		t.Errorf("expected Message: %s, got: %s", expectedResp.Message, resp.Message)
	}
}
