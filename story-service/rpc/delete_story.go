package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	common "github.com/clubo-app/clubben/protobuf/common"
	sg "github.com/clubo-app/clubben/protobuf/story"
	"github.com/clubo-app/clubben/story-service/repository"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s storyServer) DeleteStory(c context.Context, req *sg.DeleteStoryRequest) (*common.MessageResponse, error) {
	_, err := ksuid.Parse(req.StoryId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Story id")
	}

	err = s.ss.Delete(c, repository.DeleteParams{
		SId: req.StoryId,
		UId: req.RequesterId,
	})
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return &common.MessageResponse{Message: "Story removed"}, nil
}
