package rpc

import (
	"context"
	"encoding/base64"

	"github.com/clubo-app/clubben/libs/utils"
	sg "github.com/clubo-app/clubben/protobuf/story"
	"github.com/clubo-app/clubben/story-service/repository"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s storyServer) GetByUser(c context.Context, req *sg.GetByUserRequest) (*sg.PagedStories, error) {
	_, err := ksuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid User id")
	}

	p, err := base64.URLEncoding.DecodeString(req.NextPage)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Next Page Param")
	}

	stories, p, err := s.ss.GetByUser(c, repository.GetByUserParams{
		UId:   req.UserId,
		Limit: int(req.Limit),
		Page:  p,
	})
	if err != nil {
		return nil, utils.HandleError(err)
	}

	nextPage := base64.URLEncoding.EncodeToString(p)

	var result []*sg.Story
	for _, s := range stories {
		result = append(result, s.ToGRPCStory())
	}

	return &sg.PagedStories{
		Stories:  result,
		NextPage: nextPage,
	}, nil
}
