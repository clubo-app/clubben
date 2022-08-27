package rpc

import (
	"context"

	"github.com/clubo-app/packages/utils"
	sg "github.com/clubo-app/protobuf/story"
)

func (s storyServer) PresignURL(c context.Context, req *sg.PresignURLRequest) (*sg.PresignURLResponse, error) {
	url, err := s.us.PresignURL(c, req.Key)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return &sg.PresignURLResponse{Url: url}, nil
}
