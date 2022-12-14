package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	sg "github.com/clubo-app/clubben/protobuf/story"
	"github.com/clubo-app/clubben/story-service/dto"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s storyServer) CreateStory(c context.Context, req *sg.CreateStoryRequest) (*sg.Story, error) {
	_, err := ksuid.Parse(req.PartyId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Party id")
	}

	d := dto.Story{
		Id:            ksuid.New(),
		PartyId:       req.PartyId,
		UserId:        req.RequesterId,
		Url:           req.Url,
		TaggedFriends: req.TaggedFriends,
	}

	story, err := s.ss.Create(c, d)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return story.ToGRPCStory(), err
}
