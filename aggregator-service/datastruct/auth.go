package datastruct

import pbauth "github.com/clubo-app/clubben/auth-service/pb/v1"

type AggregatedAccount struct {
	Id            string             `json:"id"`
	Profile       *AggregatedProfile `json:"profile"`
	Email         string             `json:"email"`
	EmailVerified bool               `json:"email_verified"`
	ProviderId    string             `json:"provider_id"`
}

func AccountToAgg(a *pbauth.Account) *AggregatedAccount {
	if a == nil {
		return nil
	}
	return &AggregatedAccount{
		Id:            a.Id,
		Email:         a.Email,
		EmailVerified: a.EmailVerified,
		ProviderId:    a.ProviderId,
	}
}

func (a *AggregatedAccount) AddProfile(p *AggregatedProfile) *AggregatedAccount {
	a.Profile = p
	return a
}

type LoginResponse struct {
	Account *AggregatedAccount `json:"account"`
}
