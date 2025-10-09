package handler

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-arc/internal/model"
	"github.com/lubosgarancovsky/eden-arc/internal/service"
	"github.com/lubosgarancovsky/eden-arc/pkg/helpers"
	"github.com/lubosgarancovsky/go-kit/api_err"
)

type OAuth2Handler struct {
	oauthService      *service.OAuthService
	sessionCookieName string
}

func NewOAuth2Handler(oauthService *service.OAuthService) *OAuth2Handler {
	return &OAuth2Handler{
		oauthService:      oauthService,
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

	if _, err := h.oauthService.ValidateClientAtAuthorize(&authorizeQuery); err != nil {
		c.Error(err)
		return
	}

	cookie, err := c.Cookie(h.sessionCookieName)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			fullPath := c.Request.URL.RequestURI()
			c.Redirect(302, fmt.Sprintf("/login?returnTo=%s", base64.StdEncoding.EncodeToString([]byte(fullPath))))
			return
		}

		c.Error(err)
		return
	}

	_, authCode, err := h.oauthService.SaveAuthCode(cookie, &authorizeQuery)
	if err != nil {
		c.Error(err)
		return
	}

	baseURI := authCode.RedirectURI
	params := url.Values{}
	params.Add("code", authCode.Code)
	params.Add("state", authorizeQuery.State)
	redirectURI := baseURI + "?" + params.Encode()

	c.Redirect(302, redirectURI)
	return
}

func (h *OAuth2Handler) Token(c *gin.Context) {
	var tokenQuery model.TokenQuery
	if err := c.ShouldBind(&tokenQuery); err != nil {
		c.Error(err)
		return
	}

	client, err := h.oauthService.ValidateClientAtToken(&tokenQuery)
	if err != nil {
		c.Error(err)
		return
	}

	switch tokenQuery.GrantType {
	case "authorization_code":
		h.handleAuthCodeGrantType(c, client, tokenQuery)
		return
	case "refresh_token":
		h.handleRefreshGrantType(c, client, tokenQuery)
		return
	default:
		c.Error(api_err.ErrBadRequest.WithMessage("unsupported grant type"))
		return
	}

}

func (h *OAuth2Handler) Logout(c *gin.Context) {
	// TODO: Logout user
	return
}

func (h *OAuth2Handler) Login(c *gin.Context) {
	var input model.LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err)
		return
	}

	session, err := h.oauthService.SaveSession(&input, c.Request.RemoteAddr, c.Request.UserAgent())
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

func (h *OAuth2Handler) Profile(c *gin.Context) {
	userCtx, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	user, err := h.oauthService.GetProfile(userCtx.ID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, user)
}

func (h *OAuth2Handler) handleAuthCodeGrantType(c *gin.Context, client *model.Client, tokenQuery model.TokenQuery) {
	code, err := h.oauthService.ValidateAuthorizationCode(client, &tokenQuery)
	if err != nil {
		c.Error(err)
		return
	}

	sessionID, err := c.Cookie(h.sessionCookieName)
	if err != nil {
		c.Error(err)
		return
	}

	user, session, err := h.oauthService.ValidateUserSession(sessionID)
	if err != nil {
		c.Error(err)
		return
	}

	jwtPayload, err := h.oauthService.GenerateTokens(client, user, session, &tokenQuery)
	if err != nil {
		c.Error(err)
		return
	}

	h.oauthService.DeleteAuthCode(code)

	c.JSON(200, jwtPayload)
}

func (h *OAuth2Handler) handleRefreshGrantType(c *gin.Context, client *model.Client, tokenQuery model.TokenQuery) {
	authHeader := c.GetHeader("Authorization")
	if _, err := h.oauthService.ValidateBasicToken(client, authHeader); err != nil {
		c.Error(err)
		return
	}

	sessionID, err := c.Cookie(h.sessionCookieName)
	if err != nil {
		c.Error(err)
		return
	}

	user, session, err := h.oauthService.ValidateUserSession(sessionID)
	if err != nil {
		c.Error(err)
		return
	}

	jwtPayload, err := h.oauthService.GenerateTokens(client, user, session, &tokenQuery)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, jwtPayload)
}

func verifyAuthorizeQuery(query *model.AuthorizeQuery) error {
	if query.ResponseType == "" {
		return api_err.ErrBadRequest.WithMessage("response_type is required")
	}

	if query.ResponseType != "code" {
		return api_err.ErrBadRequest.WithMessage("unsupported response_type value")
	}

	if query.ClientID == "" {
		return api_err.ErrBadRequest.WithMessage("client_id is required")
	}

	if query.RedirectURI == "" {
		return api_err.ErrBadRequest.WithMessage("redirect_uri is required")
	}

	if query.State == "" {
		return api_err.ErrBadRequest.WithMessage("state is required")
	}

	if query.Scope == "" {
		return api_err.ErrBadRequest.WithMessage("scope is required")
	}

	return nil
}
