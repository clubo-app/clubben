package main

import (
	"context"
	"log"
	"strings"

	firebase "firebase.google.com/go/v4"
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
	firebaseauth "github.com/clubo-app/clubben/libs/firebase-auth"
	cg "github.com/clubo-app/clubben/protobuf/comment"
	"github.com/clubo-app/clubben/protobuf/participation"
	pbparty "github.com/clubo-app/clubben/party-service/pb/v1"
	prf "github.com/clubo-app/clubben/protobuf/profile"
	rg "github.com/clubo-app/clubben/protobuf/relation"
	"github.com/clubo-app/clubben/protobuf/search"
	sg "github.com/clubo-app/clubben/protobuf/story"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"google.golang.org/api/option"
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

  partyConn, err := grpc.Dial(c.PARTY_SERVICE_ADDRESS, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to PartyService: %v", err)
	}
	defer partyConn.Close()
	partyClient := pbparty.NewPartyServiceClient(partyConn)

	profileClient, err := prf.NewClient(c.PROFILE_SERVICE_ADDRESS)
	if err != nil {
		log.Fatalf("did not connect to profile service: %v", err)
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

	authHandler := authhandler.NewAuthHandler(authClient, profileClient, relationClient)
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

	opt := option.WithCredentialsFile(c.GOOGLE_APPLICATION_CREDENTIALS)
	firebase, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal(err)
	}

	requireAuth := firebaseauth.New(firebaseauth.Config{FirebaseApp: firebase})
	optionalAuth := firebaseauth.New(firebaseauth.Config{FirebaseApp: firebase, AuthOptional: true})

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
	auth.Get("/me", requireAuth, authHandler.GetMe)
	auth.Get("/login/:id", authHandler.GetToken)
	auth.Post("/register", authHandler.Register)
	auth.Post("/register-anon", authHandler.RegisterAnonymously)

	app.Route("/profiles", func(api fiber.Router) {
		api.Patch("/", requireAuth, profileHandler.UpdateUser)
		api.Get("/:id", optionalAuth, profileHandler.GetProfile)
		api.Get("/username-taken/:username", profileHandler.UsernameTaken)
		api.Get("/search/:query", searchHandler.SearchUsers)
	})

	party := app.Group("/parties")
	party.Post("/", requireAuth, partyHandler.CreateParty)
	party.Delete("/:id", requireAuth, partyHandler.DeleteParty)
	party.Get("/:id", optionalAuth, partyHandler.GetParty)
	party.Get("/user/:id", partyHandler.GetPartyByUser)
	party.Patch("/:id", requireAuth, partyHandler.UpdateParty)
	party.Get("/search/:query", searchHandler.SearchParties)
	party.Get("/geo-search/:lat/:lon", partyHandler.GeoSearch)

	participants := app.Group("/participants")
	participants.Delete("/:pid/invite/:uid", requireAuth, participationHandler.DeclinePartyInvite)
	participants.Put("/:pid/invite/:uid", requireAuth, participationHandler.InviteToParty)
	participants.Put("/:pid/accept/:uid", requireAuth, participationHandler.AcceptPartyInvite)
	participants.Get("/invite/:uid", participationHandler.GetUserInvites)
	participants.Put("/:pid", requireAuth, participationHandler.JoinParty)
	participants.Delete("/:pid", requireAuth, participationHandler.LeaveParty)
	participants.Get("/:pid", participationHandler.GetPartyParticipants)
	participants.Get("/user/:uid", participationHandler.GetUserPartyParticipation)
	participants.Get("/friends/:uid", participationHandler.GetFriendsPartyParticipation)

	favorite := app.Group("/favorites")
	favorite.Get("/user/:id", favoriteHandler.GetFavoritePartiesByUser)
	favorite.Get("/:pId", favoriteHandler.GetFavorisingUsersByParty)
	favorite.Put("/:pId", requireAuth, favoriteHandler.FavorParty)
	favorite.Delete("/:pId", requireAuth, favoriteHandler.DefavorParty)

	story := app.Group("/stories")
	story.Post("/", requireAuth, storyHandler.CreateStory)
	story.Delete("/:id", storyHandler.DeleteStory)
	story.Get("/party/:id", storyHandler.GetStoryByParty)
	story.Get("/user/:id", storyHandler.GetStoryByUser)
	story.Get("/presign/:key", storyHandler.PresignURL)

	friend := app.Group("/friends")
	friend.Get("/:id", optionalAuth, relationHandler.GetFriends)
	friend.Put("/request/:id", requireAuth, relationHandler.CreateFriendRequest)
	friend.Put("/accept/:id", requireAuth, relationHandler.AcceptFriendRequest)
	friend.Delete("/request/:id", requireAuth, relationHandler.DeclineFriendRequest)
	friend.Delete("/:id", requireAuth, relationHandler.RemoveFriend)

	comment := app.Group("/comments")
	comment.Post("/party/:id", requireAuth, commentHandler.CreateComment)
	comment.Get("/party/:id", commentHandler.GetCommentByParty)
	comment.Delete("/:id/party/:pId", requireAuth, commentHandler.DeleteComment)
	comment.Post("/:id/reply", requireAuth, commentHandler.CreateReply)
	comment.Get("/:id/reply", commentHandler.GetReplyByComment)
	comment.Delete("/:id/reply/:rId", requireAuth, commentHandler.DeleteReply)

	var sb strings.Builder
	sb.WriteString("0.0.0.0:")
	sb.WriteString(c.PORT)

	if err := app.Listen(sb.String()); err != nil {
		log.Fatal(err)
	}
}
