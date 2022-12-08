package repository

import (
	"context"
	"fmt"

	"firebase.google.com/go/v4/auth"
	"github.com/clubo-app/clubben/auth-service/internal/datastruct"
)

type FirebaseRepository struct {
	client *auth.Client
}

func NewFirebaseRepository(client *auth.Client) *FirebaseRepository {
	return &FirebaseRepository{client: client}
}

func (repo *FirebaseRepository) CreateAnonymous(ctx context.Context, id string) (datastruct.Account, error) {
	userParams := (&auth.UserToCreate{}).UID(id)

	user, err := repo.client.CreateUser(ctx, userParams)
	if err != nil {
		return datastruct.Account{}, err
	}

	token, err := repo.client.CustomToken(ctx, user.UID)
	if err != nil {
		return datastruct.Account{}, err
	}

	return datastruct.Account{Id: user.UID, CustomToken: token}, nil
}

type CreateAccountParams struct {
	ID       string
	Email    string
	Password string
}

func (repo *FirebaseRepository) Create(ctx context.Context, params CreateAccountParams) (datastruct.Account, error) {
	fmt.Printf("%v+", params)
	userParams := (&auth.UserToCreate{}).
		UID(params.ID).
		Email(params.Email).
		Password(params.Password)

	user, err := repo.client.CreateUser(ctx, userParams)
	if err != nil {
		return datastruct.Account{}, err
	}

	token, err := repo.client.CustomToken(ctx, user.UID)
	if err != nil {
		return datastruct.Account{}, err
	}

	account := datastruct.Account{
		Id:            user.UID,
		Email:         user.Email,
		EmailVerified: user.EmailVerified,
		ProviderId:    user.ProviderID,
		CustomToken:   token,
	}

	return account, nil
}

func (repo *FirebaseRepository) CreateFromIdToken(ctx context.Context, token string) (datastruct.Account, error) {
	verifiedToken, err := repo.client.VerifyIDToken(ctx, token)
	if err != nil {
		return datastruct.Account{}, err
	}
	fmt.Printf("%v+", verifiedToken)

	// TODO: call Create method with info from claims.
	return datastruct.Account{}, nil
}

type UpdateAccountParams struct {
	UId      string
	Password string
	Email    string
}

func (repo *FirebaseRepository) Update(ctx context.Context, params UpdateAccountParams) (datastruct.Account, error) {
	userParams := (&auth.UserToUpdate{})
	if params.Password != "" {
		userParams.Password(params.Password)
	}
	if params.Email != "" {
		userParams.Email(params.Email)
	}

	user, err := repo.client.UpdateUser(ctx, params.UId, userParams)
	if err != nil {
		return datastruct.Account{}, err
	}

	account := datastruct.Account{
		Id:    user.UID,
		Email: user.Email,
	}

	return account, nil
}

func (repo *FirebaseRepository) Delete(ctx context.Context, uId string) error {
	return repo.client.DeleteUser(ctx, uId)
}

func (repo *FirebaseRepository) GetById(ctx context.Context, uId string) (datastruct.Account, error) {
	user, err := repo.client.GetUser(ctx, uId)
	if err != nil {
		return datastruct.Account{}, err
	}

	account := datastruct.Account{
		Id:    user.UID,
		Email: user.Email,
	}

	return account, nil
}

func (repo *FirebaseRepository) GetByEmail(ctx context.Context, email string) (datastruct.Account, error) {
	user, err := repo.client.GetUserByEmail(ctx, email)
	if err != nil {
		return datastruct.Account{}, err
	}

	account := datastruct.Account{
		Id:    user.UID,
		Email: user.Email,
	}

	return account, nil
}
