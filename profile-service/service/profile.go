package service

import (
	"context"
	"database/sql"

	"github.com/clubo-app/clubben/libs/stream"
	"github.com/clubo-app/clubben/profile-service/dto"
	"github.com/clubo-app/clubben/profile-service/repository"
	"github.com/clubo-app/clubben/protobuf/events"
	"github.com/clubo-app/clubben/protobuf/profile"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProfileService interface {
	Create(context.Context, dto.Profile) (repository.Profile, error)
	Update(context.Context, dto.Profile) (repository.Profile, error)
	Delete(ctx context.Context, id string) error
	UsernameTaken(ctx context.Context, username string) bool
	GetById(ctx context.Context, id string) (repository.Profile, error)
	GetMany(ctx context.Context, ids []string) ([]repository.Profile, error)
}

type profileService struct {
	r      *repository.ProfileRepository
	stream stream.Stream
}

func NewProfileService(r *repository.ProfileRepository, stream stream.Stream) ProfileService {
	return &profileService{r: r, stream: stream}
}

func (s *profileService) Create(ctx context.Context, dp dto.Profile) (repository.Profile, error) {
	p, err := s.r.CreateProfile(ctx, repository.CreateProfileParams{
		ID:        dp.ID,
		Username:  dp.Username,
		Firstname: dp.Firstname,
		Lastname:  sql.NullString{String: dp.Lastname, Valid: dp.Lastname != ""},
		Avatar:    sql.NullString{String: dp.Avatar, Valid: dp.Avatar != ""},
	})
	if err != nil {
		return repository.Profile{}, err
	}

	s.stream.PublishEvent(&events.ProfileCreated{Profile: p.ToGRPCProfile()})

	return p, nil
}

func (s *profileService) Update(ctx context.Context, dp dto.Profile) (repository.Profile, error) {
	p, err := s.r.UpdateProfile(ctx, repository.UpdateProfileParams{
		ID:        dp.ID,
		Username:  dp.Username,
		Firstname: dp.Firstname,
		Lastname:  dp.Lastname,
		Avatar:    dp.Avatar,
	})
	if err != nil {
		return repository.Profile{}, err
	}

	updatedValues := profile.Profile{
		Id: dp.ID,
	}
	if dp.Firstname != "" {
		updatedValues.Firstname = dp.Firstname
	}
	if dp.Lastname != "" {
		updatedValues.Lastname = dp.Lastname
	}
	if dp.Username != "" {
		updatedValues.Username = dp.Username
	}

	s.stream.PublishEvent(&events.ProfileUpdated{Profile: &updatedValues})

	return p, nil
}

func (s *profileService) Delete(ctx context.Context, id string) error {
	err := s.r.DeleteProfile(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *profileService) UsernameTaken(ctx context.Context, username string) bool {
	t, err := s.r.UsernameTaken(ctx, username)
	if err != nil {
		return false
	}
	return t
}

func (s *profileService) GetById(ctx context.Context, id string) (repository.Profile, error) {
	p, err := s.r.GetProfile(ctx, id)
	if err != nil {
		return repository.Profile{}, status.Error(codes.InvalidArgument, "No Profile found")
	}

	return p, nil
}

func (s *profileService) GetMany(ctx context.Context, ids []string) ([]repository.Profile, error) {
	ps, err := s.r.GetManyProfiles(ctx, repository.GetManyProfilesParams{
		Ids:   ids,
		Limit: int32(len(ids)),
	})
	if err != nil {
		return []repository.Profile{}, err
	}

	return ps, nil
}
