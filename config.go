package main

import "github.com/coreos/go-oidc"

type config struct {
	AuthURL      string
	TokenURL     string
	Oauth2State  string
	RedirectURL  string
	UserInfoURL  string
	IdpURL       string
	ClientID     string
	ClientSecret string
	Scopes       []string
}

func loadConfig() *config {
	return &config{
		AuthURL:      "http://localhost:9000/application/o/authorize/",
		TokenURL:     "http://localhost:9000/application/o/token/",
		Oauth2State:  "random",
		RedirectURL:  "http://localhost:8080/callback",
		UserInfoURL:  "http://localhost:9000/application/o/userinfo/",
		IdpURL:       "http://localhost:9000/application/o/go-app/",
		ClientID:     "NJj8L1p4cFloYB8EQzPn2gCVPLqPwi1FQcO8xj50",
		ClientSecret: "xTZ3UOHWODNzxL612rEy5uKVBe75KHJOedzxqBSYa47zhOlU2UJjIN6uhkscVZweNYYeZsDn2vOaAeDg33gFulU6PjUfMn9Tb22vHJLVNq5t3bNNqKW2WIHrKDUYXPHu",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}
}
