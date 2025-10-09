package model

import "github.com/google/uuid"

type UserContext struct {
	ID   uuid.UUID
	Role string
}

type AuthorizeQuery struct {
	ResponseType        string `form:"response_type"`
	ClientID            string `form:"client_id"`
	RedirectURI         string `form:"redirect_uri"`
	Scope               string `form:"scope"`
	State               string `form:"state"`
	Nonce               string `form:"nonce"`
	CodeChallenge       string `form:"code_challenge"`
	CodeChallengeMethod string `form:"code_challenge_method"`
}

type TokenQuery struct {
	GrantType    string `form:"grant_type"`
	ClientID     string `form:"client_id"`
	ClientSecret string `form:"client_secret"`
	Code         string `form:"code"`
	RedirectURI  string `form:"redirect_uri"`
	Scope        string `form:"scope"`
	CodeVerifier string `form:"code_verifier"`
}

type TokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWTResponse struct {
	AccessToken  string  `json:"access_token"`
	RefreshToken string  `json:"refresh_token"`
	IdToken      *string `json:"id_token,omitempty"`
	TokenType    string  `json:"token_type"`
	ExpiresIn    int64   `json:"expires_in"`
}
