package main

import (
	"context"
	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
	"log"
)

type ProviderConfig struct {
	provider     *oidc.Provider
	oauth2Config *oauth2.Config
	oidcVerifier *oidc.IDTokenVerifier
	oauth2State  string
}

func ProviderInit() *ProviderConfig {
	// load configuration
	config := loadConfig()

	// OIDC ID Token verifier
	provider, err := oidc.NewProvider(context.Background(), config.IdpURL)
	if err != nil {
		log.Fatalf("Failed to get provider: %v", err)
	}

	oidcVerifier := provider.Verifier(&oidc.Config{ClientID: config.ClientID})

	return &ProviderConfig{
		// OAuth2 configuration
		oauth2Config: &oauth2.Config{
			ClientID:     config.ClientID,
			ClientSecret: config.ClientSecret,
			RedirectURL:  config.RedirectURL,
			Endpoint: oauth2.Endpoint{
				AuthURL:  config.AuthURL,
				TokenURL: config.TokenURL,
			},
			Scopes: config.Scopes,
		},
		provider:     provider,
		oidcVerifier: oidcVerifier,
		oauth2State:  config.Oauth2State,
	}
}
