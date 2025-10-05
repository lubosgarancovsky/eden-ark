package handler

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-arc/internal/model"
	"github.com/lubosgarancovsky/eden-arc/internal/service"
	errors2 "github.com/lubosgarancovsky/eden-arc/pkg/errors"
)

// http://localhost:9091/oauth2/authorize?response_type=code&client_id=462d3f47-1684-4a5d-8ebe-bdbf03f16fce&redirect_uri=http://google.com&scope=openid email&state=123456

type OAuth2Handler struct {
	authCodeService   *service.AuthCodeService
	sessionService    *service.SessionService
	clientService     *service.ClientService
	userService       *service.UserService
	sessionCookieName string
}

func NewOAuth2Handler(
	authCodeService *service.AuthCodeService,
	sessionService *service.SessionService,
	clientService *service.ClientService,
	userService *service.UserService,
) *OAuth2Handler {
	return &OAuth2Handler{
		authCodeService:   authCodeService,
		sessionService:    sessionService,
		clientService:     clientService,
		userService:       userService,
		sessionCookieName: "session_id",
	}
}

func (h *OAuth2Handler) Authorize(c *gin.Context) {
	var authorizeQuery model.AuthorizeQuery
	if err := c.ShouldBindQuery(&authorizeQuery); err != nil {
		c.Error(err)
		return
	}

	if err := verifyAuthorizeQuery(&authorizeQuery); err != nil {
		c.Error(err)
		return
	}

	if err := h.validateClient(&authorizeQuery); err != nil {
		c.Error(err)
		return
	}

	cookie, err := c.Cookie(h.sessionCookieName)
	fmt.Println(cookie)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			fullPath := c.Request.URL.RequestURI()
			c.Redirect(302, fmt.Sprintf("/login?returnTo=%s", base64.StdEncoding.EncodeToString([]byte(fullPath))))
			return
		}

		c.Error(err)
		return
	}

	session, err := h.sessionService.FindByToken(cookie)
	if err != nil {
		c.Error(err)
		return
	}

	authCode, err := h.authCodeService.CreateAuthCode(session, &authorizeQuery)
	if err != nil {
		c.Error(err)
		return
	}

	baseURI := authCode.RedirectURI
	params := url.Values{}
	params.Add("code", authCode.Code)
	redirectURI := baseURI + "?" + params.Encode()

	c.Redirect(302, redirectURI)
	return
}

func (h *OAuth2Handler) Token(c *gin.Context) {
	// TODO: Implement token endpoint
}

func (h *OAuth2Handler) Login(c *gin.Context) {
	var input model.LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err)
		return
	}

	user, err := h.userService.FindByEmail(input.Email)
	if err != nil {
		c.Error(err)
		return
	}

	if match := h.userService.MatchPassword(input.Password, user.PasswordHash); !match {
		c.Error(errors2.ErrInvalidCredentials)
		return
	}

	session, err := h.sessionService.Insert(user.ID, c.Request.RemoteAddr, c.Request.UserAgent())
	if err != nil {
		c.Error(err)
		return
	}

	returnToParam := c.Query("returnTo")
	if returnToParam == "" {
		returnToParam = "/"
	}

	returnTo, err := base64.StdEncoding.DecodeString(returnToParam)
	if err != nil {
		c.Error(err)
		return
	}

	// TODO: Enable TLS on production (secure parameter)
	maxAge := int(time.Until(session.ExpiresAt).Seconds())
	c.SetCookie(h.sessionCookieName, session.SessionToken, maxAge, "/", "", false, true)
	c.Redirect(302, string(returnTo))
	return
}

func (h *OAuth2Handler) validateClient(query *model.AuthorizeQuery) error {
	clientID := uuid.MustParse(query.ClientID)
	client, err := h.clientService.FindByID(clientID)
	if err != nil {
		return err
	}

	if !client.IsConfidential {
		if query.CodeChallenge == nil || *query.CodeChallenge == "" {
			return errors2.ErrBadRequest.WithMessage("code_challenge is required")
		}
		if query.CodeChallengeMethod == nil {
			return errors2.ErrBadRequest.WithMessage("code_challenge_method is required")
		}
		if *query.CodeChallengeMethod == "S256" {
			return errors2.ErrBadRequest.WithMessage("code_challenge_method is not supported")
		}
	}

	for _, v := range client.RedirectUris {
		if v == query.RedirectURI {
			return nil
		}
	}

	return errors2.ErrBadRequest.WithMessage("invalid redirect uri")
}

func verifyAuthorizeQuery(query *model.AuthorizeQuery) error {
	if query.ResponseType == "" {
		return errors2.ErrBadRequest.WithMessage("response_type is required")
	}

	if query.ResponseType != "code" {
		return errors2.ErrBadRequest.WithMessage("unsupported response_type value")
	}

	if query.ClientID == "" {
		return errors2.ErrBadRequest.WithMessage("client_id is required")
	}

	if query.RedirectURI == "" {
		return errors2.ErrBadRequest.WithMessage("redirect_uri is required")
	}

	if query.State == "" {
		return errors2.ErrBadRequest.WithMessage("state is required")
	}

	if query.Scope == "" {
		return errors2.ErrBadRequest.WithMessage("scope is required")
	}

	return nil
}
