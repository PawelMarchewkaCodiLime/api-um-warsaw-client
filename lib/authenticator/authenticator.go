package authenticator

import (
	"net/url"
)

type Authenticator struct {
	apiKey string
}

func NewAuthenticator(apiKey string) *Authenticator {
	return &Authenticator{apiKey: apiKey}
}

func (auth *Authenticator) Authenticate(query *url.Values) {
	query.Add("apikey", auth.apiKey)
}
