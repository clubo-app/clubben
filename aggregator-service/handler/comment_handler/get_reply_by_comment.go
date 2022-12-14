package commenthandler

import (
	"strconv"

	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	"github.com/clubo-app/clubben/libs/utils"
	cg "github.com/clubo-app/clubben/protobuf/comment"
	"github.com/clubo-app/clubben/protobuf/profile"
	"github.com/gofiber/fiber/v2"
)

func (h *commentHandler) GetReplyByComment(c *fiber.Ctx) error {
	cId := c.Params("id")
	nextPage := c.Query("nextPage")

	limitStr := c.Query("limit")
	limit, _ := strconv.ParseUint(limitStr, 10, 32)
	if limit > 20 {
		return fiber.NewError(fiber.StatusBadRequest, "Max limit is 20")
	}

	rs, err := h.commentClient.GetReplyByComment(c.Context(), &cg.GetReplyByCommentRequest{CommentId: cId, NextPage: nextPage, Limit: limit})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	replyAuthors := make([]string, len(rs.Replies))
	for i, r := range rs.Replies {
		replyAuthors[i] = r.AuthorId
	}

	ps, _ := h.profileClient.GetManyProfilesMap(c.Context(), &profile.GetManyProfilesRequest{Ids: utils.UniqueStringSlice(replyAuthors)})

	aggR := make([]datastruct.AggregatedReply, len(rs.Replies))
	for i, r := range rs.Replies {
		aggR[i] = datastruct.AggregatedReply{
			Id:        r.Id,
			CommentId: r.CommentId,
			Author:    ps.Profiles[r.AuthorId],
			Body:      r.Body,
			CreatedAt: r.CreatedAt,
		}
	}

	res := datastruct.PagedAggregatedReply{
		Replies:  aggR,
		NextPage: rs.NextPage,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
