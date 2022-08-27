package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/digitalocean/godo"
)

type UploadService interface {
	Upload(ctx context.Context, filename, base64File string) (string, error)
}

type uploadService struct {
	awsSess  *session.Session
	doClient *godo.Client
}

func NewUploadService(spacesEndpoint, spacesToken string) UploadService {
	region := strings.Split(spacesEndpoint, ".")[0]
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: &spacesEndpoint,
		Region:   &region,
	}))

	doClient := godo.NewFromToken(spacesToken)

	return &uploadService{awsSess: sess, doClient: doClient}
}

func (as *uploadService) Upload(ctx context.Context, filename, base64File string) (string, error) {
	decode, err := base64.StdEncoding.DecodeString(base64File)
	if err != nil {
		return "", err
	}

	uploader := s3manager.NewUploader(as.awsSess)

	res, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("avatars"),
		Key:    aws.String(filename),
		ACL:    aws.String("public-read"),
		Body:   bytes.NewReader(decode),
	})
	if err != nil {
		return "", err
	}

	flushRequest := &godo.CDNFlushCacheRequest{
		Files: []string{filename},
	}

	_, err = as.doClient.CDNs.FlushCache(ctx, "19f06b6a-3ace-4315-b086-499a0e521b76", flushRequest)
	if err != nil {
		return "", err
	}

	return res.Location, nil
}
