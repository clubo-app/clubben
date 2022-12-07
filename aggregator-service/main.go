package main

import (
	"log"
	"strings"

	"github.com/clubo-app/clubben/aggregator-service/config"
	authhandler "github.com/clubo-app/clubben/aggregator-service/handler/auth_handler"
	commenthandler "github.com/clubo-app/clubben/aggregator-service/handler/comment_handler"
	favoritehandler "github.com/clubo-app/clubben/aggregator-service/handler/favorite_handler"
	participationhandler "github.com/clubo-app/clubben/aggregator-service/handler/participation_handler"
	partyhandler "github.com/clubo-app/clubben/aggregator-service/handler/party_handler"
	profilehandler "github.com/clubo-app/clubben/aggregator-service/handler/profile_handler"
	relationhandler "github.com/clubo-app/clubben/aggregator-service/handler/relation_handler"
	searchhandler "github.com/clubo-app/clubben/aggregator-service/handler/search_handler"
	storyhandler "github.com/clubo-app/clubben/aggregator-service/handler/story_handler"
	pbauth "github.com/clubo-app/clubben/auth-service/pb/v1"
	"github.com/clubo-app/clubben/libs/utils/middleware"
	cg "github.com/clubo-app/clubben/protobuf/comment"
	"github.com/clubo-app/clubben/protobuf/participation"
	pg "github.com/clubo-app/clubben/protobuf/party"
	prf "github.com/clubo-app/clubben/protobuf/profile"
	rg "github.com/clubo-app/clubben/protobuf/relation"
	"github.com/clubo-app/clubben/protobuf/search"
	sg "github.com/clubo-app/clubben/protobuf/story"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	authConn, err := grpc.Dial(c.AUTH_SERVICE_ADDRESS, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Auth Service: %v", err)
	}
	defer authConn.Close()
	authClient := pbauth.NewAuthServiceClient(authConn)

	profileClient, err := prf.NewClient(c.PROFILE_SERVICE_ADDRESS)
	if err != nil {
		log.Fatalf("did not connect to profile service: %v", err)
	}
	partyClient, err := pg.NewClient(c.PARTY_SERVICE_ADDRESS)
	if err != nil {
		log.Fatalf("did not connect to party service: %v", err)
	}
	storyClient, err := sg.NewClient(c.STORY_SERVICE_ADDRESS)
	if err != nil {
		log.Fatalf("did not connect to story service: %v", err)
	}
	relationClient, err := rg.NewClient(c.RELATION_SERVICE_ADDRESS)
	if err != nil {
		log.Fatalf("did not connect to relation service: %v", err)
	}
	commentClient, err := cg.NewClient(c.COMMENT_SERVICE_ADDRESS)
	if err != nil {
		log.Fatalf("did not connect to comment service: %v", err)
	}
	participationClient, err := participation.NewClient(c.PARTICIPATION_SERVICE_ADDRESS)
	if err != nil {
		log.Fatalf("did not connect to participation service: %v", err)
	}
	searchClient, err := search.NewClient(c.SEARCH_SERVICE_ADDRESS)
	if err != nil {
		log.Fatalf("did not connect to search service: %v", err)
	}

	authHandler := authhandler.NewAuthHandler(authClient, profileClient)
	profileHandler := profilehandler.NewUserHandler(profileClient, relationClient, authClient)
	storyHandler := storyhandler.NewStoryHandler(storyClient, profileClient, partyClient)
	relationHandler := relationhandler.NewRelationHandler(relationClient, profileClient)
	commentHandler := commenthandler.NewCommentHandler(commentClient, profileClient)
	searchHandler := searchhandler.NewSearchHandler(searchClient)
	favoriteHandler := favoritehandler.NewFavoriteHandler(relationClient, partyClient, profileClient)
	partyHandler := partyhandler.NewPartyHandler(
		partyClient,
		profileClient,
		storyClient,
		relationClient,
		participationClient,
	)
	participationHandler := participationhandler.NewParticipationHandler(
		participationClient,
		partyClient,
		profileClient,
		relationClient,
	)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's an fiber.*Error
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			// Send custom error in json format
			return ctx.Status(code).JSON(err)
		},
	})
	app.Use(logger.New())
	app.Get("/dashboard", monitor.New())

	auth := app.Group("/auth")
	auth.Post("/register", authHandler.Register)

	profile := app.Group("/profiles")
	profile.Patch("/", middleware.AuthRequired(c.TOKEN_SECRET), profileHandler.UpdateUser)
	profile.Get("/me", middleware.AuthRequired(c.TOKEN_SECRET), profileHandler.GetMe)
	profile.Get("/:id", middleware.AuthOptional(c.TOKEN_SECRET), profileHandler.GetProfile)
	profile.Get("/username-taken/:username", profileHandler.UsernameTaken)
	profile.Get("/search/:query", searchHandler.SearchUsers)

	party := app.Group("/parties")
	party.Post("/", middleware.AuthRequired(c.TOKEN_SECRET), partyHandler.CreateParty)
	party.Delete("/:id", middleware.AuthRequired(c.TOKEN_SECRET), partyHandler.DeleteParty)
	party.Get("/:id", middleware.AuthOptional(c.TOKEN_SECRET), partyHandler.GetParty)
	party.Get("/user/:id", partyHandler.GetPartyByUser)
	party.Patch("/:id", middleware.AuthRequired(c.TOKEN_SECRET), partyHandler.UpdateParty)
	party.Get("/search/:query", searchHandler.SearchParties)
	party.Get("/geo-search/:lat/:lon", partyHandler.GeoSearch)

	participants := app.Group("/participants")
	participants.Delete("/:pid/invite/:uid", middleware.AuthRequired(c.TOKEN_SECRET), participationHandler.DeclinePartyInvite)
	participants.Put("/:pid/invite/:uid", middleware.AuthRequired(c.TOKEN_SECRET), participationHandler.InviteToParty)
	participants.Put("/:pid/accept/:uid", middleware.AuthRequired(c.TOKEN_SECRET), participationHandler.AcceptPartyInvite)
	participants.Get("/invite/:uid", participationHandler.GetUserInvites)
	participants.Put("/:pid", middleware.AuthRequired(c.TOKEN_SECRET), participationHandler.JoinParty)
	participants.Delete("/:pid", middleware.AuthRequired(c.TOKEN_SECRET), participationHandler.LeaveParty)
	participants.Get("/:pid", participationHandler.GetPartyParticipants)
	participants.Get("/user/:uid", participationHandler.GetUserPartyParticipation)
	participants.Get("/friends/:uid", participationHandler.GetFriendsPartyParticipation)

	favorite := app.Group("/favorites")
	favorite.Get("/user/:id", favoriteHandler.GetFavoritePartiesByUser)
	favorite.Get("/:pId", favoriteHandler.GetFavorisingUsersByParty)
	favorite.Put("/:pId", middleware.AuthRequired(c.TOKEN_SECRET), favoriteHandler.FavorParty)
	favorite.Delete("/:pId", middleware.AuthRequired(c.TOKEN_SECRET), favoriteHandler.DefavorParty)

	story := app.Group("/stories")
	story.Post("/", middleware.AuthRequired(c.TOKEN_SECRET), storyHandler.CreateStory)
	story.Delete("/:id", storyHandler.DeleteStory)
	story.Get("/party/:id", storyHandler.GetStoryByParty)
	story.Get("/user/:id", storyHandler.GetStoryByUser)
	story.Get("/presign/:key", storyHandler.PresignURL)

	friend := app.Group("/friends")
	friend.Get("/:id", middleware.AuthOptional(c.TOKEN_SECRET), relationHandler.GetFriends)
	friend.Put("/request/:id", middleware.AuthRequired(c.TOKEN_SECRET), relationHandler.CreateFriendRequest)
	friend.Put("/accept/:id", middleware.AuthRequired(c.TOKEN_SECRET), relationHandler.AcceptFriendRequest)
	friend.Delete("/request/:id", middleware.AuthRequired(c.TOKEN_SECRET), relationHandler.DeclineFriendRequest)
	friend.Delete("/:id", middleware.AuthRequired(c.TOKEN_SECRET), relationHandler.RemoveFriend)

	comment := app.Group("/comments")
	comment.Post("/party/:id", middleware.AuthRequired(c.TOKEN_SECRET), commentHandler.CreateComment)
	comment.Get("/party/:id", commentHandler.GetCommentByParty)
	comment.Delete("/:id/party/:pId", middleware.AuthRequired(c.TOKEN_SECRET), commentHandler.DeleteComment)
	comment.Post("/:id/reply", middleware.AuthRequired(c.TOKEN_SECRET), commentHandler.CreateReply)
	comment.Get("/:id/reply", commentHandler.GetReplyByComment)
	comment.Delete("/:id/reply/:rId", middleware.AuthRequired(c.TOKEN_SECRET), commentHandler.DeleteReply)

	var sb strings.Builder
	sb.WriteString("0.0.0.0:")
	sb.WriteString(c.PORT)

	if err := app.Listen(sb.String()); err != nil {
		log.Fatal(err)
	}
}
