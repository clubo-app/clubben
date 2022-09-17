package commenthandler

import (
	cg "github.com/clubo-app/clubben/protobuf/comment"
	"github.com/clubo-app/clubben/protobuf/profile"
	"github.com/gofiber/fiber/v2"
)

type commentHandler struct {
	commentClient cg.CommentServiceClient
	profileClient profile.ProfileServiceClient
}

type CommentHandler interface {
	CreateComment(c *fiber.Ctx) error
	DeleteComment(c *fiber.Ctx) error
	GetCommentByParty(c *fiber.Ctx) error
	CreateReply(c *fiber.Ctx) error
	DeleteReply(c *fiber.Ctx) error
	GetReplyByComment(c *fiber.Ctx) error
}

func NewCommentHandler(commentClient cg.CommentServiceClient, profileClient profile.ProfileServiceClient) CommentHandler {
	return &commentHandler{
		commentClient: commentClient,
		profileClient: profileClient,
	}
}
