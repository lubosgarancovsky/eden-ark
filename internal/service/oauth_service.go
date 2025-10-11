package service

import (
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-ark/internal/config"
	"github.com/lubosgarancovsky/eden-ark/internal/model"
	"github.com/lubosgarancovsky/eden-ark/pkg/utils"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/kit"
)

type OAuthService struct {
	cfg                 *config.Config
	clientService       *ClientService
	clientSecretService *ClientSecretService
	sessionService      *SessionService
	userService         *UserService
	authCodeService     *AuthCodeService
}

func NewOAuthService(
	cfg *config.Config,
	clientService *ClientService,
	clientSecretService *ClientSecretService,
	sessionService *SessionService,
	userService *UserService,
	authCodeService *AuthCodeService,
) *OAuthService {
	return &OAuthService{
		cfg:                 cfg,
		clientService:       clientService,
		clientSecretService: clientSecretService,
		sessionService:      sessionService,
		userService:         userService,
		authCodeService:     authCodeService,
	}
}

// ValidateClientAtToken Validate client bound parameters for all token requests
func (s *OAuthService) ValidateClientAtToken(query *model.TokenQuery) (*model.Client, error) {
	if query.ClientID == "" {
		return nil, api_err.ErrBadRequest.WithMessage("invalid client id")
	}

	clientID, err := uuid.Parse(query.ClientID)
	if err != nil {
		return nil, api_err.ErrBadRequest.WithMessage("invalid client id")
	}

	// Client exists check
	client, err := s.clientService.FindByID(clientID)
	if err != nil {
		return nil, err
	}

	// Grant type check
	if !utils.Includes(client.GrantTypes, query.GrantType) {
		return client, api_err.ErrBadRequest.WithMessage("invalid grant type")
	}

	// Client secret check for confidential clients
	if _, err := s.ValidateClientSecret(client, query.ClientSecret); err != nil {
		return client, err
	}

	return client, nil
}

// ValidateClientAtAuthorize Validate client-bound parameters for all authorized requests
func (s *OAuthService) ValidateClientAtAuthorize(query *model.AuthorizeQuery) (*model.Client, error) {
	if query.ClientID == "" {
		return nil, api_err.ErrBadRequest.WithMessage("invalid client id")
	}

	clientID, err := uuid.Parse(query.ClientID)
	if err != nil {
		return nil, api_err.ErrBadRequest.WithMessage("invalid client id")
	}

	// Client exists check
	client, err := s.clientService.FindByID(clientID)
	if err != nil {
		return nil, api_err.Wrap(api_err.ErrBadRequest.WithMessage(fmt.Sprintf("client with id %s does not exist", clientID.String())), err)
	}

	// Redirect URI check
	if !utils.Includes(client.RedirectUris, query.RedirectURI) {
		return client, api_err.ErrBadRequest.WithMessage("invalid redirect uri")
	}

	// PKCE Check
	if !client.IsConfidential {
		if query.CodeChallenge == "" {
			return client, api_err.ErrBadRequest.WithMessage("code_challenge is required")
		}
		if query.CodeChallengeMethod == "" {
			return client, api_err.ErrBadRequest.WithMessage("code_challenge_method is required")
		}
		if query.CodeChallengeMethod == "S256" {
			return client, api_err.ErrBadRequest.WithMessage("code_challenge_method is not supported")
		}
	}

	return client, nil
}

// ValidateUserSession Validate user session for authorization_code & refresh_token grant types
func (s *OAuthService) ValidateUserSession(sessionID uuid.UUID) (*model.User, *model.Session, error) {
	session, err := s.sessionService.FindByID(sessionID)
	if err != nil {
		return nil, nil, err
	}

	// TODO: Should log out user if session is expired
	if session.ExpiresAt.Before(time.Now()) {
		return nil, nil, api_err.ErrUnauthorized.WithMessage("session has expired")
	}

	user, err := s.userService.FindByID(session.UserID)
	if err != nil {
		return nil, session, err
	}

	return user, session, nil
}

// ValidateAuthorizationCode Validate request for grant type authorization code
func (s *OAuthService) ValidateAuthorizationCode(client *model.Client, query *model.TokenQuery) (*model.AuthorizationCode, error) {
	code, err := s.authCodeService.FindByCode(query.Code)
	if err != nil {
		return nil, err
	}

	if code.ClientID != client.ID {
		return nil, api_err.ErrBadRequest.WithMessage("invalid client id")
	}

	if code.ExpiresAt.Before(time.Now()) {
		return nil, api_err.ErrUnauthorized.WithMessage("code has expired")
	}

	// PKCE check for non confidential clients
	if err := s.ValidatePKCE(client, code, query); err != nil {
		return nil, err
	}

	return code, nil
}

// ValidatePKCE Validate PKCE parameters for an authorization code grant type non confidential client
func (s *OAuthService) ValidatePKCE(client *model.Client, code *model.AuthorizationCode, tokenRequest *model.TokenQuery) error {
	if client.IsConfidential {
		return nil
	}

	hash := s.HashCodeVerifier(tokenRequest.CodeVerifier, code.CodeChallengeMethod)
	if hash == "" {
		return api_err.ErrUnauthorized.WithMessage("invalid code verifier")
	}

	if hash != code.CodeChallenge {
		return api_err.ErrUnauthorized.WithMessage("invalid code verifier")
	}

	return nil
}

// ValidateClientSecret Validate client secret for confidential clients
func (s *OAuthService) ValidateClientSecret(client *model.Client, secret string) (*model.ClientSecret, error) {
	if !client.IsConfidential {
		return nil, nil
	}

	return s.clientSecretService.ValidateSecret(client.ID, secret)
}

// ValidateBasicToken Validates the auth header for confidential clients refresh_token flow
func (s *OAuthService) ValidateBasicToken(client *model.Client, authHeader string) (*model.ClientSecret, error) {
	if !client.IsConfidential {
		return nil, nil
	}

	if authHeader == "" {
		return nil, api_err.ErrUnauthorized.WithMessage("missing authorization header")
	}

	parts := strings.Split(authHeader, " ")
	if parts[0] != "Basic" || len(parts) != 2 {
		return nil, api_err.ErrUnauthorized.WithMessage("invalid authorization header")
	}

	secretString, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, api_err.ErrUnauthorized.WithMessage("invalid authorization header")
	}

	secret, err := s.ValidateClientSecret(client, string(secretString))
	if err != nil {
		return nil, err
	}

	return secret, nil
}

func (s *OAuthService) ValidateRefreshToken(refreshToken string) (uuid.UUID, error) {
	publicKey, err := kit.LoadPublicKey(s.cfg.PublicKeyPath)
	if err != nil {
		return uuid.Nil, err
	}

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(refreshToken, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return uuid.Nil, fmt.Errorf("refresh token is malformed or expired")
	}

	sidValue, ok := claims["sid"]
	if !ok {
		return uuid.Nil, fmt.Errorf("refresh token missing sid")
	}

	sidStr, ok := sidValue.(string)
	if !ok {
		return uuid.Nil, fmt.Errorf("sid is not a string")
	}

	sidUUID, err := uuid.Parse(sidStr)
	if err != nil {
		return uuid.Nil, fmt.Errorf("sid is not a valid UUID")
	}

	return sidUUID, nil
}

func (s *OAuthService) GenerateTokens(client *model.Client, user *model.User, session *model.Session, query *model.TokenQuery) (*model.JWTResponse, error) {
	privateKey, err := kit.LoadPrivateKey(s.cfg.PrivateKeyPath)
	if err != nil {
		return nil, err
	}

	var idToken string
	if hasScope(query.Scope, "openid") {
		tokenStr, err := s.SignToken(s.IDTokenClaims(client, user, session), privateKey)
		if err != nil {
			return nil, err
		}
		idToken = tokenStr
	}

	accessTokenClaims := s.AccessTokenClaims(client, user, session, query)
	accessToken, err := s.SignToken(accessTokenClaims, privateKey)
	refreshToken, err := s.SignToken(s.RefreshTokenClaims(client, user, session, query, ""), privateKey)

	return &model.JWTResponse{
		IdToken:      &idToken,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    accessTokenClaims["exp"].(int64),
	}, nil
}

func (s *OAuthService) IDTokenClaims(client *model.Client, user *model.User, session *model.Session) jwt.MapClaims {
	return jwt.MapClaims{
		"iss":         s.cfg.Issuer,
		"sub":         user.ID.String(),
		"aud":         client.ID.String(),
		"exp":         time.Now().Add(time.Duration(s.cfg.AccessExp) * time.Second).Unix(),
		"iat":         time.Now().Unix(),
		"nonce":       session.Nonce,
		"given_name":  user.FirstName,
		"family_name": user.LastName,
		"username":    user.Username,
		"email":       user.Email,
		"role":        user.Role,
		"amr":         "pwd",
		"azp":         client.ID.String(),
	}
}

func (s *OAuthService) AccessTokenClaims(client *model.Client, user *model.User, session *model.Session, query *model.TokenQuery) jwt.MapClaims {
	return jwt.MapClaims{
		"iss":       s.cfg.Issuer,
		"sub":       user.ID.String(),
		"aud":       "", // TODO: Add address of API gateway
		"exp":       time.Now().Add(time.Duration(s.cfg.AccessExp) * time.Second).Unix(),
		"iat":       time.Now().Unix(),
		"scope":     query.Scope,
		"client_id": client.ID.String(),
		"jti":       uuid.New(),
		"role":      user.Role,
		"sid":       session.ID.String(),
	}
}

func (s *OAuthService) RefreshTokenClaims(client *model.Client, user *model.User, session *model.Session, query *model.TokenQuery, prevJti string) jwt.MapClaims {
	return jwt.MapClaims{
		"iss":          s.cfg.Issuer,
		"sub":          user.ID.String(),
		"aud":          s.cfg.Issuer,
		"exp":          time.Now().Add(time.Duration(s.cfg.RefreshExp) * time.Second).Unix(),
		"iat":          time.Now().Unix(),
		"scope":        query.Scope,
		"client_id":    client.ID.String(),
		"jti":          uuid.New(),
		"role":         user.Role,
		"sid":          session.ID.String(),
		"rot":          true,
		"previous_jti": prevJti,
	}
}

func (s *OAuthService) GetProfile(userID uuid.UUID) (*model.User, error) {
	return s.userService.FindByID(userID)
}

func (s *OAuthService) SignToken(claims jwt.MapClaims, privateKey *rsa.PrivateKey) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(privateKey)
}

func (s *OAuthService) SaveAuthCode(sessionToken string, query *model.AuthorizeQuery) (*model.Session, *model.AuthorizationCode, error) {
	session, err := s.sessionService.FindByToken(sessionToken)
	if err != nil {
		return nil, nil, err
	}

	authCode, err := s.authCodeService.CreateAuthCode(session, query)
	if err != nil {
		return nil, nil, err
	}

	return session, authCode, nil
}

func (s *OAuthService) SaveNonce(session *model.Session, nonce string) (*model.Session, error) {
	return s.sessionService.UpdateNonce(session.ID, nonce)
}

func (s *OAuthService) DeleteAuthCode(code *model.AuthorizationCode) error {
	return s.authCodeService.DeleteAuthCode(code.Code)
}

func (s *OAuthService) DeleteSession(sessionToken string) error {
	return s.sessionService.DeleteByToken(sessionToken)
}

func (s *OAuthService) SaveSession(request *model.LoginRequest, ipAddr string, userAgent string, nonce string) (*model.Session, error) {
	user, err := s.GetUser(request.Username)
	if err != nil {
		return nil, err
	}

	if match := s.userService.MatchPassword(request.Password, user.PasswordHash); !match {
		return nil, api_err.ErrUnauthorized.WithMessage("invalid credentials")
	}

	session, err := s.sessionService.Insert(user.ID, ipAddr, userAgent, nonce)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (s *OAuthService) GetUser(username string) (*model.User, error) {
	if strings.Contains(username, "@") {
		return s.userService.FindByEmail(username)
	}

	return s.userService.FindByUsername(username)
}

func (s *OAuthService) HashCodeVerifier(codeVerifier string, method string) string {
	if codeVerifier == "" {
		return ""
	}

	if method == "S256" {
		hash := sha256.Sum256([]byte(codeVerifier))
		return base64.RawURLEncoding.EncodeToString(hash[:])
	}

	return ""
}

func hasScope(scope string, target string) bool {
	scopes := strings.Split(scope, " ")
	return utils.Includes(scopes, target)
}
