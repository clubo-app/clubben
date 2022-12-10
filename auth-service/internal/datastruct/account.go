package datastruct

import pbauth "github.com/clubo-app/clubben/auth-service/pb/v1"

type Account struct {
	Id            string
	Email         string
	EmailVerified bool
	Password      string
	ProviderId    string
	CustomToken   string
}

func (a Account) ToGRPCAccount() *pbauth.Account {
	return &pbauth.Account{
		Id:            a.Id,
		Email:         a.Email,
		EmailVerified: a.EmailVerified,
		ProviderId:    a.ProviderId,
	}
}
