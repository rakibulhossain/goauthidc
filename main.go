package main

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

func main() {
	provider := ProviderInit()
	// Define routes
	//http.HandleFunc("/", homeHandler)
	http.HandleFunc("/", provider.loginHandler)
	http.HandleFunc("/callback", provider.callbackHandler)

	// Start server
	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome! <a href='/login'>Login with Authentik</a>")
}

func (provider *ProviderConfig) loginHandler(w http.ResponseWriter, r *http.Request) {
	// Generate OAuth2 URL
	url := provider.oauth2Config.AuthCodeURL(provider.oauth2State, oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusFound)
}

func (provider *ProviderConfig) callbackHandler(w http.ResponseWriter, r *http.Request) {
	// Verify OAuth2 state
	if r.FormValue("state") != provider.oauth2State {
		http.Error(w, "Invalid state", http.StatusBadRequest)
		return
	}

	// Exchange code for token
	token, err := provider.oauth2Config.Exchange(context.Background(), r.FormValue("code"))
	if err != nil {
		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
		return
	}

	// Verify ID Token
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "Missing ID token", http.StatusInternalServerError)
		return
	}

	// Parse and verify ID Token
	idToken, err := provider.oidcVerifier.Verify(context.Background(), rawIDToken)
	if err != nil {
		http.Error(w, "Failed to verify ID token", http.StatusInternalServerError)
		return
	}

	// Extract claims
	var claims struct {
		Email string `json:"email"`
	}
	if err := idToken.Claims(&claims); err != nil {
		http.Error(w, "Failed to parse claims", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Email: %s\n", claims.Email)
}
