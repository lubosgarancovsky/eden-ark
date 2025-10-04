package model

type JWTRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JWTResponse struct {
	AccessToken  string  `json:"access_token"`
	RefreshToken string  `json:"refresh_token"`
	IdToken      *string `json:"id_token,omitempty"`
	TokenType    string  `json:"token_type"`
	ExpiresIn    int64   `json:"expires_in"`
}
