package rpc

import (
	"context"
	"encoding/base64"

	"github.com/clubo-app/clubben/story-service/repository"
	"github.com/clubo-app/packages/utils"
	sg "github.com/clubo-app/protobuf/story"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s storyServer) GetByParty(c context.Context, req *sg.GetByPartyRequest) (*sg.PagedStories, error) {
	_, err := ksuid.Parse(req.PartyId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Party id")
	}

	p, err := base64.URLEncoding.DecodeString(req.NextPage)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Next Page Param")
	}

	stories, p, err := s.ss.GetByParty(c, repository.GetByPartyParams{
		PId:   req.PartyId,
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
