package service

import (
	"context"

	"github.com/clubo-app/clubben/auth-service/internal/datastruct"
	"github.com/clubo-app/clubben/auth-service/internal/repository"
)

type AccountService interface {
	Create(context.Context, CreateAccountParams) (datastruct.Account, error)
	CreateAnonymously(context.Context, string) (datastruct.Account, error)
	Delete(context.Context, string) error
	Update(context.Context, UpdateAccountParams) (datastruct.Account, error)
	EmailTaken(ctx context.Context, email string) bool
	GetById(ctx context.Context, id string) (datastruct.Account, error)
	GetByEmail(ctx context.Context, email string) (datastruct.Account, error)
}

type accountService struct {
	repo *repository.FirebaseRepository
}

func NewAccountService(repo *repository.FirebaseRepository) AccountService {
	return &accountService{repo: repo}
}

type CreateAccountParams = repository.CreateAccountParams

func (s *accountService) Create(ctx context.Context, params repository.CreateAccountParams) (datastruct.Account, error) {
	return s.repo.Create(ctx, params)
}

func (s *accountService) CreateAnonymously(ctx context.Context, id string) (datastruct.Account, error) {
	return s.repo.CreateAnonymous(ctx, id)
}

func (s *accountService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

type UpdateAccountParams = repository.UpdateAccountParams

func (s *accountService) Update(ctx context.Context, params repository.UpdateAccountParams) (datastruct.Account, error) {
	return s.repo.Update(ctx, params)
}

func (s *accountService) EmailTaken(ctx context.Context, email string) bool {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return false
	}
	return datastruct.Account{} != user
}

func (s *accountService) GetById(ctx context.Context, id string) (datastruct.Account, error) {
	return s.repo.GetById(ctx, id)
}

func (s *accountService) GetByEmail(ctx context.Context, email string) (datastruct.Account, error) {
	return s.repo.GetByEmail(ctx, email)
}
