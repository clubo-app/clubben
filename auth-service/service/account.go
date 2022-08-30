package service

import (
	"context"
	"database/sql"
	"strings"

	"github.com/clubo-app/clubben/auth-service/dto"
	"github.com/clubo-app/clubben/auth-service/repository"
	"github.com/clubo-app/clubben/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AccountService interface {
	Create(context.Context, dto.Account) (repository.Account, error)
	Delete(context.Context, string) error
	Update(context.Context, dto.Account) (repository.Account, error)
	UpdateVerified(ctx context.Context, id, code string, emailVerified bool) (repository.Account, error)
	RotateEmailCode(ctx context.Context, email string) (repository.Account, error)
	EmailTaken(ctx context.Context, email string) bool
	GetById(ctx context.Context, id string) (repository.Account, error)
	GetByEmail(ctx context.Context, email string) (repository.Account, error)
}

type accountService struct {
	r *repository.AccountRepository
}

func NewAccountService(r *repository.AccountRepository) AccountService {
	return &accountService{r: r}
}

func (s *accountService) Create(ctx context.Context, d dto.Account) (repository.Account, error) {
	a, err := s.r.CreateAccount(ctx, repository.CreateAccountParams{
		ID:            d.ID,
		Email:         d.Email,
		EmailVerified: d.EmailVerified,
		EmailCode:     sql.NullString{String: d.EmailCode, Valid: d.EmailCode != ""},
		PasswordHash:  d.PasswordHash,
		Provider:      d.Provider,
		Type:          d.Type,
	})
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"email_idx\"") {
			return repository.Account{}, status.Error(codes.InvalidArgument, "Email already taken")
		}
		return repository.Account{}, err
	}

	return a, err
}

func (s *accountService) Delete(ctx context.Context, id string) error {
	return s.r.DeleteAccount(ctx, id)
}

func (s *accountService) Update(ctx context.Context, d dto.Account) (repository.Account, error) {
	return s.r.UpdateAccount(ctx, repository.UpdateAccountParams{
		ID:           d.ID,
		Email:        d.Email,
		EmailCode:    d.EmailCode,
		PasswordHash: d.PasswordHash,
	})
}

func (s *accountService) UpdateVerified(ctx context.Context, id, code string, emailVerified bool) (repository.Account, error) {
	return s.r.UpdateVerified(ctx, repository.UpdateVerifiedParams{
		ID:       id,
		Verified: emailVerified,
		Code:     sql.NullString{String: code, Valid: true},
	})
}

func (s *accountService) RotateEmailCode(ctx context.Context, email string) (repository.Account, error) {
	code, err := utils.GenerateOTP(4)
	if err != nil {
		return repository.Account{}, status.Error(codes.Internal, "Failed to generate Email Code")
	}

	return s.r.UpdateEmailCode(ctx, repository.UpdateEmailCodeParams{
		EmailCode: sql.NullString{String: code, Valid: true},
		Email:     email,
	})
}

func (s *accountService) EmailTaken(ctx context.Context, email string) bool {
	t, err := s.r.EmailTaken(ctx, email)
	if err != nil {
		return false
	}
	return t
}

func (s *accountService) GetById(ctx context.Context, id string) (repository.Account, error) {
	return s.r.GetAccount(ctx, id)
}

func (s *accountService) GetByEmail(ctx context.Context, email string) (repository.Account, error) {
	return s.r.GetAccountByEmail(ctx, email)
}
