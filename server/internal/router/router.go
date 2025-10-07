package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-arc/internal/config"
	"github.com/lubosgarancovsky/eden-arc/internal/handler"
	"github.com/lubosgarancovsky/eden-arc/internal/middleware"
	"github.com/lubosgarancovsky/eden-arc/internal/repository"
	"github.com/lubosgarancovsky/eden-arc/internal/service"
	"github.com/lubosgarancovsky/go-kit/rsql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB, cfg *config.Config) {
	parser := rsql.New()

	// Global middleware
	r.Use(middleware.ErrorMiddleware())

	// Email
	emailService := service.NewEmailService(cfg)

	// Clients
	clientRepo := repository.NewClientRepository(db)
	clientService := service.NewClientService(clientRepo)
	clientHandler := handler.NewClientHandler(parser, clientService)

	// Secrets
	clientSecretRepo := repository.NewClientSecretRepository(db)
	clientSecretService := service.NewClientSecretService(clientSecretRepo)
	clientSecretHandler := handler.NewClientSecretHandler(parser, clientSecretService)

	// Authorization code
	authCodeRepo := repository.NewAuthCodeRepository(db)
	authCodeService := service.NewAuthCodeService(authCodeRepo)

	// Session
	sessionRepo := repository.NewSessionRepository(db)
	sessionService := service.NewSessionService(cfg, sessionRepo)

	// Recovery token
	recoveryTokenRepo := repository.NewRecoveryTokenRepository(db)
	recoveryTokenService := service.NewRecoveryTokenService(cfg, recoveryTokenRepo)

	// User
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(cfg, userRepo, recoveryTokenService, emailService)
	userHandler := handler.NewUserHandler(userService)

	// OAuth2
	oauth2Service := service.NewOAuthService(cfg, clientService, clientSecretService, sessionService, userService, authCodeService)
	oauth2Handler := handler.NewOAuth2Handler(oauth2Service)

	v1 := r.Group("/v1/ark")
	protected := v1.Group("/", middleware.AuthMiddleware())

	// OAuth2 endpoints
	oauth2 := r.Group("/oauth2")
	{
		oauth2.GET("/authorize", oauth2Handler.Authorize)
		oauth2.POST("/token", oauth2Handler.Token)
		oauth2.POST("/login", oauth2Handler.Login)
		oauth2.GET("/logout", oauth2Handler.Logout)
	}

	// API endpoints
	adminClient := protected.Group("/admin")
	{
		adminClient.GET("/clients", clientHandler.FindAll)
		adminClient.GET("/clients/:clientId", clientHandler.FindByID)
		adminClient.POST("/clients", clientHandler.Create)
		adminClient.PUT("/clients/:clientId", clientHandler.Update)
		adminClient.DELETE("/clients/:clientId", clientHandler.Delete)
	}

	adminClientSecret := adminClient.Group("/clients/:clientId/secrets")
	{
		adminClientSecret.GET("/", clientSecretHandler.FindAll)
		adminClientSecret.POST("/", clientSecretHandler.Create)
		adminClientSecret.DELETE("/:clientSecretId", clientSecretHandler.Delete)
	}

	userUser := v1.Group("/users")
	{
		userUser.POST("/request-reset-password", userHandler.RequestResetPassword)
		userUser.POST("/reset-password", userHandler.ResetPassword)
	}

	// Swagger
	{
		r.GET("/api/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
