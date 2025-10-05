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
	sessionService := service.NewSessionService(sessionRepo)

	// User
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	// OAuth2
	oauth2Handler := handler.NewOAuth2Handler(authCodeService, sessionService, clientService, userService)

	v1 := r.Group("/v1/arc")

	// OAuth2 endpoints
	oauth2 := r.Group("/oauth2")
	{
		oauth2.GET("/authorize", oauth2Handler.Authorize)
		oauth2.POST("/token", oauth2Handler.Token)
		oauth2.POST("/login", oauth2Handler.Login)
	}

	protected := v1.Group("/", middleware.AuthMiddleware())

	// API endpoints
	adminClient := protected.Group("/admin")
	{
		adminClient.GET("/clients", clientHandler.FindAll)
		adminClient.GET("/clients/:clientId", clientHandler.FindByID)
		adminClient.POST("/clients", clientHandler.Create)
		adminClient.PUT("/clients/:clientId", clientHandler.Update)
		adminClient.DELETE("/admin/clients/:clientId", clientHandler.Delete)
	}

	adminClientSecret := adminClient.Group("/clients/:clientId/secrets")
	{
		adminClientSecret.GET("/", clientSecretHandler.FindAll)
		adminClientSecret.POST("/", clientSecretHandler.Create)
		adminClientSecret.DELETE("/:clientSecretId", clientSecretHandler.Delete)
	}

	// Swagger
	{
		r.GET("/api/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
