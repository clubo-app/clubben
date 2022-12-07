package datastruct

type AggregatedAccount struct {
	Id      string             `json:"id"`
	Profile *AggregatedProfile `json:"profile"`
	Email   string             `json:"email"`
}

type LoginResponse struct {
	Account AggregatedAccount `json:"account"`
}
